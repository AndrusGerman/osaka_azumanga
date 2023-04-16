package desktop

import (
	"fmt"
	"os"
	"osaka_azumanga/utils"
	"path"
)

func CreateImages(name string, image []byte) error {
	var folder = GetDesktopFolder()
	for i := 0; i < 160; i++ {
		var fileName = utils.ParsePath(path.Join(folder, fmt.Sprintf("%d-%s", i, name)))
		os.WriteFile(fileName, image, 0666)
	}
	return nil
}
