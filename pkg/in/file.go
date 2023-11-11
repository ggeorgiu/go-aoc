package in

import "os"

func ReadFileAsString(filePath string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(path + filePath)
	if err != nil {
		panic(err)
	}

	return string(dat)
}
