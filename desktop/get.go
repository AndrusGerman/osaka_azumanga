package desktop

import "golang.org/x/sys/windows/registry"

func GetDesktopFolder() string {
	reg, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders`, registry.READ)
	if err != nil {
		panic(err)
	}
	defer reg.Close()

	value, _, _ := reg.GetStringValue("Desktop")
	return value
}
