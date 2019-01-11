package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

type ConvertCommand struct {
	Format 			string		`short:"f" long:"format" description:"Specifies a file format to convert to. Supported formats are png, jpeg, gif." required:"true"`
	InFiles			[]string	`short:"i" long:"infile" description:"Specifies an input file. Can be used multiple times." required:"true"`
	Quality 		int			`short:"q" long:"quality" description:"Used when converting to jpeg. Specifies conversion quality in the range 1..100 inclusive. Higher is better." default:"100"`
	NumColours		int			`short:"n" long:"num-colours" description:"Used when converting to gif. Specifies the maximum number of colours used in the image in the range 1..256 inclusive." default:"256"`
}

func readImage(inputfile string) (image.Image, error) {
	if verbosity() > 1 {
		fmt.Printf("Reading file %s...\n", inputfile)
	}
	reader, err := os.Open(inputfile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	if verbosity() > 1 {
		fmt.Printf("Decoding image from file...\n")
	}
	ext := strings.ToLower(path.Ext(inputfile))
	var m image.Image
	if ext == ".png" {
		m, err = png.Decode(reader)
	} else if ext == ".jpg" || ext == ".jpeg" {
		m, err = jpeg.Decode(reader)
	} else if ext == ".gif" {
		m, err = gif.Decode(reader)
	} else {
		m, _, err = image.Decode(reader)
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (Command *ConvertCommand) Execute(args []string) error {
	for _, inputfile := range Command.InFiles {
		ctype := strings.Trim(Command.Format, ".")
		outfilename := inputfile + "." + ctype
		m, err := readImage(inputfile)
		if err != nil {
			return err
		}
		if verbosity() > 1 {
			fmt.Printf("Creating output file %s...\n", outfilename)
		}
		
		outfile, err := os.Create(outfilename)
		if err != nil {
			return err
		}
		defer outfile.Close()
		
		ct := strings.ToLower(ctype)
		if verbosity() > 1 {
			fmt.Printf("Writing image to %s format...\n", ct)
		}
		if ct == "jpg" || ct == "jpeg" {
			err = jpeg.Encode(outfile, m, &jpeg.Options{Quality: Command.Quality})
		} else if ct == "png" {
			err = png.Encode(outfile, m)
		} else if ct == "gif" {
			err = gif.Encode(outfile, m, &gif.Options{NumColors: Command.NumColours})
		}
		if err != nil {
			return err
		}
	}
	return nil
}

