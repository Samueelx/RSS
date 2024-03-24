package main

import "net/http"

/*
*Below is the function signature that you have to use
if you want to define an http hanlder in a way that the Go standard library expects.
It always takes a response writer as the first parameter, and a
pointer to the http request as the second parameter.
*/
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
