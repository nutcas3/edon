package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/katungi/edon/internals/runtime"
)

type EvalRequest struct {
	Code string `json:"code"`
}

type EvalResponse struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

func main() {
	// Create a new runtime instance
	rt, err := runtime.New()
	if err != nil {
		log.Fatalf("Failed to create runtime: %v", err)
	}
	defer rt.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "web/static/index.html")
			return
		}
		http.NotFound(w, r)
	})

	// Handle REPL evaluation
	http.HandleFunc("/eval", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req EvalRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create a string builder to capture output
		var output strings.Builder

		// Evaluate the code
		err := rt.Eval(req.Code)
		
		response := EvalResponse{}
		if err != nil {
			response.Error = err.Error()
		} else {
			response.Output = output.String()
			if response.Output == "" {
				response.Output = "undefined"
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
