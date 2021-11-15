package splio

import (
	"testing"
)

const universe = ""
const apiKey = ""
const testEmail = ""

func TestClient_Authenticate(t *testing.T) {
	client := NewClient(universe, "")
	apiError := client.Authenticate()
	if apiError == nil {
		t.Error("Key is invalid. Error is expected.")
	}

	client = NewClient(universe, apiKey)
	apiError = client.Authenticate()
	if apiError != nil {
		t.Error(apiError)
	}
}

func TestClient_CreateContact(t *testing.T) {
	client := NewClient(universe, apiKey)

	badContact := &Contact{
		Email: "",
	}
	apiError := client.CreateContact(badContact)
	if apiError == nil {
		t.Fatalf("Invalid required fields. Error is expected.")
	}

	firstName := "test"

	contact := &Contact{
		Email: testEmail,
		FirstName: &firstName,
	}

	apiError = client.CreateContact(contact)
	if apiError != nil {
		t.Error(apiError)
	}
}

func TestClient_ListContact(t *testing.T) {
	client := NewClient(universe, apiKey)

	searchFields := []SearchField{
		{Email, Equal, testEmail},
	}
	list, apiError := client.ListContact(1, 1, searchFields)
	if apiError != nil {
		t.Fatalf("List contact failed. Err: %v", apiError)
	}

	if list.Count != 1 {
		t.Errorf("Count result failed. Expected 1, got %d", list.Count)
	}

	if list.Elements[0].Email != testEmail {
		t.Errorf("Contact email expected %s, got %s", testEmail, list.Elements[0].Email)
	}
}

func TestClient_GetContact(t *testing.T) {
	client := NewClient(universe, apiKey)

	contact, apiError := client.GetContact(testEmail)
	if apiError != nil {
		t.Fatalf("Get contact failed. Err: %v", apiError)
	}

	if contact.Email != testEmail {
		t.Errorf("Contact email expected %s, got %s", testEmail, contact.Email)
	}
}

func TestClient_EditContact(t *testing.T) {
	client := NewClient(universe, apiKey)

	contact, apiError := client.GetContact(testEmail)
	if apiError != nil {
		t.Fatalf("Get contact failed. Err: %v", apiError)
	}

	newFirstName := "test2"
	contact.FirstName = &newFirstName

	apiError = client.EditContact(contact)
	if apiError != nil {
		t.Fatalf("Update contact failed. Err: %v", apiError)
	}
}

func TestClient_DeleteContact(t *testing.T) {
	client := NewClient(universe, apiKey)

	apiError := client.DeleteContact(testEmail)
	if apiError != nil {
		t.Fatalf("Delete contact %s failed. Err: %v", testEmail, apiError)
	}

}
