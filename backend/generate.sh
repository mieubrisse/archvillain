#!/bin/bash
set -euo pipefail
script_dirpath="$(cd "$(dirname "${0}")" && pwd)"

cd "${script_dirpath}"

echo "Generating Go code from OpenAPI specification..."
oapi-codegen -generate types -package main openapi.yaml > generated_types.go
oapi-codegen -generate gorilla -package main openapi.yaml > generated_server.go
echo "Code generation complete!"