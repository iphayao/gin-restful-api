package main

import (
	"fmt"
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var id int

func TestGetAllCustomer(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestCreateCustomer(t *testing.T) {
	r := setupRouter()

	body := []byte(`{
		"firstName": "John",
		"lastName": "Doe",
		"age": 25,
		"email": "john.doe@mail.com"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)


	id = int(resp["id"].(float64))
	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, "John", resp["firstName"])
	assert.Equal(t, "Doe", resp["lastName"])
	assert.Equal(t, float64(25), resp["age"])
	assert.Equal(t, "john.doe@mail.com", resp["email"])
}

func TestGetCustomer(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/customers/%d", id), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, float64(id), resp["id"])
	assert.Equal(t, "John", resp["firstName"])
	assert.Equal(t, "Doe", resp["lastName"])
	assert.Equal(t, float64(25), resp["age"])
	assert.Equal(t, "john.doe@mail.com", resp["email"])
}

func TestUpdateCustomer(t *testing.T) {
	r := setupRouter()

	body := []byte(`{
		"firstName": "John",
		"lastName": "Doe",
		"age": 26,
		"email": "john.doe@mail.com"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/customers/%d", id), bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, float64(id), resp["id"])
	assert.Equal(t, "John", resp["firstName"])
	assert.Equal(t, "Doe", resp["lastName"])
	assert.Equal(t, float64(26), resp["age"])
	assert.Equal(t, "john.doe@mail.com", resp["email"])
}

func TestDeleteCustomer(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/customers/%d", id), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
