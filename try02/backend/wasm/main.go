package main

import (
	"syscall/js"

	"github.com/sky0621/try-wasm/try02/backend/domain"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("validateTodoText", js.FuncOf(validateTodoText))
	<-done
}

func validateTodoText(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "ERROR: number of arguments doesn't match"
	}

	if res := domain.ValidateTodoText(args[0].String()); res != nil {
		return res.Err.Error()
	}
	return nil
}
