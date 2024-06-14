package main

import (
	"flag"
	"fmt"
	"os"
)

type generalParams struct {
	version bool
}

type allParams struct {
	general generalParams
}

func init() {
	params := allParams{}

	flag.BoolVar(&params.general.version, "version", false, "output version")
	flag.Usage = RenderManualPage()
	flag.Parse()

	if params.general.version {
		fmt.Fprintln(os.Stdout, GetVersionString())
		os.Exit(0)
	}
}
