package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kurtosis-tech/stacktrace"
)

// Server implements the generated ServerInterface
type Server struct{}

// GetHello implements the GET /hello endpoint
func (s *Server) GetHello(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World! Archvillain backend is running.",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, stacktrace.Propagate(err, "Failed to encode response").Error(), http.StatusInternalServerError)
		return
	}
}

// LaunchContainer implements the POST /launch-container endpoint
func (s *Server) LaunchContainer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	
	// Use Docker CLI directly since Go SDK is having compatibility issues
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "alpine:latest", "sh", "-c", "echo 'Hello World from container!'")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		response := ContainerResponse{
			ContainerId: "unknown",
			Output:      fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output)),
			Status:      "failed",
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ContainerResponse{
		ContainerId: "docker-cli-run",
		Output:      strings.TrimSpace(string(output)),
		Status:      "completed",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, stacktrace.Propagate(err, "Failed to encode response").Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	server := &Server{}
	r := mux.NewRouter()
	
	// Use the generated HandlerFromMux function to register routes
	HandlerFromMux(server, r)
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(stacktrace.Propagate(err, "Failed to start server"))
	}
}