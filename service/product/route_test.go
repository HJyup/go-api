package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-api/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(productStore, userStore)

	t.Run("should fail if the user create product payload is invalid", func(t *testing.T) {
		user := types.CreateProductPayload{
			Name:        "Product",
			Description: "",
			Price:       10.0,
			Quantity:    10,
		}
		marshalled, _ := json.Marshal(user)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("should correctly create a product", func(t *testing.T) {
		user := types.CreateProductPayload{
			Name:        "Product",
			Description: "A very nice product",
			Price:       10.0,
			Quantity:    10,
		}
		marshalled, _ := json.Marshal(user)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rec.Code)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) CreateProduct(_ types.Product) error {
	return nil
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(_ string) (*types.User, error) {
	return &types.User{}, fmt.Errorf("error with email")
}

func (m *mockUserStore) CreateUser(_ types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(_ int) (*types.User, error) {
	return &types.User{}, nil
}
