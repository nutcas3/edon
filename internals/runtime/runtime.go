package runtime

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/buke/quickjs-go"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/katungi/edon/internals/modules/console"
)

type Runtime struct {
	jsRuntime *quickjs.Runtime
	context   *quickjs.Context
}

const (
	prompt      = ">> (ﾉ˚Д˚)ﾉ ⌒┻━┻"
	multiPrompt = "..."
)

// REPL specific errors
var (
	ErrInterrupt = fmt.Errorf("interrupted")
	ErrExit      = fmt.Errorf("exit")
)

func New() (*Runtime, error) {
	jsRuntime := quickjs.NewRuntime()
	context := jsRuntime.NewContext()

	r := &Runtime{
		jsRuntime: &jsRuntime,
		context:   context,
	}

	// Initialize built-in modules
	if err := r.initializeBuiltins(); err != nil {
		return nil, fmt.Errorf("failed to initialize builtins: %w", err)
	}

	return r, nil
}

func (r *Runtime) initializeBuiltins() error {
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

func (r *Runtime) ExecuteFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	_, err = r.context.Eval(string(data))
	return err
}

func (r *Runtime) Close() {
	r.context.Close()
	r.jsRuntime.Close()
}

func (r *Runtime) StartREPL() error {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:            prompt,
		HistoryFile:       "/tmp/edon_history",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		HistoryLimit:      500,
	})

	if err != nil {
		return fmt.Errorf("failed to create readline instance: %w", err)
	}
	defer rl.Close()

	printWelcome()

	// The Fucking Loop
	multiline := false
	var code strings.Builder

	for {
		currentPrompt := prompt
		if multiline {
			currentPrompt = multiPrompt
		}

		rl.SetPrompt(currentPrompt)

		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				if multiline {
					// Cancel multi-line input
					multiline = false
					code.Reset()
					continue
				}
				return ErrInterrupt
			} else if err == io.EOF {
				return ErrExit
			}
			return err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handle special commands
		if !multiline {
			switch line {
			case ".exit", "exit", "quit":
				return nil
			case ".help", "help":
				printHelp()
				continue
			case ".clear":
				code.Reset()
				multiline = false
				continue
			}
		}

		// Append the line to our code buffer
		code.WriteString(line)
		code.WriteString("\n")

		// check if we nned to contine reading more lines
		if isIncomplete(line) {
			multiline = true
			continue
		}
		// Execute the code
		result, err := r.context.Eval(code.String())
		if err != nil {
			color.Red("Error: %v", err)
		} else if !result.IsUndefined() {
			// Convert result to string and print
			val := result.String()
			if val != "undefined" {
				color.Green("=> %s", val)
			}
		}

		// Reset for next input
		code.Reset()
		multiline = false
	}
}

// isIncomplete checks if the input code block is incomplete and needs more lines
func isIncomplete(line string) bool {
    line = strings.TrimSpace(line)
    
    // Check for obvious continuation cases
    if strings.HasSuffix(line, "{") ||
       strings.HasSuffix(line, "\\") ||
       strings.HasSuffix(line, ".") {
        return true
    }

    // Count brackets/braces/parentheses
    brackets := 0
    braces := 0
    parens := 0

    for _, ch := range line {
        switch ch {
        case '[':
            brackets++
        case ']':
            brackets--
        case '{':
            braces++
        case '}':
            braces--
        case '(':
            parens++
        case ')':
            parens--
        }
    }

    return brackets > 0 || braces > 0 || parens > 0
}


// createCompleter creates an autocomplete handler
func createCompleter() *readline.PrefixCompleter {
    return readline.NewPrefixCompleter(
        readline.PcItem("console.log"),
        readline.PcItem("let"),
        readline.PcItem("const"),
        readline.PcItem("function"),
        readline.PcItem("return"),
        readline.PcItem("if"),
        readline.PcItem("else"),
        readline.PcItem("for"),
        readline.PcItem("while"),
        readline.PcItem(".help"),
        readline.PcItem(".exit"),
        readline.PcItem(".clear"),
    )
}

// printWelcome prints the REPL welcome message
func printWelcome() {
    color.Cyan("Welcome to Edon REPL!")
    color.Cyan("Type .help for more information")
    fmt.Println()
}

// printHelp prints the help information
func printHelp() {
    help := `
Commands:
  .help, help    Show this help message
  .exit, exit    Exit the REPL
  .clear         Clear the current input
  
Special Keys:
  Ctrl+C         Cancel current input / Exit REPL
  Ctrl+D         Exit REPL
  Up/Down        Navigate through history
  Tab            Auto-complete
`
    color.Yellow(help)
}