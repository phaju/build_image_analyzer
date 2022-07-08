package main

import (
	"image"
	_ "image/png"
	"io/fs"
	"os"
)

type details struct {
    height int
    width int
}

func check_image_resolution(img fs.File) details {
    img_data, _, err := image.DecodeConfig(img)
    checkErr(err)
    var image_meta details
    image_meta.height = img_data.Height
    image_meta.width = img_data.Width
    return image_meta
}

func map_images(image_list []string, image_counter []int, image_size_limits []int) {
    for _, img := range(image_list) {
        // Open Image File
        image_file, err := os.Open(img)
        checkErr(err)
        defer image_file.Close()

        // Get Resolution and map the sizes to the required group
        image_data := check_image_resolution(image_file)
        if image_data.width > image_size_limits[3] || image_data.height > image_size_limits[3] {
            image_counter[3]++
        } else if image_data.width > image_size_limits[2] || image_data.height > image_size_limits[2] {
            image_counter[2]++
        } else if image_data.width > image_size_limits[1] || image_data.height > image_size_limits[1] {
            image_counter[1]++
        } else {
            image_counter[0]++
        }
    }
}
