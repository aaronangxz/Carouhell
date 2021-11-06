package constant

const (
	SEARCH_AND_FILTERS_CONSTANT_SORTFLAG = 0

	//sort flag
	SEARCH_RESULT_SORTFLAG_DEFAULT    = 0
	SEARCH_RESULT_SORTFLAG_RECENT     = 1
	SEARCH_RESULT_SORTFLAG_PRICE_HIGH = 2
	SEARCH_RESULT_SORTFLAG_PRICE_LOW  = 3
	SEARCH_RESULT_SORTFLAG_POPULAR    = 4
)

func CheckSearchAndFiltersConstant(field int, param uint32) bool {

	SortFlagList := [5]int{0, 1, 2, 3, 4}

	switch field {
	case 0:
		for _, elem := range SortFlagList {
			if param == uint32(elem) {
				return true
			}
		}
	}
	return false
}
