package iconextract

import (
	"errors"
	"fmt"
	"image"
	"strings"
	"unicode/utf16"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

// StringToUTF16Ptr converts a Go string into a pointer to a null-terminated UTF-16 wide string.
// This assumes str is of a UTF-8 compatible encoding so that it can be re-encoded as UTF-16.
func stringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}

// NumberOfIcons returns the number of icons in a file
// Windows Filepath must be the absolute file path and contain / as the directory seperators
// example "/Windows/explorer.exe"
func NumberOfIcons(filepath string, index int32) uint {
	hinst := win.GetModuleHandle(nil)
	numIcons := uint(uint64(win.ExtractIcon(hinst, stringToUTF16Ptr(filepath), int32(-1))))
	return numIcons
}

// ExtractIcon creates a image.rgba from a file
// Windows Filepath must be the absolute file path and contain / as the directory seperators
// example "/Windows/explorer.exe"
func ExtractIcon(filepath string, index int32) (*image.RGBA, error) {
	flpth := strings.Split(filepath, ".")
	extention := flpth[len(flpth)-1]
	var bmp *walk.Bitmap
	var err error
	switch extention {
	case "exe":
		fmt.Print("image is EXE\n")
		// hinst := win.GetModuleHandle(nil)
		var winhico win.HICON

		// winhico = win.ExtractIcon(hinst, stringToUTF16Ptr(filepath), index)
		result := win.SHDefExtractIcon(stringToUTF16Ptr(filepath), index, 0, &winhico, nil, 0x1)
		if win.FAILED(result) {
			return nil, errors.New("SHDefExtractIcon Error")
		}
		ico, err := walk.NewIconFromHICONForDPI(winhico, 512)
		if err != nil {
			return nil, err
		}

		sz := walk.Size{
			Height: 512,
			Width:  512,
		}
		bmp, err = walk.NewBitmapFromIcon(ico, sz)
		if err != nil {
			return nil, err
		}
		break
	case "ico":
		fmt.Print("image is ICO")
		var ico *walk.Icon
		ico, err = walk.NewIconFromFile(filepath)
		if err != nil {
			return nil, err
		}
		sz := walk.Size{
			Height: 512,
			Width:  512,
		}
		bmp, err = walk.NewBitmapFromIcon(ico, sz)
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("Image not a supported File Type")
	}
	image, err := bmp.ToImage()
	return image, err
}
