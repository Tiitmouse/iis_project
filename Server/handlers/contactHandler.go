package handlers

import (
	"errors"
	"iis_server/storage"
	"iis_server/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var contactStore *storage.ContactStore

func SetContactStore(store *storage.ContactStore) {
	contactStore = store
}

func CreateContact(c *gin.Context) {
	var newContactInput structs.Contact

	if err := c.ShouldBindJSON(&newContactInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	contact := structs.NewContact()
	contact.Type = newContactInput.Type
	contact.Value = newContactInput.Value
	contact.Name = newContactInput.Name
	if newContactInput.Sources != nil {
		contact.Sources = newContactInput.Sources
	} else {
		contact.Sources = []string{}
	}

	createdContact, err := contactStore.AddContact(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdContact)
}

func GetAllContacts(c *gin.Context) {
	allContacts := contactStore.GetAllContacts()
	c.JSON(http.StatusOK, allContacts)
}

func GetContactByID(c *gin.Context) {
	id := c.Param("id")

	contact, exists := contactStore.GetContactByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact with ID " + id + " not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var updatedContactInput structs.Contact

	if err := c.ShouldBindJSON(&updatedContactInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	updatedContact, err := contactStore.UpdateContact(id, updatedContactInput)
	if err != nil {
		if errors.Is(err, errors.New("contact with ID not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedContact)
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")

	err := contactStore.DeleteContact(id)
	if err != nil {
		if errors.Is(err, errors.New("contact with ID not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact with ID " + id + " deleted successfully"})
}
