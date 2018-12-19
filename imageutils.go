package iamgeutils

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)
var verbosity int

func readImage(inputfile string) (image.Image, error) {
	if verbosity > 1 {
		fmt.Printf("Reading file %s...\n", inputfile)
	}
	reader, err := os.Open(inputfile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	if verbosity > 1 {
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

func convert(inputfile, ctype string, quality, numColours int) {
	ctype = strings.Trim(ctype, ".")
	outfilename := inputfile + "." + ctype
	m, err := readImage(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	if verbosity > 1 {
		fmt.Printf("Creating output file %s...\n", outfilename)
	}
	
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	
	ct := strings.ToLower(ctype)
	if verbosity > 1 {
		fmt.Printf("Writing image to %s format...\n", ct)
	}
	if ct == "jpg" || ct == "jpeg" {
		err = jpeg.Encode(outfile, m, &jpeg.Options{Quality: quality})
	} else if ct == "png" {
		err = png.Encode(outfile, m)
	} else if ct == "gif" {
		err = gif.Encode(outfile, m, &gif.Options{NumColors: numColours})
	}
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
	quality := 100
	numColours := 256
	var err error
	usage := fmt.Sprintf("Usage: %s <command> <options>\n", os.Args[0])
	usage += fmt.Sprintf("Commands:\n")
	usage += fmt.Sprintf("convert                     Converts an image to another format.\n")
	usage += fmt.Sprintf("Options:\n")
	usage += fmt.Sprintf("-i --infile <inputfile>     Specifies an input file.\n")
	usage += fmt.Sprintf("-f --format <format>        Used with conversion commands. Used to specify a file format to convert to. Supported formats are png, jpeg, gif.\n")
	usage += fmt.Sprintf("-q --quality <value>		  Used when converting to jpeg. Specifies conversion quality in the range 1..100 inclusive. Higher is better. Default is 100.\n")
	usage += fmt.Sprintf("-n --num-colours <value>	  Used when converting to gif. Specifies the maximum number of colours used in the image in the range 1..256 inclusive. Default is 256.\n")
	usage += fmt.Sprintf("-v --verbose                Verbose mode.\n")
	usage += fmt.Sprintf("-h --help -?                Display this help information.\n")
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
		} else if arg == "-q" || arg == "--quality" {
			if (i + 1) < len(os.Args) {
				i++
				quality, err = strconv.Atoi(os.Args[i])
				if err != nil {
					log.Fatal(fmt.Sprintf("Error: Unable to parse quality integer: %s", err))
				}
			} else {
				log.Fatal(fmt.Sprintf("Error: %s switch provided without a quality value\n", arg))
			}
		} else if arg == "-n" || arg == "--num-colours" {
			if (i + 1) < len(os.Args) {
				i++
				numColours, err = strconv.Atoi(os.Args[i])
				if err != nil {
					log.Fatal(fmt.Sprintf("Error: Unable to parse number of colours integer: %s", err))
				}
			} else {
				log.Fatal(fmt.Sprintf("Error: %s switch provided without a quality value\n", arg))
			}
		} else if arg == "-v" || arg == "--verbose" {
			verbosity++
		} else if arg == "-h" || arg == "--help" || arg == "-?" {
			fmt.Println(usage)
			os.Exit(0)
		} else {
			log.Fatal(fmt.Sprintf("Error: Unable to parse argument %s", arg))
		}
	}
	if inputfile == "" {
		log.Fatal("Error: no input file provided.")
	}
	
	command := strings.ToLower(os.Args[1])
	if verbosity > 1 {
		fmt.Printf("Command=%s\nFile=%s\n", command, inputfile)
	}
	
	if command == "convert" {
		convert(inputfile, conversiontype, quality, numColours)
	} else if command == "jpg2png" || command == "jpeg2png" || command == "jpgtopng" || command == "jpegtopng" || command == "gif2png" || command == "giftopng" {
		convert(inputfile, "png", quality, numColours)
	} else if command == "png2jpg" || command == "png2jpeg" || command == "pngtojpg" || command == "pngtojpeg" || command == "gif2jpg" || command == "gif2jpeg" || command == "giftojpg" || command == "giftojpeg" {
		convert(inputfile, "jpeg", quality, numColours)
	} else if command == "jpg2gif" || command == "jpeg2gif" || command == "jpgtogif" || command == "jpegtogif" || command == "png2gif" || command == "pngtogif" {
		convert(inputfile, "gif", quality, numColours)
	} else {
		log.Fatal(fmt.Sprintf("Error: unknown command: %s", command))
	}
}