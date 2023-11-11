package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Guardian struct {
    ID int32 `json:"id"`
    Name string `json:"name"`
}

type BungieResponse[T any] struct {
    Response T;
    ErrorCode int32;
    ThrottleSeconds int32;
    Message string;
    MessageData MessageData;
    DetailedErrorTrace string;
}

type MessageData struct {
    DictionaryContents string;
    DictionaryKeyTyp string;
}

type UserInfoCard struct {
    SupplementalDisplayName string `json:"supplementalDisplayName"`;
    IconPath string `json:"iconPath"`;
    CrossSaveOverride int32 `json:"crossSaveOverride"`;
    ApplicableMembershipType []int32 `json:"applicableMembershipType"`;
    IsPublic bool `json:"isPublic"`;
    MembershipType int32 `json:"membershipType"`;
    MembershipId int64 `json:"membershipId"`;
    BungieGlobalDisplayName string `json:"bungieGlobalDisplayName"`;
    BungieGlobalDisplayNameCode int16 `json:"bungieGlobalDisplayNameCode"`;
}

type DestinyProfileResponse struct {
    ResponseMintedTimestamp time.Time `json:"responseMintedTimestamp"`;
    SecondaryComponentsMintedTimestamp time.Time `json:"secondaryComponentsMintedTimestamp"`;
    VendorReceipts VendorReceipts `json:"vendorReceipts"`;
    ProfileInventory ProfileInventory `json:"profileInventory"`;
    ProfileCurrencies ProfileCurrencies `json:"profileCurrencies"`;
    Profile Profile  `json:"profile"`;
    PlatformSilver PlatformSilver `json:"platformSilver"`;
    profileKiosks ProfileKiosks `json:"profileKiosks"`;

}

type VendorReceipts struct {
}
type ProfileInventory struct {
}
type ProfileCurrencies struct {
}
type Profile struct {
}
type PlatformSilver struct {
}
type ProfileKiosks struct {
}

var Guardians []Guardian =  []Guardian{{ID: 1, Name: "guardian1"}, {ID: 2, Name:"guardian2"},{ID: 3, Name:  "guardian3"},{ID: 4, Name:  "aztecross"},{ID:5, Name:"datto"},{ID:6, Name:  "itsmefallout"},{ID:7, Name: "travelDanielle"},{ID: 8, Name:"taekwondork"}};


func main() {
    
	router := gin.Default()
    router.LoadHTMLGlob("templates/*");
    router.Static("/static", "./static/");

    router.GET("/", serveHomePage);
    router.GET("/get-guardians-by-prefix", getGuardiansByNamePrefix);
    router.GET("/get-guardian-by-id", constructVaultForGuardian);

	err := router.Run("127.0.0.1:8080")
	if err != nil {
        msg := fmt.Sprintf("error starting server: %v", err);
        log.Fatal(msg)
	}

}

func serveHomePage(c *gin.Context) {

    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Destiny Vault Inspector",
    })
}

func getGuardiansByNamePrefix(c *gin.Context) {
    prefix := c.Query("searchString");
    if prefix == "" {
        c.HTML(http.StatusOK, "search-results.tmpl", gin.H{
            "results": []Guardian{},
        });
        return;
    }
    results := filter(Guardians, prefix);
    c.HTML(http.StatusOK, "search-results.tmpl", gin.H{
        "results": results,
    });
}

func constructVaultForGuardian(c *gin.Context) {
    guardianId := c.Query("guardianId");
    parsedId, err := strconv.ParseInt(guardianId, 10, 32);
    if err != nil {
        c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
            "err": fmt.Sprintf("Error fetching Guardian: %v", err),
        });
        return;
    }
    for _,guardian := range Guardians {
        if guardian.ID == int32(parsedId) {
            c.HTML(http.StatusOK, "vault.tmpl", gin.H{
                "Guardian": guardian,
            });
            return;
        }
    }
    c.HTML(http.StatusNotFound, "vault.tmpl", gin.H{
        "guardian": nil,
    });
}

func getMembershipTypesByGuardianName(guardianName string, client *http.Client) ([]UserInfoCard) {
    bungieClient := &http.Client{};
    postBody,_ := json.Marshal(map[string]string{"displayName": guardianName});
    responseBody := bytes.NewBuffer(postBody);
    request := getBungieRequest("POST", "Platform/Destiny2/SearchDestinyPlayerByBungieName/-1", responseBody);
    request.Header.Add("Content-Type", "application/json");

    resp, err := bungieClient.Do(request);
    if err != nil {
        return []UserInfoCard{};
    }

    body, err := io.ReadAll(resp.Body);
    if err != nil {
        return []UserInfoCard{}
    }

    defer resp.Body.Close();
    var bungieResponse *BungieResponse[[]UserInfoCard];

    err = json.Unmarshal(body, bungieResponse);
    if err != nil {
        return []UserInfoCard{};
    }

    if bungieResponse.ErrorCode > 1 {
        handleBungieErrorCode[[]UserInfoCard](bungieResponse);
    }
    return bungieResponse.Response;
}

func getGuardianProfileByIdAndType(guardianId int32, membershipType int16) (*DestinyProfile) {
    bungieClient := &http.Client{};
    path := fmt.Sprintf("Destiny2/%v/Profile/%v/", membershipType, guardianId);
    request := getBungieRequest("GET", path, nil);
    response, err := bungieClient.Do(request);
    if err != nil {
        return *DestinyProfile{};
    }

    body, err := io.ReadAll(response.Body);
    if err != nil {
        return *DestinyProfile{};
    }
    var bungieResponse *BungieResponse[DestinyProfileResponse];
    err := json.Unmarshal(body, bungieResponse);

    if err != nil {
        return *DestinyProfile{};
    }

    if bungieResponse.ErrorCode > 1 {
        handleBungieErrorCode[DestinyProfileResponse](bungieResponse);
    }
    return &bungieResponse.Response;
}

func filter(ss []Guardian, prefix string) (ret []Guardian) {
    for _,s := range ss {
        if strings.HasPrefix(s.Name, prefix) {
            ret = append(ret, s);
        }
    }
    return
}

func getBungieRequest(method string, path string, body io.Reader) *http.Request {
    baseUrl :=os.Getenv("BUNGIE_API_URL");
    apiKey := os.Getenv("BUNGIE_API_KEY");
    url := fmt.Sprintf("%v/%v", baseUrl, path);
    request,err := http.NewRequest(method, url, body);
    if err != nil {
        log.Fatal("error constructing HTTP client for Bungie API");
    }
    request.Header.Add("X-API-KEY", apiKey);
    return request;
}

func handleBungieErrorCode[T any](response *BungieResponse[T]) {
    //do the needful
}
