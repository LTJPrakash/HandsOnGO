package controller

import (
	"context"
	"ecommers/model"
	"ecommers/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Creates a new product with the provided details
// @Tags Product
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data"
// @Success 201 {object} model.Product "Product created successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error creating product"
// @Router /api/products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		// we will see response manager
		utils.BadRequestResponse(err, w)
		return
	}

	product.Id = primitive.NewObjectID()
	_, err = productCollection.InsertOne(context.Background(), product)
	if err != nil {
		utils.ErrorResponse(err, w)
		return
	}

	utils.SuccessResponse("Product created successfully", product, w)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieves a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.Product "Product fetched successfully"
// @Failure 400 {string} string "Invalid ID format"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Error fetching product"
// @Router /api/product/{id} [get]

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	// Get the product ID from the URL parameters
	params := mux.Vars(r)
	productId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.BadRequestResponse(err, w)
		return
	}

	// Query the product collection for the specific ID
	var product model.Product
	err = productCollection.FindOne(context.Background(), bson.M{"_id": productId}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.NotFoundResponse("Product not found", w)
		} else {
			utils.ErrorResponse(err, w)
		}
		return
	}

	utils.SuccessResponse("Product fetched successfully", product, w)
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieves a list of all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} model.Product "Products fetched successfully"
// @Failure 500 {string} string "Error fetching products"
// @Router /api/products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	cur, err := productCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		utils.ErrorResponse(err, w)
		return
	}

	var products []primitive.M
	for cur.Next(context.Background()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			utils.ErrorResponse(err, w)
			return
		}
		products = append(products, product)
	}

	utils.SuccessResponse("Products fetched successfully", products, w)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Updates an existing product with the provided details
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body model.Product true "Product data"
// @Success 200 {object} model.Product "Product updated successfully"
// @Failure 400 {string} string "Invalid ID format or request payload"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Error updating product"
// @Router /api/product/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	productId, _ := primitive.ObjectIDFromHex(params["id"])

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.BadRequestResponse(err, w)
		return
	}

	update := bson.M{"$set": product}
	_, err = productCollection.UpdateOne(context.Background(), bson.M{"_id": productId}, update)
	if err != nil {
		utils.ErrorResponse(err, w)
		return
	}

	utils.SuccessResponse("Product updated successfully", product, w)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Deletes a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string "Product deleted successfully"
// @Failure 400 {string} string "Invalid ID format"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Error deleting product"
// @Router /api/product/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	productId, _ := primitive.ObjectIDFromHex(params["id"])

	_, err := productCollection.DeleteOne(context.Background(), bson.M{"_id": productId})
	if err != nil {
		utils.ErrorResponse(err, w)
		return
	}

	utils.SuccessResponse("Product deleted successfully", nil, w)
}
