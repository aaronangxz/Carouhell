package constant

const (
	ACC_CREATION_SECURITY_QUESTION = 0
	ACC_STATUS                     = 1
	USER_REVIEW_STARS              = 2

	SECURITY_QUESTION_1 = 0
	SECURITY_QUESTION_2 = 1
	SECURITY_QUESTION_3 = 2
	SECURITY_QUESTION_4 = 3
	SECURITY_QUESTION_5 = 4

	ACC_STATUS_ACTIVE   = 0
	ACC_STATUS_INACTIVE = 1
	ACC_STATUS_DISABLED = 2
	ACC_STATUS_DELETED  = 3

	USER_REVIEW_ONE_STAR   = 0
	USER_REVIEW_TWO_STAR   = 1
	USER_REVIEW_THREE_STAR = 2
	USER_REVIEW_FOUR_STAR  = 3
	USER_REVIEW_FIVE_STAR  = 4
)

func CheckAccConstant(field int, param uint32) bool {

	SecurityQuestionList := [5]int{0, 1, 2, 3, 4}
	AccStatusList := [4]int{0, 1, 2, 3}
	UserReviewList := [5]int{0, 1, 2, 3, 4}

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
	}
	return false
}
