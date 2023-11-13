# AllOfDestinyDefinitionsDestinyInventoryItemDefinitionPreview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScreenStyle** | **string** | A string that the game UI uses as a hint for which detail screen to show for the item. You, too, can leverage this for your own custom screen detail views. Note, however, that these are arbitrarily defined by designers: there&#x27;s no guarantees of a fixed, known number of these - so fall back to something reasonable if you don&#x27;t recognize it. | [optional] [default to null]
**PreviewVendorHash** | **int32** | If the preview data is derived from a fake \&quot;Preview\&quot; Vendor, this will be the hash identifier for the DestinyVendorDefinition of that fake vendor. | [optional] [default to null]
**ArtifactHash** | **int32** | If this item should show you Artifact information when you preview it, this is the hash identifier of the DestinyArtifactDefinition for the artifact whose data should be shown. | [optional] [default to null]
**PreviewActionString** | **string** | If the preview has an associated action (like \&quot;Open\&quot;), this will be the localized string for that action. | [optional] [default to null]
**DerivedItemCategories** | [**[]DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition**](Destiny.Definitions.Items.DestinyDerivedItemCategoryDefinition.md) | This is a list of the items being previewed, categorized in the same way as they are in the preview UI. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

