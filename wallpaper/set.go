package wallpaper

import (
	"log"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var regName = "WallPaper"

func SetWallpaper(path string) error {
	if err := SetWallpaperInRegedit(path); err != nil {
		return err
	}
	return RestartConfigUser()
}

func RestartConfigUser() error {
	cmd := exec.Command("RUNDLL32.EXE", "user32.dll,UpdatePerUserSystemParameters")
	out, err := cmd.Output()
	log.Println("cmdoutput: ", string(out))
	return err
}

func SetWallpaperInRegedit(path string) error {
	reg, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.WRITE)
	if err != nil {
		return err
	}
	defer reg.Close()

	// val, _, _ := reg.GetStringValue(regName)
	// log.Println("Current Wallpaper: ", val)

	return reg.SetStringValue(regName, strings.ReplaceAll(path, "\\", "/"))
}
