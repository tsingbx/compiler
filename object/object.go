package object

import (
	"fmt"

	"github.com/tsingbx/compiler/code"
	"github.com/tsingbx/interpreter/object"
)

const (
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"
)

type CompiledFunction struct {
	Instructions code.Instructions
	NumLocals    int
}

func (cf *CompiledFunction) Type() object.ObjectType {
	return COMPILED_FUNCTION_OBJ
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
