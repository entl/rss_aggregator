package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}
