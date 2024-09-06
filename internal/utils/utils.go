package utils

type M map[string]interface{}

func SendResult(code int, msg string, data interface{}) M {
	if data != nil {
		return M{
			"code": code,
			"msg":  msg,
			"data": data,
		}
	} else {
		return M{
			"code": code,
			"msg":  msg,
		}
	}
}

func Find(val string, slice []string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
