package MongoLite

import (
	"os"
	"testing"
)

func setup() {
	// Remove the data file before each test to ensure a clean state.
	os.Remove(dataFile)
}

func TestInsertDocument(t *testing.T) {
	setup()
	err := InsertDocument(`{"name": "test", "value": 123}`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the document was inserted.
	results, err := SearchDocuments(`{"name": "test"}`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := `[{"name":"test","value":123}]`
	if results != expected {
		t.Fatalf("Expected %v, got %v", expected, results)
	}
}

func TestSearchDocuments(t *testing.T) {
	setup()
	InsertDocument(`{"name": "test1", "value": 123}`)
	InsertDocument(`{"name": "test2", "value": 456}`)

	results, err := SearchDocuments(`{"value": 123}`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := `[{"name":"test1","value":123}]`
	if results != expected {
		t.Fatalf("Expected %v, got %v", expected, results)
	}
}

func TestDeleteDocuments(t *testing.T) {
	setup()
	InsertDocument(`{"name": "test1", "value": 123}`)
	InsertDocument(`{"name": "test2", "value": 456}`)

	deletedCount, err := DeleteDocuments(`{"value": 123}`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if deletedCount != 1 {
		t.Fatalf("Expected 1 document to be deleted, got %d", deletedCount)
	}

	// Verify the document was deleted.
	results, err := SearchDocuments(`{"value": 123}`)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := `[]`
	if results != expected {
		t.Fatalf("Expected %v, got %v", expected, results)
	}
}
