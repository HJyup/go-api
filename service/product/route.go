package product

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go-api/types"
	"go-api/utils"
	"net/http"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", handler.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodPost)
}

func (handler *Handler) handleGetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := handler.store.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		var errValidation validator.ValidationErrors
		errors.As(err, &errValidation)
		utils.WriteError(w, http.StatusBadRequest, errValidation)
	}

	err := handler.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
