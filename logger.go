package main

import (
	"log"
)

// var (
// 	stderr  = LogWriter{log.New(os.Stderr, "", log.Lshortfile)}
// 	stdnull = LogWriter{log.New(io.Discard, "", 0)}
// 	stdout  = LogWriter{log.New(os.Stdout, "", 0)}
// )

type LogWriter struct {
	*log.Logger
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	return lw.Writer().Write(p)
}
