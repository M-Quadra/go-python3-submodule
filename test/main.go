package main

import (
	"os"

	"github.com/M-Quadra/go-python3-submodule/py"
)

func init() {
	py.Initialize()
	// defer py.Finalize()
	if !py.IsInitialized() {
		os.Exit(-1)
	}
}

func main() {

}
