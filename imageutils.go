package main

import (
	"fmt"
	//"image"
	"log"
	"os"
	"strings"
)

func main() {
	inputfile := ""
	verbosity := 1
	usage := fmt.Sprintf("Usage: %s <command> -i <inputfile>", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-i" {
			if (i + 1) < len(os.Args) {
				i++
				inputfile = os.Args[i]
			} else {
				log.Fatal(fmt.Sprintf("Error: -i switch provided without an input file\n%s", usage))
			}
		} else if arg == "-v" || arg == "--verbose" {
			verbosity++
		} else if arg == "-h" || arg == "--help" || arg == "-?" {
			fmt.Println(usage)
			os.Exit(0)
		}
	}
	if inputfile == "" && len(os.Args) == 3 {
		inputfile = os.Args[2]
		if verbosity > 1 {
			fmt.Println("-i switch not provided. Using third argument as filename.")
		}
	}
	if inputfile == "" {
		log.Fatal("Error: no input file provided.")
	}
	
	command := strings.ToLower(os.Args[1])
	if verbosity > 1 {
		fmt.Printf("Command=%s\nFile=%s\n", command, inputfile)
	}
}