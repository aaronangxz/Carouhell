package constant

const (
	ACC_CREATION_SECURITY_QUESTION = 0
	ACC_STATUS                     = 1
	USER_REVIEW_STARS              = 2
	SELLER_TYPE                    = 3

	SECURITY_QUESTION_1 = 0
	SECURITY_QUESTION_2 = 1
	SECURITY_QUESTION_3 = 2
	SECURITY_QUESTION_4 = 3
	SECURITY_QUESTION_5 = 4

	ACC_STATUS_ACTIVE   = 0
	ACC_STATUS_INACTIVE = 1
	ACC_STATUS_DISABLED = 2
	ACC_STATUS_DELETED  = 3

	USER_REVIEW_ONE_STAR   = 1
	USER_REVIEW_TWO_STAR   = 2
	USER_REVIEW_THREE_STAR = 3
	USER_REVIEW_FOUR_STAR  = 4
	USER_REVIEW_FIVE_STAR  = 5

	SELLER_TYPE_OFFICIAL  = 1
	SELLER_TYPE_PREFERRED = 2
)

func CheckAccConstant(field int, param uint32) bool {

	SecurityQuestionList := [5]int{0, 1, 2, 3, 4}
	AccStatusList := [4]int{0, 1, 2, 3}
	UserReviewList := [5]int{1, 2, 3, 4, 5}
	SellerTypeList := [2]int{1, 2}

	switch field {
	case ACC_CREATION_SECURITY_QUESTION:
		for _, elem := range SecurityQuestionList {
			if param == uint32(elem) {
				return true
			}
		}
	case ACC_STATUS:
		for _, elem := range AccStatusList {
			if param == uint32(elem) {
				return true
			}
		}
	case USER_REVIEW_STARS:
		for _, elem := range UserReviewList {
			if param == uint32(elem) {
				return true
			}
		}
	case SELLER_TYPE:
		for _, elem := range SellerTypeList {
			if param == uint32(elem) {
				return true
			}
		}
	}
	return false
}
