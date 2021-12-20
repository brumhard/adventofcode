package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brumhard/adventofcode/aocconv"
)

//go:embed input.txt
var inputFile string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	input, err := inputFromString(string(inputFile))
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

type input struct {
	enhancementString []string
	image             [][]string
}

func inputFromString(inputStr string) (input, error) {
	parts := aocconv.StrToStrSlice(inputStr, aocconv.WithDelimeter("\n\n"))
	enhancementString := aocconv.StrToStrSlice(parts[0], aocconv.WithDelimeter(""))

	var image [][]string
	for _, line := range aocconv.StrToStrSlice(parts[1]) {
		image = append(image, aocconv.StrToStrSlice(line, aocconv.WithDelimeter("")))
	}

	return input{
		enhancementString: enhancementString,
		image:             image,
	}, nil
}

func padImage(image [][]string, padding int) [][]string {
	newImage := make([][]string, len(image)+2*padding)
	for i := range newImage {
		newImage[i] = make([]string, len(image[0])+2*padding)
		for y := range newImage[i] {
			if y-padding < 0 || i-padding < 0 {
				newImage[i][y] = "."
				continue
			}
			if y > len(image)-1+padding || i > len(image)-1+padding {
				newImage[i][y] = "."
				continue
			}
			newImage[i][y] = image[i-padding][y-padding]
		}
	}

	return newImage
}

func enhanceImage(p [][]string, enhancementString []string) [][]string {
	var enhancedImage [][]string
	for i := range p {
		var enhancedImageLine []string
		for y := range p[i] {
			if y < 1 || i < 1 || i >= len(p)-2 || y >= len(p[i])-2 {
				enhancedImageLine = append(enhancedImageLine, enhancementString[getNumFromPixels(strings.Repeat(p[i][y], 9))])
				continue
			}
			binString := p[i-1][y-1] + p[i-1][y] + p[i-1][y+1] + p[i][y-1] + p[i][y] + p[i][y+1] + p[i+1][y-1] + p[i+1][y] + p[i+1][y+1]
			index := getNumFromPixels(binString)
			enhancedImageLine = append(enhancedImageLine, enhancementString[index])
		}
		enhancedImage = append(enhancedImage, enhancedImageLine)
	}

	return enhancedImage
}

func getNumFromPixels(pixelString string) int {
	bitString := strings.NewReplacer(".", "0", "#", "1").Replace(pixelString)
	num, err := strconv.ParseInt(bitString, 2, 0)
	if err != nil {
		panic("invalid number")
	}

	return int(num)
}

func SolvePart1(input input) int {
	currentImage := input.image
	currentImage = padImage(currentImage, 10)

	for i := 0; i < 2; i++ {
		currentImage = enhanceImage(currentImage, input.enhancementString)
	}

	litPixels := 0
	for i := range currentImage {
		fmt.Println()
		for y := range currentImage[i] {
			fmt.Print(currentImage[i][y])
			if currentImage[i][y] == "#" {
				litPixels++
			}
		}
	}

	return litPixels
}

func SolvePart2(input input) int {
	currentImage := input.image
	currentImage = padImage(currentImage, 100)

	for i := 0; i < 50; i++ {
		currentImage = enhanceImage(currentImage, input.enhancementString)
	}

	litPixels := 0
	for i := range currentImage {
		fmt.Println()
		for y := range currentImage[i] {
			fmt.Print(currentImage[i][y])
			if currentImage[i][y] == "#" {
				litPixels++
			}
		}
	}

	return litPixels
}
