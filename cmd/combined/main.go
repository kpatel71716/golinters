package main

import (
	"flag"

	"github.com/mongodb-forks/golinters/deferfor"
	"github.com/mongodb-forks/golinters/errresp"
	"github.com/mongodb-forks/golinters/mustcheck"
	"github.com/mongodb-forks/golinters/printf"
	"github.com/mongodb-forks/golinters/println"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func init() {
	// go vet always adds this flag for certain packages in the standard library,
	// which causes "flag provided but not defined" errors when running with
	// custom vet tools.
	// So we just declare it here and swallow the flag.
	// See https://github.com/golang/go/issues/34053 for details.
	// TODO: Remove this once above issue is resolved.
	flag.String("unsafeptr", "", "")
}

func main() {
	unitchecker.Main(
		println.Analyzer,
		printf.Analyzer,
		mustcheck.Analyzer,
		deferfor.Analyzer,
		errresp.Analyzer,
	)
}
