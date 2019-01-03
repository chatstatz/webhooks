package main

import (
	"fmt"
	"log"
	"os"
)

// Log level prefixes
const (
	// WARN  = "[WARN] "
	DEBUG = "[DEBU] "
	INFO  = "[INFO] "
	ERROR = "[ERRO] "
	FATAL = "[FATA] "
)

func println(l string, v ...interface{}) {
	if l == "" {
		log.Println(fmt.Sprint(v...))
		return
	}

	log.Println(fmt.Sprintf("%s%s", l, fmt.Sprint(v...)))
}

func ldebug(v ...interface{}) {
	if verbose {
		println(DEBUG, v...)
	}
}

func ldebugf(f string, v ...interface{}) {
	if verbose {
		ldebug(fmt.Sprintf(f, v...))
	}
}

func linfo(v ...interface{}) {
	println(INFO, v...)
}

func linfof(f string, v ...interface{}) {
	linfo(fmt.Sprintf(f, v...))
}

func lerror(v ...interface{}) {
	println(ERROR, v...)
}

func lerrorf(f string, v ...interface{}) {
	lfatal(fmt.Sprintf(f, v...))
}

func lfatal(v ...interface{}) {
	println(FATAL, v...)
	os.Exit(1)
}

func lfatalf(f string, v ...interface{}) {
	lfatal(fmt.Sprintf(f, v...))
}
