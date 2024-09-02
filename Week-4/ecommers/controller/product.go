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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
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

// GetProductByID retrieves a single product by its ID
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
