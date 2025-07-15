FROM mcr.microsoft.com/devcontainers/go:1.23

# Install additional tools needed for development
RUN apt-get update && apt-get install -y \
    curl \
    jq \
    && rm -rf /var/lib/apt/lists/*

# Install Docker CLI (will be provided by the docker-outside-of-docker feature)
# The feature handles Docker CLI installation and configuration

# Set working directory
WORKDIR /workspace

# Copy go module files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Make scripts executable
RUN chmod +x start.sh stop.sh

# Expose the backend port
EXPOSE 8080