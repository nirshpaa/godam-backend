package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Products : struct for set Products Dependency Injection
type Products struct {
	App   http.Handler
	Token string
}

// Run : http handler for run products testing
func (u *Products) Run(t *testing.T) {
	created := u.Create(t)
	id := created["data"].(map[string]interface{})["id"].(string)
	u.List(t)
	u.View(t, id)
	u.Update(t, id)
	u.Delete(t, id)
}

// List : http handler for returning list of products
func (u *Products) List(t *testing.T) {
	req := httptest.NewRequest("GET", "/products", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", u.Token)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("getting: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var list map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"status_code":    string("REBEL-200"),
		"status_message": string("OK"),
		"data": []interface{}{
			map[string]interface{}{
				"id":                  string("test-product-1"),
				"code":                string("PROD-1"),
				"name":                "Tes",
				"purchase_price":      float64(1),
				"sale_price":          float64(1),
				"minimum_stock":       float64(25),
				"image_url":           "",
				"company_id":          string("test-company-1"),
				"brand_id":            string("test-brand-1"),
				"product_category_id": string("test-category-1"),
			},
		},
	}

	if diff := cmp.Diff(want, list); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}

// Create : http handler for create new product
func (u *Products) Create(t *testing.T) map[string]interface{} {
	var created map[string]interface{}
	jsonBody := `
		{
			"code": "PROD-200",
			"name": "Tes",
			"purchase_price": 1,
			"sale_price": 1,
			"minimum_stock": 25,
			"brand_id": "test-brand-1",
			"product_category_id": "test-category-1"
		}
	`
	body := strings.NewReader(jsonBody)

	req := httptest.NewRequest("POST", "/products", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", u.Token)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusCreated != resp.Code {
		t.Fatalf("posting: expected status code %v, got %v", http.StatusCreated, resp.Code)
	}

	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	c := created["data"].(map[string]interface{})

	if c["id"] == "" || c["id"] == nil {
		t.Fatal("expected non-empty product id")
	}

	want := map[string]interface{}{
		"status_code":    "REBEL-200",
		"status_message": "OK",
		"data": map[string]interface{}{
			"id":                  c["id"],
			"code":                "PROD-200",
			"name":                "Tes",
			"purchase_price":      float64(1),
			"sale_price":          float64(1),
			"minimum_stock":       float64(25),
			"company_id":          "test-company-1",
			"brand_id":            "test-brand-1",
			"product_category_id": "test-category-1",
		},
	}

	if diff := cmp.Diff(want, created); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}

	return created
}

// View : http handler for retrieve product by id
func (u *Products) View(t *testing.T, id string) {
	req := httptest.NewRequest("GET", "/products/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", u.Token)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Fatalf("retrieving: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var fetched map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&fetched); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"status_code":    "REBEL-200",
		"status_message": "OK",
		"data": map[string]interface{}{
			"id":                  id,
			"code":                "PROD-200",
			"name":                "Tes",
			"purchase_price":      float64(1),
			"sale_price":          float64(1),
			"minimum_stock":       float64(25),
			"company_id":          "test-company-1",
			"brand_id":            "test-brand-1",
			"product_category_id": "test-category-1",
		},
	}

	// Fetched product should match the one we created.
	if diff := cmp.Diff(want, fetched); diff != "" {
		t.Fatalf("Retrieved product should match created. Diff:\n%s", diff)
	}
}

// Update : http handler for update product by id
func (u *Products) Update(t *testing.T, id string) {
	jsonBody := `
		{
			"code": "PROD-201",
			"name": "Tes Updated",
			"purchase_price": 2,
			"sale_price": 2,
			"minimum_stock": 30,
			"brand_id": "test-brand-1",
			"product_category_id": "test-category-1"
		}
	`
	body := strings.NewReader(jsonBody)

	req := httptest.NewRequest("PUT", "/products/"+id, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", u.Token)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Fatalf("updating: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var updated map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"status_code":    "REBEL-200",
		"status_message": "OK",
		"data": map[string]interface{}{
			"id":                  id,
			"code":                "PROD-201",
			"name":                "Tes Updated",
			"purchase_price":      float64(2),
			"sale_price":          float64(2),
			"minimum_stock":       float64(30),
			"company_id":          "test-company-1",
			"brand_id":            "test-brand-1",
			"product_category_id": "test-category-1",
		},
	}

	if diff := cmp.Diff(want, updated); diff != "" {
		t.Fatalf("Updated product should match expected. Diff:\n%s", diff)
	}
}

// Delete : http handler for delete product by id
func (u *Products) Delete(t *testing.T, id string) {
	req := httptest.NewRequest("DELETE", "/products/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", u.Token)
	resp := httptest.NewRecorder()

	u.App.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Fatalf("deleting: expected status code %v, got %v", http.StatusOK, resp.Code)
	}

	var deleted map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&deleted); err != nil {
		t.Fatalf("decoding: %s", err)
	}

	want := map[string]interface{}{
		"status_code":    "REBEL-200",
		"status_message": "OK",
	}

	if diff := cmp.Diff(want, deleted); diff != "" {
		t.Fatalf("Response did not match expected. Diff:\n%s", diff)
	}
}
