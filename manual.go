package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Tmpl is a basic man page[-ish] looking template
const ManualPageTemplate = `
{{define "manual"}}
NAME
  {{.BinName}} - what does this program do?

SYNOPSIS
  $ {{.BinName}}
  $ {{.BinName}} [-h|help]

DESCRIPTION
  {{.BinName}} what does this program do, but in more detail.

EXAMPLES
  $ {{.BinName}} -h

OPTIONS
{{.Options}}
VERSION
  version:  {{.BuildVersion}}
  compiled: {{.CompiledBy}}
  built:    {{.BuildTimestamp}}

{{end}}
`

// these vars are built at compile time, DO NOT ALTER
var (
	// Version adds build information
	BinName string
	// Version adds build information
	BuildVersion string
	// BuildTimestamp adds build information
	BuildTimestamp string
	// CompiledBy adds the make/model that was used to compile
	CompiledBy string
)

// Usage wraps a set of `Info` and creates a flag.Usage func
func RenderManualPage() func() {
	t := template.Must(template.New("manual").Parse(ManualPageTemplate))

	return func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		t.Execute(os.Stdout, struct {
			BinName        string
			Options        string
			BuildVersion   string
			BuildTimestamp string
			CompiledBy     string
		}{
			Options:        def.String(),
			BinName:        BinName,
			BuildVersion:   BuildVersion,
			BuildTimestamp: BuildTimestamp,
			CompiledBy:     CompiledBy,
		})
	}
}

func RenderManualPageMulti(flags []*flag.FlagSet) func() {
	t := template.Must(template.New("manual").Parse(ManualPageTemplate))

	return func() {
		var def bytes.Buffer
		for _, f := range flags {
			fmt.Fprintf(&def, "\nSUBCOMMAND: %s\n", f.Name())
			f.SetOutput(&def)
			f.PrintDefaults()
		}

		t.Execute(os.Stdout, struct {
			Options        string
			BinName        string
			BuildVersion   string
			BuildTimestamp string
			CompiledBy     string
		}{
			Options:        def.String(),
			BinName:        BinName,
			BuildVersion:   BuildVersion,
			BuildTimestamp: BuildTimestamp,
			CompiledBy:     CompiledBy,
		})
	}
}

func GetBuildInfo() string {
	t := template.Must(template.New("buildinfo").Parse(
		`{{define "buildinfo"}}{{.BinName}} version {{.BuildVersion}}{{end}}`,
	))

	var s strings.Builder
	t.Execute(&s, struct {
		BinName      string
		BuildVersion string
	}{
		BinName:      BinName,
		BuildVersion: BuildVersion,
	})

	return s.String()
}
