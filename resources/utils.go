package resources

import (
	"os"
	"osaka_azumanga/utils"
	"path"
)

var appFolder = "andruscodex/osaka_azumanga"

func GetPathFolder() string {
	var dir, err = os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return utils.ParsePath(path.Join(dir, appFolder))
}
