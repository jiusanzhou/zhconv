package main

import (
	"flag"
	"fmt"
	"log"
)

type Options struct {
	m         int
	cmd       string
	inputs    []string
	outputDir string
}

var opts *Options

func init() {
	m := flag.Int("concurrent", 1, "Concurrent of process, then more the faster.")
	outputDir := flag.String("output-dir", "output_data", "Output dir for data.")

	flag.Parse()

	other := flag.Args()
	if len(other) <= 1 {
		fmt.Println(other)
		log.Fatalln("You must offer command and a file to convert.")
	}

	cmd := other[0]
	switch cmd {
	case "2s", "2t":
	default:
		log.Fatalln("command should only be `2s` or `2t`")
	}

	inputs := other[1:]

	opts = &Options{
		m:         *m,
		outputDir: *outputDir,

		cmd:    cmd,
		inputs: inputs,
	}

/*	if s, err := os.Stat(opts.cmd); err != nil && os.IsNotExist(err) {
		os.MkdirAll(opts.outputDir, os.ModeDir)
	} else {
		if !s.IsDir() {

			os.MkdirAll(opts.outputDir, os.ModeDir)
		}
	}*/
}
