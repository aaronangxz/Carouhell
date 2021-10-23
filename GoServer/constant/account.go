package constant

const (
	ACC_CREATION_SECURITY_QUESTION = 0

	SECURITY_QUESTION_1 = 0
	SECURITY_QUESTION_2 = 1
	SECURITY_QUESTION_3 = 2
	SECURITY_QUESTION_4 = 3
	SECURITY_QUESTION_5 = 4

	ACC_STATUS_ACTIVE   = 0
	ACC_STATUS_INACTIVE = 1
	ACC_STATUS_DISABLED = 2
	ACC_STATUS_DELETED  = 3
)

func CheckAccConstant(field int, param uint32) bool {

	SecurityQuestionList := [5]int{0, 1, 2, 3, 4}

	switch field {
	case 0:
		for _, elem := range SecurityQuestionList {
			if param == uint32(elem) {
				return true
			}
		}
	}
	return false
}
