package console

import (
	"fmt"
	"github.com/buke/quickjs-go"
)

func Init(ctx *quickjs.Context) error {
	global := ctx.Globals()
	console := ctx.Object()
	defer console.Free()

	// Implement console.log
	log := ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		for i, arg := range args {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(arg.String())
		}
		fmt.Println()
		return ctx.Undefined()
	})
	defer log.Free()

	console.Set("log", log)
	global.Set("console", console)

	return nil
}
