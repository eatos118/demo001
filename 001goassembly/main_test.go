package main

import (
	"testing"

	"github.com/eatos118/demo001/001goassembly/asmconst"
	"github.com/eatos118/demo001/001goassembly/asmfunc"
)

func TestMain(t *testing.T) {
	println(asmconst.TestInt)

	println(asmconst.Hello)

	var i int32 = 9
	asmfunc.CompareAndSwapInt32(&i, 9, 10)
	t.Log(i)
}
