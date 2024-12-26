package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type userRequest struct {
	Prompt string `json:"prompt"`
}

type ApiResponse struct {
	Response string `json:"response"`
}

func main() {
	router := gin.Default()
	router.POST("/request", postMessage)
	router.Run("localhost:8081")
}

func getOllamaResponse(model, prompt string) (response string) {

	url := "http://localhost:11434/api/generate"

	ollamaRequest := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(
		ollamaRequest,
	)
	if err != nil {
		return
	}
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the JSON response into the struct
	var apiResp ApiResponse
	json.Unmarshal(body, &apiResp)
	return apiResp.Response
}

func postMessage(c *gin.Context) {
	var prompt userRequest
	if err := c.BindJSON(&prompt); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, getOllamaResponse("llama3.2:3b", prompt.Prompt))

}
