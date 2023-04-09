package routers

import (
	sys "armadabackend/logging"
	"net/http"
)

func errorOut400(w *http.ResponseWriter, msg string, err error) {
	sys.Log.Errorf("%s: %s", msg, err)
	http.Error(*w, "400 bad request", http.StatusBadRequest)
}

func errorOut500(w *http.ResponseWriter, msg string, err error) {
	sys.Log.Errorf("%s: %s", msg, err)
	http.Error(*w, "500 internal server error", http.StatusInternalServerError)
}
