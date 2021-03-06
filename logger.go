// +build !windows

package main

import (
	"github.com/ArthurHlt/gominlog"
	"os"
	"log"
)

var logger *gominlog.MinLog = gominlog.NewMinLogWithWriter("cf-sync", gominlog.Linfo, true, log.Ldate | log.Ltime, os.Stdout)
