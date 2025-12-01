package api

import (
	"encoding/json"
	"errors"
	"go-circleci/types"
	"net/http"
	"strconv"
	"strings"
)

// extractIDFromPath extracts the product ID from a URL path like "/products/123"
// Returns the ID and an error if the ID is invalid or missing
func extractIDFromPath(path string) (int, error) {
	// Split path by "/"
	parts := strings.Split(strings.Trim(path, "/"), "/")
	
	// Expected format: ["products", "{id}"]
	if len(parts) < 2 {
		return 0, errors.New("product ID is required")
	}
	
	// Get the last segment as ID
	idStr := parts[len(parts)-1]
	
	// Parse as integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid product ID format: must be an integer")
	}
	
	if id <= 0 {
		return 0, errors.New("invalid product ID: must be greater than 0")
	}
	
	return id, nil
}

// validateProductName validates that a product name is not empty
func validateProductName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("product name is required")
	}
	return nil
}

// validatePrice validates that a price is non-negative
func validatePrice(price float64) error {
	if price < 0 {
		return errors.New("product price must be greater than or equal to 0")
	}
	return nil
}

// validateStock validates that stock is non-negative
func validateStock(stock int) error {
	if stock < 0 {
		return errors.New("product stock must be greater than or equal to 0")
	}
	return nil
}

// handleGetAllProducts handles GET /products requests
// Returns all products in the database as a JSON array
func (s *ApiServer) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	
	products, err := s.svc.GetAllProducts(r.Context())
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to retrieve products"})
		return
	}
	
	writeJson(w, http.StatusOK, products)
}

// handleGetProduct handles GET /products/{id} requests
// Returns a single product by ID
func (s *ApiServer) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	
	// Extract ID from path
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	
	// Get product from service
	product, err := s.svc.GetProductByID(r.Context(), id)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") {
			writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		// Check if it's a validation error
		if strings.Contains(err.Error(), "invalid") {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		// Other errors are internal server errors
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to retrieve product"})
		return
	}
	
	writeJson(w, http.StatusOK, product)
}

// handleCreateProduct handles POST /products requests
// Creates a new product and returns it with HTTP 201 status
func (s *ApiServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	
	// Parse JSON request body
	var req types.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON format"})
		return
	}
	
	// Create product via service
	product, err := s.svc.CreateProduct(r.Context(), &req)
	if err != nil {
		// Check if it's a validation error
		if strings.Contains(err.Error(), "required") || 
		   strings.Contains(err.Error(), "must be") ||
		   strings.Contains(err.Error(), "invalid") {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		// Other errors are internal server errors
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to create product"})
		return
	}
	
	// Return 201 Created with the new product
	writeJson(w, http.StatusCreated, product)
}

// handleUpdateProduct handles PUT /products/{id} requests
// Updates an existing product and returns the updated product
func (s *ApiServer) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Only allow PUT method
	if r.Method != http.MethodPut {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	
	// Extract ID from path
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	
	// Parse JSON request body
	var req types.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON format"})
		return
	}
	
	// Update product via service
	product, err := s.svc.UpdateProduct(r.Context(), id, &req)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") {
			writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		// Check if it's a validation error
		if strings.Contains(err.Error(), "required") || 
		   strings.Contains(err.Error(), "must be") ||
		   strings.Contains(err.Error(), "invalid") {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		// Other errors are internal server errors
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to update product"})
		return
	}
	
	writeJson(w, http.StatusOK, product)
}

// handleDeleteProduct handles DELETE /products/{id} requests
// Deletes a product and returns a success message
func (s *ApiServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	
	// Extract ID from path
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	
	// Delete product via service
	err = s.svc.DeleteProduct(r.Context(), id)
	if err != nil {
		// Check if it's a not found error
		if strings.Contains(err.Error(), "not found") {
			writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		// Check if it's a validation error
		if strings.Contains(err.Error(), "invalid") {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		// Other errors are internal server errors
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to delete product"})
		return
	}
	
	writeJson(w, http.StatusOK, map[string]string{"message": "product deleted successfully"})
}
