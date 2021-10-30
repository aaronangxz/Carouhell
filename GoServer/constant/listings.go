package constant

const (
	//listing constant groups
	LISTING_CONSTANT_TYPE_ITEM_STATUS        = 0
	LISTING_CONSTANT_TYPE_ITEM_CATEGORY      = 1
	LISTING_CONSTANT_TYPE_REACTION_TYPE      = 2
	LISTING_CONSTANT_TYPE_LOCATION           = 3
	LISTING_CONSTANT_TYPE_LOCATION_DIRECTION = 4

	//item_status
	ITEM_STATUS_ALL     = 0
	ITEM_STATUS_NORMAL  = 1
	ITEM_STATUS_SOLDOUT = 2
	ITEM_STATUS_DELETED = 3

	//item_category
	ITEM_CATEGORY_FOOD_DRINKS          = 0
	ITEM_CATEGORY_COMPUTERS_TECH       = 1
	ITEM_CATEGORY_MOBILE_PHONE_GADGETS = 2
	ITEM_CATEGORY_MENS_FASHION         = 3
	ITEM_CATEGORY_WOMENS_FASHION       = 4
	ITEM_CATEGORY_PETS                 = 5
	ITEM_CATEGORY_FURNITURE            = 6
	ITEM_CATEGORY_CARS                 = 7
	ITEM_CATEGORY_BEAUTY               = 8
	ITEM_CATEGORY_GAMES                = 9
	ITEM_CATEGORY_HOME_APPLIANCES      = 10
	ITEM_CATEGORY_BABIES_KIDS          = 11
	ITEM_CATEGORY_HOBBIES_TOYS         = 12
	ITEM_CATEGORY_TICKETS              = 13
	ITEM_CATEGORY_SPORTS               = 14
	ITEM_CATEGORY_VIRTUAL_GOODS        = 15
	ITEM_CATEGORY_PROPERTY             = 16
	ITEM_CATEGORY_LUXURY               = 17

	//listing_reaction_type
	LISTING_REACTION_TYPE_LIKE    = 0
	LISTING_REACTION_TYPE_COMMENT = 1

	//item_location
	LOCATION_ANG_MO_KIO      = 0
	LOCATION_BEDOK           = 1
	LOCATION_BISHAN          = 2
	LOCATION_BOON_LAY        = 3
	LOCATION_BUKIT_BATOK     = 4
	LOCATION_BUKIT_MERAH     = 5
	LOCATION_BUKIT_PANJANG   = 6
	LOCATION_BUKIT_TIMAH     = 7
	LOCATION_CHANGI          = 8
	LOCATION_CHOA_CHU_KANG   = 9
	LOCATION_CLEMENTI        = 10
	LOCATION_DOWNTOWN        = 11
	LOCATION_GEYLANG         = 12
	LOCATION_HOUGANG         = 13
	LOCATION_JURONG_EAST     = 14
	LOCATION_JURONG_WEST     = 15
	LOCATION_KALLANG         = 16
	LOCATION_LIM_CHU_KANG    = 17
	LOCATION_MANDAI          = 18
	LOCATION_MARINA          = 19
	LOCATION_MARINE_PARADE   = 20
	LOCATION_NEWTON          = 21
	LOCATION_NOVENA          = 22
	LOCATION_ORCHARD         = 23
	LOCATION_OUTRAM          = 24
	LOCATION_PAYA_LEBAR      = 25
	LOCATION_PIONEER         = 26
	LOCATION_PUNGGOL         = 27
	LOCATION_QUEENSTOWN      = 28
	LOCATION_RIVER_VALLEY    = 29
	LOCATION_ROCHOR          = 30
	LOCATION_SELETAR         = 31
	LOCATION_SEMBAWANG       = 32
	LOCATION_SENGKANG        = 33
	LOCATION_SERANGOON       = 34
	LOCATION_SIMPANG         = 35
	LOCATION_SINGAPORE_RIVER = 36
	LOCATION_SUNGEI_KADUT    = 37
	LOCATION_TAMPINES        = 38
	LOCATION_TANGLIN         = 39
	LOCATION_TENGAH          = 40
	LOCATION_TOA_PAYOH       = 41
	LOCATION_TUAS            = 42
	LOCATION_WOODLANDS       = 43
	LOCATION_YISHUN          = 44

	LOCATION_DIRECTION_NORTH   = 0
	LOCATION_DIRECTION_SOUTH   = 1
	LOCATION_DIRECTION_EAST    = 2
	LOCATION_DIRECTION_WEST    = 3
	LOCATION_DIRECTION_CENTRAL = 4

	LISTING_TRANSACTION_STATUS_SUCCESS = 0
	LISTING_TRANSACTION_STATUS_PENDING = 0
	LISTING_TRANSACTION_STATUS_FAIL    = 0
)

func CheckListingConstant(field int, param uint32) bool {

	itemStatusList := [4]uint32{0, 1, 2, 3}
	itemCategoryList := [18]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17}
	listingReactionList := [2]uint32{0, 1}
	locationList := [45]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44}
	listingDirectionList := [5]uint32{0, 1, 2, 3, 4}

	switch field {
	case LISTING_CONSTANT_TYPE_ITEM_STATUS:
		if param > uint32(len(itemStatusList)-1) {
			return false
		}
		return itemStatusList[param] == param
	case LISTING_CONSTANT_TYPE_ITEM_CATEGORY:
		if param > uint32(len(itemCategoryList)-1) {
			return false
		}
		return itemCategoryList[param] == param
	case LISTING_CONSTANT_TYPE_REACTION_TYPE:
		if param > uint32(len(listingReactionList)-1) {
			return false
		}
		return listingReactionList[param] == param
	case LISTING_CONSTANT_TYPE_LOCATION:
		if param > uint32(len(locationList)-1) {
			return false
		}
		return locationList[param] == param
	case LISTING_CONSTANT_TYPE_LOCATION_DIRECTION:
		if param > uint32(len(listingDirectionList)-1) {
			return false
		}
		return listingDirectionList[param] == param
	}
	return false
}

func GetLocationFromDirection(param uint32) []uint32 {
	var (
		locationNorthList   = []uint32{0, 2, 17, 18, 27, 31, 32, 35, 37, 41, 43, 44}
		locationSouthList   = []uint32{5, 7, 28, 19, 21, 23}
		locationEastList    = []uint32{1, 8, 12, 13, 16, 25, 33, 34, 38}
		locationWestList    = []uint32{3, 4, 6, 9, 10, 14, 15, 26, 40, 42}
		locationCentralList = []uint32{11, 19, 20, 22, 24, 29, 30, 36, 39}
	)

	switch param {
	case LOCATION_DIRECTION_NORTH:
		return locationNorthList
	case LOCATION_DIRECTION_SOUTH:
		return locationSouthList
	case LOCATION_DIRECTION_EAST:
		return locationEastList
	case LOCATION_DIRECTION_WEST:
		return locationWestList
	case LOCATION_DIRECTION_CENTRAL:
		return locationCentralList
	}
	return nil
}
