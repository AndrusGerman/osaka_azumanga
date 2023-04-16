package resources

import (
	_ "embed"
	"path"
)

//go:embed assets/main.jpg
var MainImage []byte
var MainImageName = "main.jpg"

func GetResourcePath(name string) string {
	return path.Join(GetPathFolder(), name)
}
