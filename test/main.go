package main

import (
	"os"
	"python/py"
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
