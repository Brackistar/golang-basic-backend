package api

import (
	"net/http"
)

func HandleServerless(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
