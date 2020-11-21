package main

import (
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("isNumber", js.FuncOf(isNumber))
	<-done
}

func isNumber(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "ERROR: number of arguments doesn't match"
	}

	return args[0].IsNaN
}
