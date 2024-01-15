package file

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func checkImageType(filepath string) (result string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		log.Println(err)
		return
	}
	switch format {
	case "jpeg", "jpg":
		result = "jpg"
	case "png":
		result = "png"
	default:
		err = fmt.Errorf("Unknown file type")
		log.Println(err)
	}
	return
}

func encodeJpg(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	dir, filename := filepath.Split(path)
	filename = strings.Split(filename, ".")[0]
	output, err := os.Create(dir + filename + ".webp")
	if err != nil {
		log.Println(err)
		return
	}

	defer output.Close()
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Println(err)
		return
	}

	err = webp.Encode(output, img, options)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func encodePng(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	dir, filename := filepath.Split(path)
	filename = strings.Split(filename, ".")[0]
	output, err := os.Create(dir + filename + ".webp")
	if err != nil {
		log.Println(err)
		return
	}

	defer output.Close()
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Println(err)
		return
	}

	err = webp.Encode(output, img, options)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
