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

func HandleGet(store *cache.Store) func(*server.Request) *server.Response {
	return func(r *server.Request) *server.Response {

		if r.Params["id"] == "" {
			return &server.Response{StatusCode: 400, Body: []byte("Bad Request")}
		}

		keyVal, err := store.Get(r.Params["id"])
		if err != nil {
			response := err.Error()
			return &server.Response{StatusCode: 404, Body: []byte(response)}
		}

		return &server.Response{StatusCode: 200, Body: keyVal}

	}

}

func HandlePut(store *cache.Store) func(*server.Request) *server.Response {
	return func(r *server.Request) *server.Response {

		if r.Params["id"] == "" || len(r.Body) == 0 {
			return &server.Response{StatusCode: 400, Body: []byte("invalid/empty body")}
		}

		ok := json.Valid(r.Body)

		if !ok {
			return &server.Response{StatusCode: 400, Body: []byte("invalid")}
		}

		val := r.Body

		store.Set(r.Params["id"], val)

		return &server.Response{StatusCode: 201, Body: val}
	}
}

func HandleDelete(store *cache.Store) func(*server.Request) *server.Response {
	return func(r *server.Request) *server.Response {

		if r.Params["id"] == "" {
			return &server.Response{StatusCode: 400, Body: []byte("Bad Request")}
		}

		val, err := store.Delete(r.Params["id"])
		if err != nil {
			response := err.Error()
			return &server.Response{StatusCode: 404, Body: []byte(response)}
		}

		return &server.Response{StatusCode: 200, Body: val}
	}
}
