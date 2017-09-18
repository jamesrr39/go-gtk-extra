package gogtk2extra

import (
	"image"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

func NewGtkImageFromImage(picture image.Image) *gtk.Image {
	return gtk.NewImageFromPixbuf(NewGdkPixBufFromImage(picture))
}

func NewGdkPixBufFromImage(picture image.Image) *gdkpixbuf.Pixbuf {
	width := picture.Bounds().Max.X
	height := picture.Bounds().Max.Y

	pixbuf := gdkpixbuf.NewPixbuf(gdkpixbuf.GDK_COLORSPACE_RGB, true, 8, width, height)
	pixelSlice := pixbuf.GetPixels()

	const bytesPerPixel = 4
	indexInPixelSlice := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colour := picture.At(x, y)
			r, g, b, a := colour.RGBA()

			pixelSlice[indexInPixelSlice] = uint32ColourToByte(r)
			pixelSlice[indexInPixelSlice+1] = uint32ColourToByte(g)
			pixelSlice[indexInPixelSlice+2] = uint32ColourToByte(b)
			pixelSlice[indexInPixelSlice+3] = uint32ColourToByte(a)

			indexInPixelSlice += bytesPerPixel
		}
	}

	return pixbuf
}

func uint32ColourToByte(value uint32) byte {
	const ratio = float64(256) / float64(65536)
	byteValue := ratio * float64(value)
	if byteValue > 255 {
		return byte(255)
	}
	return byte(byteValue)
}
