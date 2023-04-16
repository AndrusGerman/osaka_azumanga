package resources

import (
	"os"
	"path"
)

var appFolder = "andruscodex/osaka_azumanga"

func GetPathFolder() string {
	var dir, err = os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return path.Join(dir, appFolder)
}
