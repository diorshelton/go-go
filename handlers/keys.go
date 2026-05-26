package handlers

import (
	"encoding/json"
	"go-go/cache"
	"go-go/server"
)

func HandleGetAll(store *cache.Store) func(*server.Request) *server.Response {
	return func(r *server.Request) *server.Response {

		storeMap := store.All()

		storeAsBytes, err := json.Marshal(storeMap)
		if err != nil {
			response := err.Error()
			return &server.Response{StatusCode: 500, Body: []byte(response)}
		}

		return &server.Response{StatusCode: 200, Body: storeAsBytes}
	}

}
