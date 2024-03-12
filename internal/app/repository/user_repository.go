package repository

func GetUserPassword(username string) (string, error) {
	if username == "anand" {
		return "123", nil
	}
	return "kkk", nil
}

func GetUserID(username string) (string, error) {
	if username == "anand" {
		return "ID123", nil
	}
	return "asfsaf", nil
}
