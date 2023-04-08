package routers

import (
	sys "armadabackend/logging"
	db "armadabackend/services/databaseServices"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type courseWBody struct {
	Message string            `json:"message"`
	Result  []db.CourseRecord `json:"result"`
}

// Create

// Read
func coursesGetForCohort(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		sys.Log.Errorf("coursesGetForCohort: failed to parse id from request params: %s", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	courseRecords, err := db.GetCohortCourses(id)
	if err != nil {
		sys.Log.Errorf("coursesGetForCohort: unable to get courses for cohort id '%d': %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	msg := fmt.Sprintf("successfully fetched courses associated with cohort id '%d'", id)
	sys.Log.Info(msg)

	data := courseWBody{
		Message: msg,
		Result:  courseRecords,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("coursesGetForCohort: unable to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(wBody)
}

// Update

// Delete
