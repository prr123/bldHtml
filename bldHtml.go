// bldHtml.go
// program that builds a webiste
// date 6/6/2022
// author: prr, azul software
// copyright: 2022 prr, azul software
//
// usage:
// ./bldHtml inp
//
// reads inpYaml/inp.yaml
// generates output file: output/inp.html
//
package main

import (
	"os"
	"fmt"
)

func main () {

	argNum := len(os.Args)

	inpfilNam := ""
	switch argNum {
	case 1:
		fmt.Println("no input file provided!")
		fmt.Printf("usage: %s \"input file name\" \n", os.Args[0])
		os.Exit(1)
	case 2:
		inpfilNam = os.Args[1]
	default:
		fmt.Println("too many arguments in cmd line!")
		fmt.Printf("usage: %s \"input file name\" \n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("input file: %s\n", inpfilNam)

	// create input file name and open file
	inpYamlNam := "inpYaml/" + inpfilNam + ".yaml"

	inpfil, err := os.Open(inpYamlNam)
	if err != nil {
		fmt.Printf("error os.Open: %v\n", err)
	}

	outFilNam := "output/" + inpfilNam + ".html"
	outfil, err := os.Create(outFilNam)
	if err != nil {
		fmt.Printf("error os.Create %v\n", err)
		os.Exit(1)
	}
	fmt.Println("*** success ***")
}
