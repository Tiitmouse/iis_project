package storage

import (
	"errors"
	"iis_server/structs"
	"log"
	"sync"
)

type ContactStore struct {
	mu       sync.RWMutex
	contacts map[string]structs.Contact
}

func NewContactStore() *ContactStore {
	return &ContactStore{
		contacts: make(map[string]structs.Contact),
	}
}

func (cs *ContactStore) AddContact(contact structs.Contact) (structs.Contact, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if contact.ID == "" {
		return structs.Contact{}, errors.New("contact ID cannot be empty")
	}
	if _, exists := cs.contacts[contact.ID]; exists {
		return structs.Contact{}, errors.New("contact with ID already exists")
	}

	cs.contacts[contact.ID] = contact
	log.Printf("Storage: Added contact ID %s\n", contact.ID)
	return contact, nil
}

func (cs *ContactStore) GetContactByID(id string) (structs.Contact, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	contact, exists := cs.contacts[id]
	return contact, exists
}

func (cs *ContactStore) GetAllContacts() []structs.Contact {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	allContacts := make([]structs.Contact, 0, len(cs.contacts))
	for _, contact := range cs.contacts {
		allContacts = append(allContacts, contact)
	}
	return allContacts
}

func (cs *ContactStore) UpdateContact(id string, updatedContact structs.Contact) (structs.Contact, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, exists := cs.contacts[id]
	if !exists {
		return structs.Contact{}, errors.New("contact with ID not found")
	}

	updatedContact.ID = id
	cs.contacts[id] = updatedContact
	log.Printf("Storage: Updated contact ID %s\n", id)
	return updatedContact, nil
}

func (cs *ContactStore) DeleteContact(id string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, exists := cs.contacts[id]
	if !exists {
		return errors.New("contact with ID not found")
	}

	delete(cs.contacts, id)
	log.Printf("Storage: Deleted contact ID %s\n", id)
	return nil
}
