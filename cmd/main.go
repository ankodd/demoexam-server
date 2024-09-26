package main

import (
	"github.com/ankodd/demoexam/core/internal/app"
	"github.com/ankodd/demoexam/core/internal/utils/sl"
	"os"
)

func main() {
	logger := sl.New()
	if err := app.Run(logger); err != nil {
		sl.Err(logger, err)
		os.Exit(1)
	}
}
