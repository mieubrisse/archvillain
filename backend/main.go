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

type Response struct {
	Message string `json:"message"`
}

type ContainerResponse struct {
	ContainerID string `json:"container_id"`
	Output      string `json:"output"`
	Status      string `json:"status"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World! Archvillain backend is running.",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, stacktrace.Propagate(err, "Failed to encode response").Error(), http.StatusInternalServerError)
		return
	}
}

func launchContainerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	
	// Use Docker CLI directly since Go SDK is having compatibility issues
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "alpine:latest", "sh", "-c", "echo 'Hello World from container!'")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		response := ContainerResponse{
			ContainerID: "unknown",
			Output:      fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output)),
			Status:      "failed",
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ContainerResponse{
		ContainerID: "docker-cli-run",
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
	r := mux.NewRouter()
	
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/launch-container", launchContainerHandler).Methods("POST")
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(stacktrace.Propagate(err, "Failed to start server"))
	}
}