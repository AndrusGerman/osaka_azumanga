package resources

import (
	"os"
)

func SetResources() error {
	os.MkdirAll(GetPathFolder(), 0777)

	var filePath = GetResourcePath(MainImageName)
	os.WriteFile(filePath, MainImage, 0666)
	return nil
}
