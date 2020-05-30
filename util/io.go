package util

import "os"

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func FetchAppDataPath() string {
	return os.Getenv("APPDATA")
}
