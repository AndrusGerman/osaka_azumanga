package resources

import (
	"log"
	"os"
)

func SetResources() error {
	os.MkdirAll(GetPathFolder(), 0777)

	var filePath = GetResourcePath(MainImageName)
	log.Println("resources set in: ", filePath)
	var file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(MainImage)
	return nil
}
