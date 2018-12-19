package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strings"
)
var verbosity int

// Converts an image to jpeg format.
// inputfile: File to be converted.
// quality: Conversion quality from 1 to 100 inclusive. Higher is better.
func convertToJpeg(inputfile string, quality int) {
	if verbosity > 1 {
		fmt.Printf("Reading file %s...\n", inputfile)
	}
	reader, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	if verbosity > 1 {
		fmt.Printf("Decoding image from file...\n")
	}
	var m image.Image
	if strings.ToLower(path.Ext(inputfile)) == ".png" {
		m, err = png.Decode(reader)
	} else {
		m, _, err = image.Decode(reader)
	}
	if err != nil {
		log.Fatal(err)
	}
	outfilename := inputfile + ".jpeg"
	if verbosity > 1 {
		fmt.Printf("Creating output file %s...\n", outfilename)
	}
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	
	if verbosity > 1 {
		fmt.Printf("Writing image to %s...\n", outfilename)
	}
	err = jpeg.Encode(outfile, m, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Fatal(err)
	}
}

// Converts an image to png.
func convertToPng(inputfile string) {
	log.Fatal("TODO: Not yet implemented.")
}

func main() {
	verbosity = 1
	conversiontype := ""
	inputfile := ""
	usage := fmt.Sprintf("Usage: %s <command> <options>", os.Args[0])
	usage += fmt.Sprintf("Commands:\n")
	usage += fmt.Sprintf("convert                     Converts an image to another format.")
	usage += fmt.Sprintf("Options:\n")
	usage += fmt.Sprintf("-i --infile <inputfile>     Specifies an input file.")
	usage += fmt.Sprintf("-f --format <format>        Used with conversion commands. Used to specify a file format to convert to. e.g. png, .jpg")
	usage += fmt.Sprintf("-v --verbose                Verbose mode.")
	usage += fmt.Sprintf("-h --help -?                Display this help information.")
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-i" || arg == "--infile" {
			if (i + 1) < len(os.Args) {
				i++
				inputfile = os.Args[i]
			} else {
				log.Fatal(fmt.Sprintf("Error: -i switch provided without an input file\n%s", usage))
			}
		} else if arg == "-f" || arg == "--format" {
			if (i + 1) < len(os.Args) {
				i++
				conversiontype = os.Args[i]
			} else {
				log.Fatal(fmt.Sprintf("Error: %s switch provided without a format\n%s", usage))
			}
		} else if arg == "-v" || arg == "--verbose" {
			verbosity++
		} else if arg == "-h" || arg == "--help" || arg == "-?" {
			fmt.Println(usage)
			os.Exit(0)
		}
	}
	if inputfile == "" {
		log.Fatal("Error: no input file provided.")
	}
	
	command := strings.ToLower(os.Args[1])
	if verbosity > 1 {
		fmt.Printf("Command=%s\nFile=%s\n", command, inputfile)
	}
	ct := strings.ToLower(conversiontype)
	if command == "jpg2png" || command == "jpeg2png" || command == "jpgtopng" || command == "jpegtopng" || (command == "convert" && strings.ToLower(conversiontype) == "png") {
		convertToPng(inputfile)
	} else if command == "png2jpg" || command == "png2jpeg" || (command == "convert" && (ct == "jpg" || ct == "jpeg" || ct == ".jpg" || ct == ".jpeg")) {
		convertToJpeg(inputfile, 100)
	} else {
		log.Fatal(fmt.Sprintf("Error: unknown command: %s", command))
	}
}