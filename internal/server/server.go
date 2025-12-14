package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/katungi/edon/internal/runtime"
)

type EvalRequest struct {
	Code string `json:"code"`
}

type EvalResponse struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Server struct {
	rt   *runtime.Runtime
	port string
}

func New(staticDir string) (*Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rt, err := runtime.New()
	if err != nil {
		return nil, err
	}

	return &Server{
		rt:   rt,
		port: port,
	}, nil
}

func (s *Server) Close() {
	if s.rt != nil {
		s.rt.Close()
	}
}

func (s *Server) Start(staticDir string) error {
	// Serve static files
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, staticDir+"/index.html")
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

		// Capture stdout
		oldStdout := os.Stdout
		pipeReader, pipeWriter, _ := os.Pipe()
		os.Stdout = pipeWriter

		// Evaluate the code
		evalErr := s.rt.Eval(req.Code)

		// Read the output
		pipeWriter.Close()
		var buf strings.Builder
		io.Copy(&buf, pipeReader)
		os.Stdout = oldStdout
		pipeReader.Close()

		response := EvalResponse{}
		if evalErr != nil {
			response.Error = evalErr.Error()
		} else {
			output := buf.String()
			if output == "" {
				output = "=> " + req.Code
			}
			response.Output = output
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Printf("Server starting on http://localhost:%s\n", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}
