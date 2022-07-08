package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func list_builds() []fs.DirEntry {
    build_list, err := os.ReadDir("./builds")
    checkErr(err)
    return build_list
}

func detect_images(game_folder string) []string {
    var image_list []string
	filepath.Walk(game_folder, func(path string, info os.FileInfo, err error) error {
        checkErr(err)
		if !info.IsDir() {
			if filepath.Ext(info.Name()) == ".png" {
                image_list = append(image_list, path)
            }}
        return nil
	})
    return image_list
}

