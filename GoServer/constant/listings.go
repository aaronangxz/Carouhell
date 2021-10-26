package constant

const (
	//listing constant groups
	LISTING_CONSTANT_TYPE_ITEM_STATUS   = 0
	LISTING_CONSTANT_TYPE_ITEM_CATEGORY = 1
	LISTING_CONSTANT_TYPE_SHIPPING_TYPE = 2
	LISTING_CONSTANT_TYPE_PAYMENT_TYPE  = 3
	LISTING_CONSTANT_TYPE_REACTION_TYPE = 4
	LISTING_CONSTANT_TYPE_LOCATION      = 5

	//item statuses
	ITEM_STATUS_ALL     = 0
	ITEM_STATUS_NORMAL  = 1
	ITEM_STATUS_SOLDOUT = 2
	ITEM_STATUS_DELETED = 3

	//item categories
	ITEM_CATEGORY_ALL         = 0
	ITEM_CATEGORY_FOOD        = 1
	ITEM_CATEGORY_ELECTRONICS = 2

	//shipping type
	SHIPPING_TYPE_MAIL    = 0
	SHIPPING_TYPE_COURIER = 1
	SHIPPING_TYPE_COD     = 2

	//payment type
	PAYMENT_TYPE_WALLET        = 0
	PAYMENT_TYPE_CREDIT_CARD   = 1
	PAYMENT_TYPE_BANK_TRANSFER = 2
	PAYMENT_TYPE_INSTALMENT    = 3
	PAYMENT_TYPE_CASH          = 4

	LISTING_REACTION_TYPE_LIKE    = 0
	LISTING_REACTION_TYPE_COMMENT = 1

	LOCATION_1  = 0
	LOCATION_2  = 1
	LOCATION_3  = 2
	LOCATION_4  = 3
	LOCATION_5  = 4
	LOCATION_6  = 5
	LOCATION_7  = 6
	LOCATION_8  = 7
	LOCATION_9  = 8
	LOCATION_10 = 9
)

func CheckListingConstant(field int, param uint32) bool {

	ItemStatusList := [4]int{0, 1, 2, 3}
	ItemCategoryList := [2]int{0, 1}
	ShippingTypeList := [3]int{0, 1, 2}
	PaymentTypeList := [5]int{0, 1, 2, 3, 4}
	ListingReactionList := [2]int{0, 1}
	LocationList := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	switch field {
	case 0:
		for _, elem := range ItemStatusList {
			if param == uint32(elem) {
				return true
			}
		}
	case 1:
		for _, elem := range ItemCategoryList {
			if param == uint32(elem) {
				return true
			}
		}
	case 2:
		for _, elem := range ShippingTypeList {
			if param == uint32(elem) {
				return true
			}
		}
	case 3:
		for _, elem := range PaymentTypeList {
			if param == uint32(elem) {
				return true
			}
		}
	case 4:
		for _, elem := range ListingReactionList {
			if param == uint32(elem) {
				return true
			}
		}
	case 5:
		for _, elem := range LocationList {
			if param == uint32(elem) {
				return true
			}
		}
	}
	return false
}
