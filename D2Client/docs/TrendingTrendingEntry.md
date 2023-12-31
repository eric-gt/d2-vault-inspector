# TrendingTrendingEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **float64** | The weighted score of this trending item. | [optional] [default to null]
**IsFeatured** | **bool** |  | [optional] [default to null]
**Identifier** | **string** | We don&#x27;t know whether the identifier will be a string, a uint, or a long... so we&#x27;re going to cast it all to a string. But either way, we need any trending item created to have a single unique identifier for its type. | [optional] [default to null]
**EntityType** | **int32** | An enum - unfortunately - dictating all of the possible kinds of trending items that you might get in your result set, in case you want to do custom rendering or call to get the details of the item. | [optional] [default to null]
**DisplayName** | **string** | The localized \&quot;display name/article title/&#x27;primary localized identifier&#x27;\&quot; of the entity. | [optional] [default to null]
**Tagline** | **string** | If the entity has a localized tagline/subtitle/motto/whatever, that is found here. | [optional] [default to null]
**Image** | **string** |  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Link** | **string** |  | [optional] [default to null]
**WebmVideo** | **string** | If this is populated, the entry has a related WebM video to show. I am 100% certain I am going to regret putting this directly on TrendingEntry, but it will work so yolo | [optional] [default to null]
**Mp4Video** | **string** | If this is populated, the entry has a related MP4 video to show. I am 100% certain I am going to regret putting this directly on TrendingEntry, but it will work so yolo | [optional] [default to null]
**FeatureImage** | **string** | If isFeatured, this image will be populated with whatever the featured image is. Note that this will likely be a very large image, so don&#x27;t use it all the time. | [optional] [default to null]
**Items** | [**[]TrendingTrendingEntry**](Trending.TrendingEntry.md) | If the item is of entityType TrendingEntryType.Container, it may have items - also Trending Entries - contained within it. This is the ordered list of those to display under the Container&#x27;s header. | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | If the entry has a date at which it was created, this is that date. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

