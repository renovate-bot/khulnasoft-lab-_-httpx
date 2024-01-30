package main

import (
	"fmt"
	"log"

	"github.com/khulnasoft-lab/goflags"
	"github.com/khulnasoft-lab/gologger"
	"github.com/khulnasoft-lab/gologger/levels"
	"github.com/khulnasoft-lab/httpx/runner"
)

func main() {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)

	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: goflags.StringSlice{"scanme.sh", "khulnasoft.com", "localhost"},
		//InputFile: "./targetDomains.txt", // path to file containing the target domains list
		OnResult: func(r runner.Result) {
			// handle error
			if r.Err != nil {
				fmt.Printf("[Err] %s: %s\n", r.Input, r.Err)
				return
			}
			fmt.Printf("%s %s %d\n", r.Input, r.Host, r.StatusCode)
		},
	}

	if err := options.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()
}
