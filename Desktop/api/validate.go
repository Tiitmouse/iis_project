package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const _URL_TMPl = "http://localhost:8088/upload/%s"

func Validate(filename string, data []byte, method string) (string, error) {
	// Create a buffer to hold the request body
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Create a form file field
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %w", err)
	}

	// Write the file data to the form file field
	_, err = fileWriter.Write(data)
	if err != nil {
		return "", fmt.Errorf("error writing file data: %w", err)
	}

	// Close the multipart writer to finalize the form
	err = bodyWriter.Close()
	if err != nil {
		return "", fmt.Errorf("error closing multipart writer: %w", err)
	}

	// Create a new request
	url := fmt.Sprintf(_URL_TMPl, method)
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set the content type to multipart/form-data
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return "", fmt.Errorf(string(respBody))
	}
	// Return the response as a string
	return string(respBody), nil
}
