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

	if errs := domain.ValidateTodoText(args[0].String()); errs != nil {
		var errMaps []map[string]interface{}
		for _, e := range errs {
			errMaps = append(errMaps, map[string]interface{}{
				"status_code":   400,
				"error_message": e.Err.Error(),
				"field":         e.Field,
				"value":         e.Value,
			})
		}
		return errMaps
	}
	return nil
}
