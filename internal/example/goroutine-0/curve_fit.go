package main

import (
	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pyfloat "github.com/M-Quadra/go-python3-submodule/v10/py-float"
	pyimport "github.com/M-Quadra/go-python3-submodule/v10/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v10/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/v10/py-object"
)

// GetPopt python function
func GetPopt(trainX, trainY []int) []float64 {
	Lock()
	defer Unlock()

	trainXPy := pylist.FromInts(trainX)
	defer py.DecRef(trainXPy)
	trainYPy := pylist.FromInts(trainY)
	defer py.DecRef(trainYPy)

	curvefit := pyimport.ImportModule("curvefit")
	defer py.DecRef(curvefit)

	getPopt := pyobject.GetAttrString(curvefit, "getPopt")
	defer py.DecRef(getPopt)

	res := pyobject.CallFunctionObjArgs(getPopt, trainXPy, trainYPy)
	defer py.DecRef(res)

	opt := make([]float64, 0, pylist.Size(res))
	for i := 0; i < pylist.Size(res); i++ {
		item := pyfloat.AsFloat64(pylist.GetItem(res, i))
		opt = append(opt, item)
	}
	return opt
}
