package constant

const (
	LISTING_CONSTANT_TYPE_ITEM_STATUS   = 0
	LISTING_CONSTANT_TYPE_ITEM_CATEGORY = 1
	LISTING_CONSTANT_TYPE_SHIPPING_TYPE = 2
	LISTING_CONSTANT_TYPE_PAYMENT_TYPE  = 3

	ITEM_STATUS_ALL     = 0
	ITEM_STATUS_NORMAL  = 1
	ITEM_STATUS_SOLDOUT = 2
	ITEM_STATUS_DELETED = 3

	ITEM_CATEGORY_FOOD        = 0
	ITEM_CATEGORY_ELECTRONICS = 1

	SHIPPING_TYPE_MAIL    = 0
	SHIPPING_TYPE_COURIER = 1
	SHIPPING_TYPE_COD     = 2

	PAYMENT_TYPE_WALLET        = 0
	PAYMENT_TYPE_CREDIT_CARD   = 1
	PAYMENT_TYPE_BANK_TRANSFER = 2
	PAYMENR_TYPE_INSTALMENT    = 3
	PAYMENT_TYPE_CASH          = 4
)

func CheckListingConstant(field int, param uint32) bool {

	ItemStatusList := [4]int{0, 1, 2, 3}
	ItemCategoryList := [2]int{0, 1}
	ShippingTypeList := [3]int{0, 1, 2}
	PaymentTypeList := [5]int{0, 1, 2, 3, 4}

	switch field {
	case 1:
		for _, elem := range ItemStatusList {
			if param == uint32(elem) {
				return true
			}
		}
	case 2:
		for _, elem := range ItemCategoryList {
			if param == uint32(elem) {
				return true
			}
		}
	case 3:
		for _, elem := range ShippingTypeList {
			if param == uint32(elem) {
				return true
			}
		}
	case 4:
		for _, elem := range PaymentTypeList {
			if param == uint32(elem) {
				return true
			}
		}
	}
	return false
}
