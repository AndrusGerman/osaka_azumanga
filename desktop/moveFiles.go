package desktop

import (
	"log"
	"os"
	"os/exec"
	"osaka_azumanga/utils"
	"path"
)

var filesFolder = "backup"

func MoveFiles() error {
	var folder = GetDesktopFolder()
	var newFoder = utils.ParsePath(path.Join(folder, filesFolder))

	var files, err = os.ReadDir(folder)
	if err != nil {
		return err
	}

	os.MkdirAll(newFoder, 0777)

	for _, de := range files {
		var fileToMove = utils.ParsePath(path.Join(folder, de.Name()))
		log.Println(fileToMove, newFoder)
		cmd := exec.Command("cmd", "/C", "move", fileToMove, newFoder)
		bt, err := cmd.Output()
		log.Println(string(bt), err)
	}

	return nil
}
