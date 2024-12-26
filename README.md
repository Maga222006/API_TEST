# API Test Project

A Go-based API service that integrates with Ollama for text generation.

## Features

- RESTful API using Gin web framework
- Integration with Ollama's local LLM service
- JSON request/response handling
- Streaming support (configured but not enabled by default)

## Prerequisites

- Go 1.x
- Ollama running locally on port 11434
- LLama 3.2 3B model installed in Ollama

## Installation

1. Clone the repository:
```bash
git clone git@github.com:Maga222006/API_TEST.git
cd API_TEST
```

2. Install dependencies:
```bash
go mod download
```

## Usage

1. Make sure Ollama is running locally with the llama3.2:3b model
2. Start the server:
```bash
go run main.go
```

The server will start on `localhost:8081`

### API Endpoints

#### POST /request
Accepts JSON requests with the following structure:
```json
{
    "prompt": "Your prompt text here"
}
```

Returns JSON response with generated text:
```json
{
    "response": "Generated response from Ollama"
}
```

## Configuration

- Default port: 8081
- Default Ollama endpoint: http://localhost:11434/api/generate
- Default model: llama3.2:3b 