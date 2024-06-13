package main

import (
	"fmt"
	"strings"
)

// these vars are built at compile time, DO NOT ALTER
var (
	// Version adds build information
	binName string
	// Version adds build information
	buildVersion string
	// BuildTimestamp adds build information
	buildTimestamp string
	// CompiledBy adds the make/model that was used to compile
	compiledBy string
)

type buildInfo struct {
	BinName        string
	BuildVersion   string
	BuildTimestamp string
	CompiledBy     string
}

func getBuildInfo() buildInfo {
	if binName == "" {
		binName = `BIN_NAME`
	}

	bi := buildInfo{
		BinName:        binName,
		BuildVersion:   buildVersion,
		BuildTimestamp: buildTimestamp,
		CompiledBy:     compiledBy,
	}

	return bi
}

func (bi buildInfo) getBinName() string {
	return bi.BinName
}

func (bi buildInfo) getBuildVersion() string {
	return bi.BuildVersion
}

func (bi buildInfo) getBuildTimestamp() string {
	return bi.BuildTimestamp
}

func (bi buildInfo) getCompiledBy() string {
	return bi.CompiledBy
}

func (bi buildInfo) String() string {
	var rtn strings.Builder
	fmt.Fprintf(&rtn, "bin: %s\n", bi.BinName)
	fmt.Fprintf(&rtn, "build version: %s\n", bi.BuildVersion)
	fmt.Fprintf(&rtn, "build timestamp: %s\n", bi.BuildTimestamp)
	fmt.Fprintf(&rtn, "compiled by: %s\n", bi.CompiledBy)
	return rtn.String()
}
