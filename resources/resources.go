package resources

import (
	_ "embed"
	"osaka_azumanga/utils"
	"path"
)

//go:embed assets/main.jpg
var MainImage []byte
var MainImageName = "osaka_azumanga_andruscodex.jpg"

func GetResourcePath(name string) string {
	return utils.ParsePath(path.Join(GetPathFolder(), name))
}
