/*
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * API version: 2.18.0
 * Contact: support@bungie.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package D2Client

// If this item can exist in an inventory, this block will be non-null. In practice, every item that currently exists has one of these blocks. But note that it is not necessarily guaranteed.
type AllOfDestinyDefinitionsDestinyInventoryItemDefinitionInventory struct {
	// If this string is populated, you can't have more than one stack with this label in a given inventory. Note that this is different from the equipping block's unique label, which is used for equipping uniqueness.
	StackUniqueLabel string `json:"stackUniqueLabel,omitempty"`
	// The maximum quantity of this item that can exist in a stack.
	MaxStackSize int32 `json:"maxStackSize,omitempty"`
	// The hash identifier for the DestinyInventoryBucketDefinition to which this item belongs. I should have named this \"bucketHash\", but too many things refer to it now. Sigh.
	BucketTypeHash int32 `json:"bucketTypeHash,omitempty"`
	// If the item is picked up by the lost loot queue, this is the hash identifier for the DestinyInventoryBucketDefinition into which it will be placed. Again, I should have named this recoveryBucketHash instead.
	RecoveryBucketTypeHash int32 `json:"recoveryBucketTypeHash,omitempty"`
	// The hash identifier for the Tier Type of the item, use to look up its DestinyItemTierTypeDefinition if you need to show localized data for the item's tier.
	TierTypeHash int32 `json:"tierTypeHash,omitempty"`
	// If TRUE, this item is instanced. Otherwise, it is a generic item that merely has a quantity in a stack (like Glimmer).
	IsInstanceItem bool `json:"isInstanceItem,omitempty"`
	// The localized name of the tier type, which is a useful shortcut so you don't have to look up the definition every time. However, it's mostly a holdover from days before we had a DestinyItemTierTypeDefinition to refer to.
	TierTypeName string `json:"tierTypeName,omitempty"`
	// The enumeration matching the tier type of the item to known values, again for convenience sake.
	TierType int32 `json:"tierType,omitempty"`
	// The tooltip message to show, if any, when the item expires.
	ExpirationTooltip string `json:"expirationTooltip,omitempty"`
	// If the item expires while playing in an activity, we show a different message.
	ExpiredInActivityMessage string `json:"expiredInActivityMessage,omitempty"`
	// If the item expires in orbit, we show a... more different message. (\"Consummate V's, consummate!\")
	ExpiredInOrbitMessage string `json:"expiredInOrbitMessage,omitempty"`
	SuppressExpirationWhenObjectivesComplete bool `json:"suppressExpirationWhenObjectivesComplete,omitempty"`
	// A reference to the associated crafting 'recipe' item definition, if this item can be crafted.
	RecipeItemHash int32 `json:"recipeItemHash,omitempty"`
}
