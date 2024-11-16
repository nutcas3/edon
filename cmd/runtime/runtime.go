package runtime

import (
	"fmt"
	"io/ioutil"

	"github.com/buke/quickjs-go"
)

type Runtime struct {
	jsRuntime *quickjs.Runtime
	context   *quickjs.Context
}

func New() (*Runtime, error) {
	jsRuntime := quickjs.NewRuntime()
	context := jsRuntime.NewContext()

	r := &Runtime{
		jsRuntime: &jsRuntime,
		context:   context,
	}

	// Initialize built-in modules and stuff
	if err := r.initializeBuiltins(); err != nil {
		return nil, fmt.Errorf("failed to initialize builtins: %w", err)
	}

	return r, nil
}

func (r  *Runtime) initializeBuiltins() error {
	// Add console module
	if err := console.Init(r.context); err != nil {
		return fmt.Errorf("failed to initialize console: %w", err)
	}
	return nil
}

func (r *Runtime) Eval(script string) error {
	_, err := r.context.Eval(script)
	return err
}

func (r *Runtime) ExecuteFile(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = r.context.Eval(string(data))
	return err
}


// func (r *Runtime) Close() {
//     r.context.Free()
//     r.jsRuntime.Free()
// }

func (r *Runtime) StartREPL() error {
    // TODO: Implement REPL
    return fmt.Errorf("REPL not implemented yet")
}