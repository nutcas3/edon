package runtime

import "flag"

func runtime () {
  // Parse command line flags
  evalScript := flag.String("eval", "", "Script to evaluate")
  flag.Parse()

  // Initialize the runtime
  rt, err = runtime.New()
  if err != nil {
    log.Fatal(err)
  }
  defer rt.Close()

  // if script provided, run it
  if *evalScript != "" {
    if err := rt.Eval(*evalScript); err != nil {
      fmt.Fprintf(os.Stderr, "Error: %v\n", err)
      os.Exit(1)
    }
    return
  }

  // if file provided as argument, execute it
  if len(flag.Args()) > 0 {
    if err := rt.ExecuteFile(flags.Args()[0]); err != nil {
      fmt.Fprintf(os.Stderr, "Error: %v\n", err)
      os.Exit(2)
    }
    return
  }

  // otherwise start REPL
   if err := rt.StartRepl(); err != nil {
    fmt.printF(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
   } 
 }

