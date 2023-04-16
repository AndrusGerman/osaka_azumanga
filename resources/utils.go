package resources

import (
	"log"
	"os"
	"osaka_azumanga/utils"
	"path"
)

var appFolder = "andruscodex/osaka_azumanga"

func GetPathFolder() string {
	var dir, err = os.UserConfigDir()
	if err != nil {
		log.Println("err get user config folder,", err)
		return ""
	}
	return utils.ParsePath(path.Join(dir, appFolder))
}
