package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    image := make([][]uint8, dx)
    for i := range image {
        image[i] = make([]uint8, dy)
    }

    for i:= 0; i < dx; i++ {
        for j:= 0; j < dx; j++ {
            //image[i][j] = uint8(i*j)
            //image[i][j] = uint8(i*^j)
            //image[i][j] = uint8((i+j)/2)
        }
    }

    return image
}

func main() {
    pic.Show(Pic)
}
