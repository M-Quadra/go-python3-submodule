package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pylist "github.com/M-Quadra/go-python3-submodule/v11/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/v11/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v11/py-unicode"
)

func init() {
	py.Finalize()
	py.Initialize()
	// defer py.Finalize()
	if !py.IsInitialized() {
		os.Exit(-1)
	}

	paths := pysys.GetObject("path")
	if paths == nil {
		os.Exit(-1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	wdUnicode := pyunicode.FromString(wd)
	defer py.DecRef(wdUnicode)
	if !pylist.Append(paths, wdUnicode) {
		os.Exit(-1)
	}
}

var _rand = rand.New(rand.NewSource(114514))

func main() {
}
