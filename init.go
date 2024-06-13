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
	bi := getBuildInfo()
	params := allParams{}

	flag.BoolVar(&params.general.version, "version", false, "output version/build information")
	flag.Usage = Usage(Info{
		Bin:            bi.getBinName(),
		Version:        bi.getBuildVersion(),
		CompiledBy:     bi.getCompiledBy(),
		BuildTimestamp: bi.getBuildTimestamp(),
	})
	flag.Parse()

	if params.general.version {
		fmt.Fprint(os.Stdout, getBuildInfo().String())
		os.Exit(0)
	}
}
