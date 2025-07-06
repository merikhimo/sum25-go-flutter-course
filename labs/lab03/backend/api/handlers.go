package api

import (
	"encoding/json"
	"fmt"
	"lab03-backend/models"
	"lab03-backend/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler holds the storage instance
type Handler struct {
	// TODO: Add storage field of type *storage.MemoryStorage
	storage *storage.MemoryStorage
}

// NewHandler creates a new handler instance
func NewHandler(storage *storage.MemoryStorage) *Handler {
	// TODO: Return a new Handler instance with provided storage
	handler := Handler{storage: storage}
	return &handler
}

// SetupRoutes configures all API routes
func (h *Handler) SetupRoutes() *mux.Router {
	// TODO: Create a new mux router
	// TODO: Add CORS middleware
	// TODO: Create API v1 subrouter with prefix "/api"
	// TODO: Add the following routes:
	// GET /messages -> h.GetMessages
	// POST /messages -> h.CreateMessage
	// PUT /messages/{id} -> h.UpdateMessage
	// DELETE /messages/{id} -> h.DeleteMessage
	// GET /status/{code} -> h.GetHTTPStatus
	// GET /health -> h.HealthCheck
	// TODO: Return the router
	return nil
}

// GetMessages handles GET /api/messages
func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GetMessages handler
	// Get all messages from storage
	// Create successful API response
	// Write JSON response with status 200
	// Handle any errors appropriately
	//messages := h.storage.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// CreateMessage handles POST /api/messages
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement CreateMessage handler
	// Parse JSON request body into CreateMessageRequest
	// Validate the request
	// Create message in storage
	// Create successful API response
	// Write JSON response with status 201
	// Handle validation and storage errors appropriately
}

// UpdateMessage handles PUT /api/messages/{id}
func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement UpdateMessage handler
	// Extract ID from URL path variables
	// Parse JSON request body into UpdateMessageRequest
	// Validate the request
	// Update message in storage
	// Create successful API response
	// Write JSON response with status 200
	// Handle validation, parsing, and storage errors appropriately
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.writeJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "missing message ID in URL",
		})
		return
	}
}

// DeleteMessage handles DELETE /api/messages/{id}
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.writeJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "missing message ID in URL",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		h.writeJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid message ID",
		})
		return
	}
	err = h.storage.Delete(id)
	if err != nil {
		h.writeJSON(w, http.StatusNotFound, models.APIResponse{
			Success: false,
			Error:   "message not found",
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetHTTPStatus handles GET /api/status/{code}
func (h *Handler) GetHTTPStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	codeStr, ok := vars["code"]
	if !ok {
		h.writeJSON(w, http.StatusBadRequest,
			models.APIResponse{Success: false, Error: "status code missing"})
		return
	}
	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 100 || code > 599 {
		h.writeJSON(w, http.StatusBadRequest,
			models.APIResponse{Success: false, Error: "invalid status code"})
		return
	}

	resp := models.HTTPStatusResponse{
		StatusCode:  code,
		ImageURL:    "https://http.cat/" + strconv.Itoa(code),
		Description: getHTTPStatusDescription(code),
	}

	// 4. Оборачиваем в унифицированный API‑ответ и отправляем
	h.writeJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    resp,
	})
}

// HealthCheck handles GET /api/health
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement HealthCheck handler
	// Create a simple health check response with:
	//   - status: "ok"
	//   - message: "API is running"
	//   - timestamp: current time
	//   - total_messages: count from storage
	// Write JSON response with status 200
	//res := models.APIResponse{}
}

// Helper function to write JSON responses
func (h *Handler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	// TODO: Implement writeJSON helper
	// Set Content-Type header to "application/json"
	// Set status code
	// Encode data as JSON and write to response
	// Log any encoding errors
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		// Логируем ошибку — можно использовать log, либо любой другой логгер
		fmt.Printf("writeJSON: failed to encode JSON response: %v", err)
	}
}

// Helper function to write error responses
func (h *Handler) writeError(w http.ResponseWriter, status int, message string) {
	// TODO: Implement writeError helper
	// Create APIResponse with Success: false and Error: message
	// Use writeJSON to send the error response
	res := models.APIResponse{Success: false, Data: "", Error: message}
	h.writeJSON(w, status, res)
}

// Helper function to parse JSON request body
func (h *Handler) parseJSON(r *http.Request, dst interface{}) error {
	// TODO: Implement parseJSON helper
	// Create JSON decoder from request body
	// Decode into destination interface
	// Return any decoding errors
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

// Helper function to get HTTP status description
func getHTTPStatusDescription(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"

	default:
		return "Unknown Status"
	}
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	// TODO: Implement CORS middleware
	// Set the following headers:
	// Access-Control-Allow-Origin: *
	// Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
	// Access-Control-Allow-Headers: Content-Type, Authorization
	// Handle OPTIONS preflight requests
	// Call next handler for non-OPTIONS requests

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
