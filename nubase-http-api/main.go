//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world nubase-http-api --out gen ../wit
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ucarion/urlpath"
	"go.wasmcloud.dev/component/net/wasihttp"
)

// Since we don't run this program like a CLI, the `main` function is empty.
// Instead, setup is done via the `init` function
func main() {}

var getStatusPath = urlpath.New("/api/v1/_status")
var documentsPath = urlpath.New("/api/v1/documents")
var documentsByIdPath = urlpath.New("/api/v1/documents/:id")
var documentsLatestPath = urlpath.New("/api/v1/documents/latest")

// Entrypoint for the WebAssembly component
func init() {

	wasihttp.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

		// GET /api/v1/_status
		if _, ok := getStatusPath.Match(r.URL.Path); ok && r.Method == http.MethodGet {
			handleStatus(w, r)
			return
		}

		// GET /api/v1/documents
		if _, ok := documentsPath.Match(r.URL.Path); ok && r.Method == http.MethodGet {
			handleGetAllDocuments(w, r)
			return
		}

		// PUT /api/v1/documents
		if _, ok := documentsPath.Match(r.URL.Path); ok && r.Method == http.MethodPut {
			handlePutDocument(w, r)
			return
		}

		// GET /api/v1/documents/:id
		if match, ok := documentsByIdPath.Match(r.URL.Path); ok && r.Method == http.MethodGet {
			handleGetSingleDocumentById(w, r, match.Params["id"])
			return
		}

		// DELETE /api/v1/documents/:id
		if match, ok := documentsByIdPath.Match(r.URL.Path); ok && r.Method == http.MethodDelete {
			handleDeleteSingleDocumentById(w, r, match.Params["id"])
			return
		}

		// GET /api/v1/documents/latest
		if _, ok := documentsLatestPath.Match(r.URL.Path); ok && r.Method == http.MethodGet {
			handleGetLatestDocument(w, r)
			return
		}

		sendErrorResponse(w, fmt.Sprintf("unrecognized path [%s]", r.URL.Path))

	})
}

// Handle the status endpoint
func handleStatus(w http.ResponseWriter, r *http.Request) {
	sendSuccessResponse(w, "ok")
}

// Handle retrieving all documents in a paginated manner
func handleGetAllDocuments(w http.ResponseWriter, r *http.Request) {
	sendSuccessResponse(w, "handleGetAllDocuments")
}

// Handle retrieving all documents in a paginated manner
func handlePutDocument(w http.ResponseWriter, r *http.Request) {
	sendSuccessResponse(w, "handlePutDocument")
}

// Handle retrieving a single document by ID
func handleGetSingleDocumentById(w http.ResponseWriter, r *http.Request, documentId string) {
	sendSuccessResponse(w, "handleGetSingleDocumentById")
}

// Handle inserting a single document
func handleInsertSingleDocument(w http.ResponseWriter, r *http.Request) {
	sendSuccessResponse(w, "handleInsertSingleDocument")
}

// Handle deleting a single document by ID
func handleDeleteSingleDocumentById(w http.ResponseWriter, r *http.Request, documentId string) {
	sendSuccessResponse(w, "handleDeleteSingleDocumentById")
}

// Handle getting the latest inserted document
func handleGetLatestDocument(w http.ResponseWriter, r *http.Request) {
	sendSuccessResponse(w, "handleGetLatestDocument")
}

// Response is a generic struct that holds any data of type T
type SuccessResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
}

func makeSuccessResponse[T any](data T) SuccessResponse[T] {
	return SuccessResponse[T]{
		Status:  "success",
		Data:    data,
	}
}

func sendSuccessResponse[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(makeSuccessResponse(data)); err != nil {
		http.Error(w, fmt.Sprintf("unexpected failure while building response: %s", err), http.StatusInternalServerError)
	}
}

// Response is a generic struct that holds any data of type T
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func makeErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  msg,
	}
}

func sendErrorResponse(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(makeErrorResponse(msg)); err != nil {
		http.Error(w, fmt.Sprintf("unexpected failure while building response: %s", err), http.StatusInternalServerError)
	}
}
