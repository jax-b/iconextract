package main

import (
	"image/png"
	"os"

	"github.com/jax-b/iconextract"
)

func main() {
	image, err := iconextract.ExtractIcon("/Windows/explorer.exe", 0)
	f, err := os.Create("outimage2.png")
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, image)
	if err != nil {
		// Handle error
	}
}
