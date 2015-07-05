// Copyright 2012 - 2015 The ASCIIToSVG Contributors
// All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asciitosvg/asciitosvg"
)

func mainImpl() error {
	in := flag.String("i", "-", "Path to input text file. If set to \"-\" (hyphen), stdin is used (default).")
	out := flag.String("o", "-", "Path to output SVG file. If set to \"-\" (hyphen), stdout is used (default).")
	//noBlur := flag.Bool("b", false, "disable drop-shadow blur")
	//font := flag.String("f", "", "font family to use")
	//scaleX := flag.Int("x", 9, "X grid scale in pixels. Default to 9.")
	//scaleY := flag.Int("y", 16, "Y grid scale in pixels. Default to 16.")
	flag.Parse()

	var input []byte
	var err error
	if *in == "-" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(*in)
	}
	if err != nil {
		return err
	}

	canvas := asciitosvg.NewCanvas(input)
	boxes := canvas.FindBoxes()
	svg := boxes.ToSVG()
	if *out == "-" {
		_, err := os.Stdout.Write(svg)
		return err
	}
	return ioutil.WriteFile(*out, svg, 0666)
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "a2s: %s\n", err)
		os.Exit(1)
	}
}
