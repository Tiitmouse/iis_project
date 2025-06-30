package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type TokenStore struct {
	mu       sync.Mutex
	Tokens   map[string]string `json:"tokens"`
	filePath string
}

func NewTokenStore(filePath string) (*TokenStore, error) {
	store := &TokenStore{
		Tokens:   make(map[string]string),
		filePath: filePath,
	}
	if err := store.load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}
	return store, nil
}

func (s *TokenStore) AddToken(userID, token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Tokens[userID] = token
	return s.save()
}

func (s *TokenStore) GetToken(userID string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Tokens[userID]
}

func (s *TokenStore) RemoveToken(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Tokens, userID)
	return s.save()
}

func (s *TokenStore) save() error {
	data, err := json.MarshalIndent(s.Tokens, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}

func (s *TokenStore) load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.Tokens)
}
