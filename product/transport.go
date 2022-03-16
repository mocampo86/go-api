package product

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GetAllFindAllRequest struct {
}

//Product
func DecodeGetRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return GetAllFindAllRequest{}, err
}

func DecodePostProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("could not read the body")
	}

	var product Product
	err = json.Unmarshal(body, &product)
	return product, nil
}

func DecodeProductFilterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("could not read the body")
	}

	var filter ProductFilterRequest
	err = json.Unmarshal(body, &filter)
	return filter, nil
}

func DecodeProductIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	log.Println("Obteined param: ", idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	request := ProductIdRequest{ProductId: id}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
