# HALO Runtime

Basic implementation of a Deno like runtime in GO based on QuickJS

Where are we now?

- we just implemented the runtime, and the repl

What are we working on next?

- Implement a module cache-ing system
- Implement the module parsing system to parse url imports
- Implement the module resolution system to resolve url imports
- Implement the module loading system to load modules

How to get started?

Currently we have a REPL that is pretty much ready to go, but we need to
implement the module loading system first for it to work.

```bash
go run cmd/edon/main.go
```

That should run the repl and allow you to run code in the repl.
