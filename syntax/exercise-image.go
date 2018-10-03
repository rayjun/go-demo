package main
 
import (
    "image"
    "image/color"
	"golang.org/x/tour/pic"
)
 
type Image struct{
    x, y, w, h int
    v uint8
}
 
func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}
 
func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.w, img.h)
}
 
func (img Image) At(x, y int) color.Color {
    return color.RGBA{img.v, img.v, 255, 255}
}
 
func main() {
    m := Image{1, 2, 200, 30, 11}
    pic.ShowImage(m)
}
