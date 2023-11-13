package main

//go:generate oapi-codegen --package=main -generate=types -o ./bungie.gen.go ./openapi.json
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/eric-gt/d2-vault-inspector/D2Client"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Guardian struct {
    ID int32 `json:"id"`
    Name string `json:"name"`
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
    client := getBungieClient();
    playerName := c.Query("searchString");
    response, err := client.Destiny2Api.Destiny2SearchDestinyPlayerByBungieName(c, D2Client.UserExactSearchRequest{DisplayName: playerName}, -1);
    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "search-results.tmpl", gin.H{
            "results":[]D2Client.UserUserInfoCard{},
            "message": "Results could not be loaded, the server encountered an error",
        });
        return;
    }
    body, err := io.ReadAll(response.Body);
    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "search-results.tmpl", gin.H{
            "results":[]D2Client.UserUserInfoCard{},
            "message": "Results could not be loaded, the server encountered an error",
        });
        return;
    }
    var d2Response D2Client.BungieResponse[[]D2Client.UserUserInfoCard];
    err = json.Unmarshal(body, d2Response);

    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "search-results.tmpl", gin.H{
            "results":[]D2Client.UserUserInfoCard{},
            "message": "Results could not be loaded. Refresh the page and try again",
        });
        return;
    }

    if response.StatusCode > 300 {
        logrus.Error(response);
        c.HTML(http.StatusInternalServerError, "search-results.tmpl", gin.H{
            "results":[]D2Client.UserUserInfoCard{},
            "message": "Results could not be loaded, the Destiny API responded with an error",
        });
        return;
    }
    if d2Response.ErrorCode > 1 {
        handleBungieErrorCode[[]D2Client.UserUserInfoCard](d2Response);
        return;
    }
    results := d2Response.Response;
    if results == nil {
        c.HTML(http.StatusNotFound, "search-results.tmpl", gin.H{
            "results": []D2Client.UserUserInfoCard{},
        });
    }
    c.HTML(http.StatusOK, "search-results.tmpl", gin.H{
        "results": results,
    });
}

func constructVaultForGuardian(c *gin.Context) {
    guardianId := c.Query("guardianId");
    membershipType := c.Query("membershipType");
    parsedId, err := strconv.ParseInt(guardianId, 10, 32);
    if err != nil {
        c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
            "err": fmt.Sprintf("Error fetching Guardian: %v", err),
        });
        return;
    }
    parsedMembership,err := strconv.ParseInt(membershipType, 10, 32);
    if err != nil {
        c.HTML(http.StatusBadRequest, "err.tmpl", gin.H{
            "err": fmt.Sprintf("Error fetching Guardian: %v", err),
        });
        return;
    }
    client := getBungieClient();
    opts := &D2Client.Destiny2ApiDestiny2GetProfileOpts{};
    response, err := client.Destiny2Api.Destiny2GetProfile(c, parsedId, int32(parsedMembership), opts)
    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "vault.tmpl", gin.H{
            "error": "Unable to load vault, the server encountered an error",
        });
        return
    }

    body, err := io.ReadAll(response.Body);
    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "vault.tmpl", gin.H{
            "error": "Unable to load vault, the server encountered an error",
        });
        return
    }

    if response.Status != "200 OK" {
        logrus.Error(response);
        c.HTML(http.StatusInternalServerError, "vault.tmpl", gin.H{
            "error": "Unable to load vault, the server received an error from Bungie",
        });
        return
    }

    var bungieResponse D2Client.BungieResponse[D2Client.DestinyResponsesDestinyProfileResponse];
    err = json.Unmarshal(body, bungieResponse);
    if err != nil {
        logrus.Error(err);
        c.HTML(http.StatusInternalServerError, "vault.tmpl", gin.H{
            "error": "Unable to load vault, the server encountered an error",
        });
        return
    }

    if bungieResponse.ErrorCode > 1 {
        handleBungieErrorCode[D2Client.DestinyResponsesDestinyProfileResponse](bungieResponse);
        return;
    }

   if bungieResponse.Response == nil {
        c.HTML(http.StatusNotFound, "vault.tmpl", gin.H{
            "guardian": nil,
            "error": "Error fetching vault data, the Bungie API did not return a profile",
        });
        return;
   }

    guardian := bungieResponse.Response;

    c.HTML(http.StatusOK, "vault.tmpl", gin.H{
        "Guardian": guardian,
    });
    return;
}

func filter(ss []Guardian, prefix string) (ret []Guardian) {
    for _,s := range ss {
        if strings.HasPrefix(s.Name, prefix) {
            ret = append(ret, s);
        }
    }
    return
}

func getBungieClient() (client *D2Client.APIClient) {
    cfg := D2Client.NewConfiguration();
    cfg.AddDefaultHeader("X-Api-Key", os.Getenv("BUNGIE_API_KEY"));
    cfg.UserAgent = "D2-Vault-Inspector";
    client = D2Client.NewAPIClient(cfg);
    return; 
}

func handleBungieErrorCode[T any](response D2Client.BungieResponse[T]) {
    //do the needful
}
