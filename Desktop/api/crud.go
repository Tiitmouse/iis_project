package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Contact struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Value   string   `json:"value"`
	Name    string   `json:"name,omitempty"`
	Sources []string `json:"sources,omitempty"`
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Secure struct {
	accessToken  string
	refreshToken string
	httpClient   *http.Client
}

func NewSecure() *Secure {
	instance := &Secure{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	log.Printf("NewSecure: Instance created with address %p, httpClient: %+v", instance, instance.httpClient)
	return instance
}

func (s *Secure) doRequest(method, url string, body []byte) (*http.Response, error) {
	log.Printf("doRequest: Entered. Method: %s, URL: %s. Instance: %p", method, url, s)

	var reqBodyReader *bytes.Reader
	if body != nil {
		reqBodyReader = bytes.NewReader(body)
	}

	log.Printf("doRequest: About to call http.NewRequest. Method: '%s', URL: '%s', reqBodyReader is nil: %t", method, url, reqBodyReader == nil)

	var req *http.Request
	var err error

	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, reqBodyReader)
	}

	if err != nil {
		log.Printf("Error creating new HTTP request in doRequest: %v. Method: %s, URL: %s", err, method, url)
		return nil, err
	}
	if req == nil {
		log.Printf("doRequest: http.NewRequest returned nil request AND nil error. This is unexpected. Method: %s, URL: %s", method, url)
		return nil, fmt.Errorf("http.NewRequest returned nil request without error for URL: %s", url)
	}

	log.Printf("Attempting to make request in doRequest. Method: %s, URL: %s", method, url)

	if s.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+s.accessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	if req == nil {
		log.Println("Request object (req) is nil before httpClient.Do")
		return nil, fmt.Errorf("internal error: request object is nil")
	}

	return s.httpClient.Do(req)
}

func (s *Secure) Login(username, password string) error {
	if s == nil {
		log.Println("Secure instance (s) is nil in Login")
		return fmt.Errorf("internal server error: secure context not initialized")
	}
	log.Printf("Login: Called on instance %p. Initial s.accessToken: '%s', s.httpClient: %+v", s, s.accessToken, s.httpClient)

	loginURL := "http://localhost:8088/api/login"

	credentials := map[string]string{
		"username": username,
		"password": password,
	}
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		return fmt.Errorf("failed to marshal login credentials: %w", err)
	}

	resp, err := s.doRequest(http.MethodPost, loginURL, jsonData)
	if err != nil {
		return fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status: %s", resp.Status)
	}

	var loginResp loginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return fmt.Errorf("failed to decode login response: %w", err)
	}

	s.accessToken = loginResp.AccessToken
	s.refreshToken = loginResp.RefreshToken
	log.Printf("Login: Completed for instance %p. New s.accessToken: '%s'", s, s.accessToken)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if s.refreshToken == "" {
					log.Println("No refresh token, stopping refresh attempts.")
					return
				}
				log.Println("Attempting token refresh...")
				s.refresh()
			}
		}
	}()
	return nil
}

func (s *Secure) Logout() error {
	if s == nil {
		log.Println("Secure instance (s) is nil in Logout")
		return fmt.Errorf("internal server error: secure context not initialized")
	}
	s.accessToken = ""
	s.refreshToken = ""
	return nil
}

func (s *Secure) refresh() {
	if s == nil {
		log.Println("Secure instance (s) is nil in refresh")
		return
	}
	refreshURL := "http://localhost:8088/api/refresh"

	if s.refreshToken == "" {
		log.Println("No refresh token available, cannot refresh session.")
		return
	}

	payload := map[string]string{
		"refresh_token": s.refreshToken,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal refresh token payload: %v", err)
		return
	}

	resp, err := s.doRequest(http.MethodPost, refreshURL, jsonData)
	if err != nil {
		log.Printf("Refresh token request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Refresh token failed with status: %s", resp.Status)
		s.accessToken = ""
		s.refreshToken = ""
		return
	}

	var refreshResp loginResponse
	if err := json.NewDecoder(resp.Body).Decode(&refreshResp); err != nil {
		log.Printf("Failed to decode refresh token response: %v", err)
		return
	}

	s.accessToken = refreshResp.AccessToken
	if refreshResp.RefreshToken != "" {
		s.refreshToken = refreshResp.RefreshToken
	}
	log.Println("Token refreshed successfully.")
}

const baseURL = "http://localhost:8088/api/contacts"

func (s *Secure) FetchContacts() ([]Contact, error) {
	resp, err := s.doRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Printf("! new error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("! new error: %v", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch contacts: %s", resp.Status)
	}

	var contacts []Contact
	fmt.Printf("resp.Body: %v\n", resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&contacts); err != nil {
		log.Printf("! new error: %v", err)
		return nil, err
	}
	return contacts, nil
}

func (s *Secure) DeleteContact(id string) error {
	if s == nil {
		log.Println("Secure instance (s) is nil in DeleteContact")
		return fmt.Errorf("internal server error: secure context not initialized")
	}
	url := fmt.Sprintf("%s/%s", baseURL, id)
	resp, err := s.doRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete contact: %s", resp.Status)
	}
	return nil
}

func (s *Secure) CreateContact(contact Contact) (Contact, error) {
	if s == nil {
		log.Println("Secure instance (s) is nil in CreateContact")
		return Contact{}, fmt.Errorf("internal server error: secure context not initialized")
	}
	jsonData, err := json.Marshal(contact)
	if err != nil {
		return Contact{}, err
	}

	resp, err := s.doRequest(http.MethodPost, baseURL, jsonData)
	if err != nil {
		return Contact{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return Contact{}, fmt.Errorf("failed to create contact: %s", resp.Status)
	}

	var createdContact Contact
	if err := json.NewDecoder(resp.Body).Decode(&createdContact); err != nil {
		return Contact{}, err
	}
	return createdContact, nil
}

func (s *Secure) UpdateContact(id string, contact Contact) (Contact, error) {
	if s == nil {
		log.Println("Secure instance (s) is nil in UpdateContact")
		return Contact{}, fmt.Errorf("internal server error: secure context not initialized")
	}
	jsonData, err := json.Marshal(contact)
	if err != nil {
		return Contact{}, err
	}

	url := fmt.Sprintf("%s/%s", baseURL, id)
	resp, err := s.doRequest(http.MethodPut, url, jsonData)
	if err != nil {
		return Contact{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Contact{}, fmt.Errorf("failed to update contact: %s", resp.Status)
	}

	var updatedContact Contact
	if err := json.NewDecoder(resp.Body).Decode(&updatedContact); err != nil {
		return Contact{}, err
	}
	return updatedContact, nil
}

func (s *Secure) FetchContact(id string) (Contact, error) {
	if s == nil {
		log.Println("Secure instance (s) is nil in FetchContact")
		return Contact{}, fmt.Errorf("internal server error: secure context not initialized")
	}
	url := fmt.Sprintf("%s/%s", baseURL, id)
	resp, err := s.doRequest(http.MethodGet, url, nil)
	if err != nil {
		return Contact{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Contact{}, fmt.Errorf("failed to fetch contact: %s", resp.Status)
	}

	var contact Contact
	if err := json.NewDecoder(resp.Body).Decode(&contact); err != nil {
		return Contact{}, err
	}
	return contact, nil
}
