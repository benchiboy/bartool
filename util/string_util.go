package util

import (
	"fmt"
	"os"
	"regexp"

	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func MatchString(oldStr string, matchConf string) []string {
	var valid = regexp.MustCompile(matchConf)
	fmt.Println(valid.FindAllString(oldStr, -1))
	strTemp := valid.FindAllString(oldStr, -1)
	if strTemp != nil && len(strTemp) > 0 {
		return strTemp
	} else {
		return nil
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CreateQrcode(fileName string, subIcon string, text string, width int, height int) (string, error) {
	// Create the barcode
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)
	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)
	// create the outp
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, qrCode)
	if subIcon != "" {
		bigPngFile, err := os.Open(fileName)
		defer bigPngFile.Close()
		if err != nil {
			log.Println(err)
			return "", err
		}
		bigPng, err := png.Decode(bigPngFile)
		bigBound := bigPng.Bounds()
		subJpgFile, err := os.Open(subIcon)
		defer subJpgFile.Close()
		if err != nil {
			log.Println(err)
			return "", err
		}
		subImg, _ := jpeg.Decode(subJpgFile)
		subBound := subImg.Bounds()
		offset := image.Pt(bigBound.Dx()/2-subBound.Dx()/2, bigBound.Dy()/2-subBound.Dy()/2)

		resultFile, _ := os.Create(fileName)
		defer resultFile.Close()
		resultPng := image.NewNRGBA(bigBound)
		draw.Draw(resultPng, bigBound, bigPng, image.ZP, draw.Src)
		draw.Draw(resultPng, bigBound.Add(offset), subImg, image.ZP, draw.Src)
		jpeg.Encode(resultFile, resultPng, &jpeg.Options{100})

		return fileName, nil
	}
	return fileName, nil
}
