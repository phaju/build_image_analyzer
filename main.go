package main

import (
	"log"
	"os"
)

func main() {
    image_size_limits := []int{0, 400, 800, 1200}
    image_counter := []int{0, 0, 0, 0}
    // Get all builds present inside build folder and unzip them one by one
    image_list := detect_images("builds")
    log.Println("Total Files: ", len(image_list))
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
    log.Println(image_counter)
}
