package main

import (
	"fmt"
	"os"

	"github.com/M-Quadra/go-python3-submodule/py"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
)

func init() {
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

func main() {

}
