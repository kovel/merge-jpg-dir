package main

import (
	"flag"
	gim "github.com/ozankasikci/go-image-merge"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var imageDir string
	var outputFile string
	flag.StringVar(&imageDir, "dir", "./images", "");
	flag.StringVar(&outputFile, "out", "./output.jpg", "");
	flag.Parse()

	files, err := ioutil.ReadDir(imageDir)
	if err != nil {
		log.Fatal(err)
	}

	var grids = make([]*gim.Grid, 0)
	for _, f := range files {
		log.Println(f.Name())
		grids = append(grids, &gim.Grid{
			ImageFilePath: path.Join("images", f.Name()),
		})
	}

	// accepts *Grid instances, grid unit count x, grid unit count y
	// returns an *image.RGBA object
	rgba, err := gim.New(grids, 1, len(grids)).Merge(); if err != nil {
		log.Fatal(err)
	}

	// save the output to jpg or png
	file, err := os.Create(outputFile)
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
	err = png.Encode(file, rgba)
}
