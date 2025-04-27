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
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %w", err)
	}

	_, err = fileWriter.Write(data)
	if err != nil {
		return "", fmt.Errorf("error writing file data: %w", err)
	}

	err = bodyWriter.Close()
	if err != nil {
		return "", fmt.Errorf("error closing multipart writer: %w", err)
	}

	url := fmt.Sprintf(_URL_TMPl, method)
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return "", fmt.Errorf("%s", string(respBody))
	}
	return string(respBody), nil
}
