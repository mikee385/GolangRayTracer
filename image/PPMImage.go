package image

import (
	"fmt"
	"github.com/mikee385/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/table"
	"math"
	"os"
)

type PPMImage struct {
	fileName string
}

func NewPPMImage(fileName string) PPMImage {
	return PPMImage{
		fileName: fileName,
	}
}

func (image *PPMImage) Save(pixelTable *table.Table) error {
	file, err := os.OpenFile(image.fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var width = pixelTable.Width()
	var height = pixelTable.Height()

	_, err = file.WriteString(fmt.Sprintf("P6\n%v %v\n255\n", width, height))
	if err != nil {
		return err
	}

	var colorBytes = make([]byte, 3)
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			var pixel = pixelTable.Get(row, column).(color.ColorRGB)
			colorBytes[0] = convertToByte(pixel.Red)
			colorBytes[1] = convertToByte(pixel.Green)
			colorBytes[2] = convertToByte(pixel.Blue)

			_, err = file.Write(colorBytes)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func convertToByte(value float32) byte {
	return byte(math.Max(math.Min(float64(value)*255, 255), 0))
}
