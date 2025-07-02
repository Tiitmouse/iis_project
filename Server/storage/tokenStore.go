package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type TokenStore struct {
	mu       sync.Mutex
	Tokens   map[string]string `json:"tokens"`
	filePath string
}

var (
	tokenStoreInstance *TokenStore
	tokenStoreOnce     sync.Once
)

func GetTokenStore() *TokenStore {
	tokenStoreOnce.Do(func() {
		store := &TokenStore{
			Tokens:   make(map[string]string),
			filePath: "tokens.json", // default
		}
		if err := store.load(); err != nil {
			if !os.IsNotExist(err) {
				fmt.Println("Error loading token store:", err)
			}
		}
		tokenStoreInstance = store
	})
	return tokenStoreInstance
}

// kept for backward compatibility
func NewTokenStore(filePath string) (*TokenStore, error) {
	tokenStoreOnce.Do(func() {
		store := &TokenStore{
			Tokens:   make(map[string]string),
			filePath: filePath,
		}
		if err := store.load(); err != nil {
			if !os.IsNotExist(err) {
				fmt.Println("Error loading token store:", err)
			}
		}
		tokenStoreInstance = store
	})
	return tokenStoreInstance, nil
}

func SetTokenStoreFilePath(filePath string) {
	if tokenStoreInstance == nil {
		tokenStoreOnce.Do(func() {
			store := &TokenStore{
				Tokens:   make(map[string]string),
				filePath: filePath,
			}
			if err := store.load(); err != nil {
				if !os.IsNotExist(err) {
					fmt.Println("Error loading token store:", err)
				}
			}
			tokenStoreInstance = store
		})
	}
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
