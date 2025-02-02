package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// Document represents a generic JSON document.
type Document map[string]interface{}

// dataFile is the path to our local storage file.
const dataFile = "mongo.local"

// lock serializes file access.
var lock sync.Mutex

// insertDocument takes a JSON string, parses it, and appends it to the file.
func InsertDocument(jsonStr string) error {
	var doc Document
	if err := json.Unmarshal([]byte(jsonStr), &doc); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Convert back to JSON for uniform formatting.
	output, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	lock.Lock()
	defer lock.Unlock()

	// Open (or create) the file in append mode.
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Append the JSON document with a newline.
	if _, err := f.Write(append(output, '\n')); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// searchDocuments reads the file and returns a JSON array (as string) of documents
// that match all key-value pairs in the query.
func SearchDocuments(queryStr string) (string, error) {
	var query map[string]interface{}
	if err := json.Unmarshal([]byte(queryStr), &query); err != nil {
		return "", fmt.Errorf("failed to parse query JSON: %w", err)
	}

	lock.Lock()
	defer lock.Unlock()

	// Open the file for reading.
	f, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// No data yet.
			return "[]", nil
		}
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var results []Document
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		var doc Document
		if err := json.Unmarshal([]byte(line), &doc); err != nil {
			// Skip invalid JSON.
			continue
		}
		if matchQuery(doc, query) {
			results = append(results, doc)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	output, err := json.Marshal(results)
	if err != nil {
		return "", fmt.Errorf("failed to encode results: %w", err)
	}
	if len(results) == 0 {
		return "[]", nil
	}
	return string(output), nil
}

// deleteDocuments removes documents that match the query.
// It returns the number of documents deleted.
func DeleteDocuments(queryStr string) (int, error) {
	var query map[string]interface{}
	if err := json.Unmarshal([]byte(queryStr), &query); err != nil {
		return 0, fmt.Errorf("failed to parse query JSON: %w", err)
	}

	lock.Lock()
	defer lock.Unlock()

	// Read the existing file content.
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil // Nothing to delete.
		}
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var updatedLines []string
	deletedCount := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var doc Document
		if err := json.Unmarshal([]byte(line), &doc); err != nil {
			// If this line is invalid JSON, skip it.
			continue
		}
		if matchQuery(doc, query) {
			deletedCount++
			// Do not include in updatedLines.
		} else {
			updatedLines = append(updatedLines, line)
		}
	}

	newContent := strings.Join(updatedLines, "\n")
	if newContent != "" {
		newContent += "\n" // Ensure file ends with a newline.
	}
	if err := ioutil.WriteFile(dataFile, []byte(newContent), 0644); err != nil {
		return deletedCount, fmt.Errorf("failed to write updated data: %w", err)
	}

	return deletedCount, nil
}

// matchQuery returns true if every key-value pair in the query exists in doc
// (the value comparison is done via formatted strings for simplicity).
func matchQuery(doc Document, query map[string]interface{}) bool {
	for key, qVal := range query {
		dVal, ok := doc[key]
		if !ok {
			return false
		}
		// Using formatted strings for a simple deep equality check.
		if fmt.Sprintf("%v", dVal) != fmt.Sprintf("%v", qVal) {
			return false
		}
	}
	return true
}
