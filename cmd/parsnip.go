package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jamiekieranmartin/parsnip"
)

const cliVersion = "0.0.1"

const helpMessage = `
Parsnip is a minimal, powerful, text to JSON converter.
	parsnip v%s

By default, parsnip converts input text to JSON based on a given expression
	parsnip "(\S+) (\S+)" "Jamie Martin"
	
	{"1":"Jamie","2":"Martin"}

Named groups can be used to map key-value pairs.
	parsnip "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
	
	{"first":"Jamie","last":"Martin"}

Write to file
	parsnip -out "./result.json" "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	out := flag.String("out", "", "Output file")

	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("parsnip v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	// collect all other non-parsed arguments from the CLI as files to be run
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("error: 2 arguments expected")
		return
	}

	// parse input given expression
	parsed, err := parsnip.Parse(args[0], args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	// translate to json
	jsoned, err := json.MarshalIndent(parsed, "", "	")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	// get file path
	file, err := filepath.Abs(*out)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	if file != "" {
		// write to file
		err = ioutil.WriteFile(file, jsoned, os.ModePerm)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}

		fmt.Printf("written to %s\n", file)
	} else {
		// print to stdout
		fmt.Println(string(jsoned))
	}
}
