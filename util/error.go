package util

func CheckError(err error) {
	if err != nil {
		Error(err.Error())
	}
}
