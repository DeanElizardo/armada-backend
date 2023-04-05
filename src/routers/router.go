package routers

import "net/http"

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/cohort/all", cohortGetAll)

	return router
}
