package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest_golang/constants"
	"github.com/rest_golang/data"
	"github.com/rest_golang/mongo"
	"github.com/rest_golang/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Post(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	body := data.Doc{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	err = json.Unmarshal(requestBody, &body)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	client, err := mongo.ConnectToDataBase(ctx, constants.MongoURL)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database(constants.DataBase).Collection(constants.Collection)

	_, err = collection.InsertOne(ctx, body)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	utils.SucessResponse(w)
}

func Get(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	client, err := mongo.ConnectToDataBase(ctx, constants.MongoURL)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database(constants.DataBase).Collection(constants.Collection)

	curr, err := collection.Find(ctx, bson.D{})
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	var docArray []data.Doc

	for curr.Next(ctx) {
		var each data.Doc
		err := curr.Decode(&each)
		if err != nil {
			utils.ErrorResponse(w, err, 500)
			return
		}
		docArray = append(docArray, each)
	}

	utils.SucessResponseWithData(w, docArray)
}

func Put(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	params := mux.Vars(r)
	id := params["id"]
	body := data.Doc{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	err = json.Unmarshal(requestBody, &body)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	client, err := mongo.ConnectToDataBase(ctx, constants.MongoURL)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database(constants.DataBase).Collection(constants.Collection)

	res, err := collection.UpdateOne(ctx, bson.M{"name": id},
		bson.D{
			{"$set", body},
		})
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	utils.SucessResponseWithData(w, res.ModifiedCount)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	params := mux.Vars(r)
	id := params["id"]

	client, err := mongo.ConnectToDataBase(ctx, constants.MongoURL)
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database(constants.DataBase).Collection(constants.Collection)

	res, err := collection.DeleteOne(ctx, bson.M{"name": id})
	if err != nil {
		utils.ErrorResponse(w, err, 500)
		return
	}

	utils.SucessResponseWithData(w, res.DeletedCount)

}
