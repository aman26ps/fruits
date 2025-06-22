package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetAllFruits(t *testing.T) {
	fruits = []Fruit{
		{ID: 1, Name: "apple", Color: "red"},
	}

	req := httptest.NewRequest("GET", "/fruits", nil)
	w := httptest.NewRecorder()

	getAllFruits(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result []Fruit
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil || len(result) != 1 {
		t.Errorf("Expected 1 fruit, got: %v", result)
	}
}

func TestAddFruit(t *testing.T) {
	fruits = []Fruit{}

	body := `{"fruit":"kiwi","color":"green"}`
	req := httptest.NewRequest("POST", "/fruits", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	addFruit(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result Fruit
	json.NewDecoder(w.Body).Decode(&result)

	if result.Name != "kiwi" {
		t.Errorf("Expected fruit name kiwi, got %s", result.Name)
	}
}
