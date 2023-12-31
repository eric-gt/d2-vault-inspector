# DestinyDefinitionsDestinyItemCategoryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**Visible** | **bool** | If True, this category should be visible in UI. Sometimes we make categories that we don&#x27;t think are interesting externally. It&#x27;s up to you if you want to skip on showing them. | [optional] [default to null]
**Deprecated** | **bool** | If True, this category has been deprecated: it may have no items left, or there may be only legacy items that remain in it which are no longer relevant to the game. | [optional] [default to null]
**ShortTitle** | **string** | A shortened version of the title. The reason why we have this is because the Armory in German had titles that were too long to display in our UI, so these were localized abbreviated versions of those categories. The property still exists today, even though the Armory doesn&#x27;t exist for D2... yet. | [optional] [default to null]
**ItemTypeRegex** | **string** | The janky regular expression we used against the item type to try and discern whether the item belongs to this category. | [optional] [default to null]
**GrantDestinyBreakerType** | **int32** | If the item in question has this category, it also should have this breaker type. | [optional] [default to null]
**PlugCategoryIdentifier** | **string** | If the item is a plug, this is the identifier we expect to find associated with it if it is in this category. | [optional] [default to null]
**ItemTypeRegexNot** | **string** | If the item type matches this janky regex, it does *not* belong to this category. | [optional] [default to null]
**OriginBucketIdentifier** | **string** | If the item belongs to this bucket, it does belong to this category. | [optional] [default to null]
**GrantDestinyItemType** | **int32** | If an item belongs to this category, it will also receive this item type. This is now how DestinyItemType is populated for items: it used to be an even jankier process, but that&#x27;s a story that requires more alcohol. | [optional] [default to null]
**GrantDestinySubType** | **int32** | If an item belongs to this category, it will also receive this subtype enum value.  I know what you&#x27;re thinking - what if it belongs to multiple categories that provide sub-types?  The last one processed wins, as is the case with all of these \&quot;grant\&quot; enums. Now you can see one reason why we moved away from these enums... but they&#x27;re so convenient when they work, aren&#x27;t they? | [optional] [default to null]
**GrantDestinyClass** | **int32** | If an item belongs to this category, it will also get this class restriction enum value.  See the other \&quot;grant\&quot;-prefixed properties on this definition for my color commentary. | [optional] [default to null]
**TraitId** | **string** | The traitId that can be found on items that belong to this category. | [optional] [default to null]
**GroupedCategoryHashes** | **[]int32** | If this category is a \&quot;parent\&quot; category of other categories, those children will have their hashes listed in rendering order here, and can be looked up using these hashes against DestinyItemCategoryDefinition.  In this way, you can build up a visual hierarchy of item categories. That&#x27;s what we did, and you can do it too. I believe in you. Yes, you, Carl.  (I hope someone named Carl reads this someday) | [optional] [default to null]
**ParentCategoryHashes** | **[]int32** | All item category hashes of \&quot;parent\&quot; categories: categories that contain this as a child through the hierarchy of groupedCategoryHashes. It&#x27;s a bit redundant, but having this child-centric list speeds up some calculations. | [optional] [default to null]
**GroupCategoryOnly** | **bool** | If true, this category is only used for grouping, and should not be evaluated with its own checks. Rather, the item only has this category if it has one of its child categories. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

