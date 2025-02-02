package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

// InsertDocument is the C-accessible function for inserting a document.
// It receives a C string (JSON document) and returns a C string with a result message.
//
//export InsertDocument
func InsertDocument(cjson *C.char) *C.char {
	jsonStr := C.GoString(cjson)
	err := insertDocument(jsonStr)
	var result string
	if err != nil {
		result = "InsertDocument error: " + err.Error()
	} else {
		result = "InsertDocument: document successfully inserted"
	}
	return C.CString(result)
}

// SearchDocument is the C-accessible function for searching documents.
// It receives a C string (JSON query) and returns a C string containing a JSON array of
// matching documents, or an error message.
//
//export SearchDocument
func SearchDocument(cquery *C.char) *C.char {
	queryStr := C.GoString(cquery)
	results, err := searchDocuments(queryStr)
	var result string
	if err != nil {
		result = "SearchDocument error: " + err.Error()
	} else {
		result = results
	}
	return C.CString(result)
}

// DeleteDocument is the C-accessible function for deleting documents.
// It receives a C string (JSON query) and returns a C string with a result message.
//
//export DeleteDocument
func DeleteDocument(cquery *C.char) *C.char {
	queryStr := C.GoString(cquery)
	deletedCount, err := deleteDocuments(queryStr)
	var result string
	if err != nil {
		result = "DeleteDocument error: " + err.Error()
	} else {
		result = fmt.Sprintf("DeleteDocument: %d documents deleted", deletedCount)
	}
	return C.CString(result)
}
