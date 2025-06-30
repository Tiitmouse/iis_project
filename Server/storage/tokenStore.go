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

var (
	tokenStoreInstance *TokenStore
	tokenStoreOnce     sync.Once
)

// GetTokenStore returns the singleton instance of TokenStore
func GetTokenStore() *TokenStore {
	tokenStoreOnce.Do(func() {
		store := &TokenStore{
			Tokens:   make(map[string]string),
			filePath: "tokens.json", // default file path
		}
		if err := store.load(); err != nil {
			if !os.IsNotExist(err) {
				// Log error but continue with empty store
				// In production, you might want to handle this differently
			}
		}
		tokenStoreInstance = store
	})
	return tokenStoreInstance
}

// NewTokenStore is deprecated, use GetTokenStore() instead
// Kept for backward compatibility
func NewTokenStore(filePath string) (*TokenStore, error) {
	tokenStoreOnce.Do(func() {
		store := &TokenStore{
			Tokens:   make(map[string]string),
			filePath: filePath,
		}
		if err := store.load(); err != nil {
			if !os.IsNotExist(err) {
				// Log error but continue with empty store
			}
		}
		tokenStoreInstance = store
	})
	return tokenStoreInstance, nil
}

// SetTokenStoreFilePath sets a custom file path for the token store
// This should be called before the first call to GetTokenStore() to take effect
func SetTokenStoreFilePath(filePath string) {
	if tokenStoreInstance == nil {
		// If instance doesn't exist yet, we can set the path
		tokenStoreOnce.Do(func() {
			store := &TokenStore{
				Tokens:   make(map[string]string),
				filePath: filePath,
			}
			if err := store.load(); err != nil {
				if !os.IsNotExist(err) {
					// Log error but continue with empty store
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
