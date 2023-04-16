package main

import (
	"log"
	"osaka_azumanga/desktop"
	"osaka_azumanga/resources"
	"osaka_azumanga/wallpaper"
)

func main() {

	var err error
	if err = desktop.MoveFiles(); err != nil {
		log.Println("err set resources ", err)
		return
	}

	if err = resources.SetResources(); err != nil {
		log.Println("err set resources ", err)
		return
	}
	if err = wallpaper.SetWallpaper(resources.GetResourcePath(resources.MainImageName)); err != nil {
		log.Println("err set wallpaper ", err)
		return
	}
}
