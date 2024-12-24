package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

var matrix [256][256]int

func saveImg(filePath string, maxValue int) {
	img := image.NewGray(image.Rect(0, 0, 256, 256))

	maxLogValue := math.Log(float64(1 + maxValue))
	for idx := 0; idx < 256; idx++ {
		for idy := 0; idy < 256; idy++ {
			var grayValue uint8
			if maxValue > 0 {
				logValue := math.Log(float64(1 + matrix[idx][idy]))

				grayValue = uint8(logValue * 255 / maxLogValue)
			} else {
				grayValue = 255
			}

			img.Set(idx, idy, color.Gray{Y: grayValue})
		}
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error while creating a file: %v.", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatalf("Error while encoding image: %v", err)
		return
	}

	log.Printf("Image saved into %v", filePath)
}

func readFile(inputFileName string) []byte {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("Error while reading file %v. %v", inputFileName, err)
		return nil
	}

	return content
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage:\n\tgo run main.go <input_file> <output_file>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	content := readFile(inputFileName)
	if content == nil || len(content) < 1 {
		return
	}

	var maxMatrixValue int
	for idx := 0; idx < len(content)-1; idx++ {
		a, b := content[idx], content[idx+1]
		matrix[a][b] += 1
		maxMatrixValue = max(matrix[a][b], maxMatrixValue)
	}

	saveImg(outputFileName, maxMatrixValue)

	fmt.Printf("File name: %v; File size: %v", inputFileName, len(content))
}
