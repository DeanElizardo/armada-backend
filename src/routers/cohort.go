package routers

import (
	sys "armadabackend/logging"
	db "armadabackend/services/databaseServices"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type cohortWBody struct {
	Message string            `json:"message"`
	Result  []db.CohortRecord `json:"result"`
}

type cohortRBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create
func cohortCreate(w http.ResponseWriter, r *http.Request) {
	// New cohort name is sent in the request body
	var rBody cohortRBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		sys.Log.Errorf("cohortCreate: failed to read request body: %s", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	cohort, err := db.CreateCohortRecord(rBody.Name)
	if err != nil {
		sys.Log.Errorf("cohortCreate: failed to create new cohort: %s", err)
		http.Error(w, "500 internal server error", http.StatusBadRequest)
		return
	}
	sys.Log.Infof("successfully created cohort (id=%d, name=%s)", cohort[0].ID, cohort[0].Name)

	data := cohortWBody{
		Message: "Success: created new cohort",
		Result:  cohort,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("cohortCreate: failed to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}

	w.Write(wBody)
}

// Read
func cohortGetAll(w http.ResponseWriter, r *http.Request) {
	cohorts, err := db.GetAllCohortRecords()
	if err != nil {
		sys.Log.Errorf("cohortGetAll: failed to get all cohorts: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	sys.Log.Info("successfully fetched all cohorts")

	data := cohortWBody{
		Message: "Success: fetched all cohorts",
		Result:  cohorts,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("cohortGetAll: failed to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(wBody)
}

func cohortGetSingle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		sys.Log.Errorf("cohortGetSingle: unable to parse cohort id param: %s", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
	}

	cohort, err := db.GetSingleCohort(id)
	if err != nil {
		sys.Log.Errorf("cohortGetSingle: failed to get cohort: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	sys.Log.Infof("successfully fetched single cohort (id=%d)", id)

	data := cohortWBody{
		Message: "success: fetched single cohort",
		Result:  cohort,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("cohortGetSingle: unable to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(wBody)
}

// Update
func cohortUpdate(w http.ResponseWriter, r *http.Request) {
	var rBody cohortRBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		sys.Log.Errorf("cohortUpdate: failed to read request body: %s", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	cohort, err := db.UpdateCohort(rBody.ID, rBody.Name)
	if err != nil {
		sys.Log.Errorf("cohortUpdate: failed to update cohort '%s %s': %s", rBody.ID, rBody.Name, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	sys.Log.Infof("successfully updated cohort (id=%s)", rBody.ID)

	data := cohortWBody{
		Message: "success: updated cohort",
		Result:  cohort,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("cohortUpdate: failed to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(wBody)
}

// Delete
func cohortDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		sys.Log.Errorf("cohortDelete: failed to parse id from request params: %s", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	deleteCount, err := db.DeleteCohort(id)
	if err != nil {
		sys.Log.Errorf("cohortDelete: failed to delete cohort '%s': %s", id, err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	if deleteCount != 0 {
		sys.Log.Infof("successfully deleted cohort (id=%s)", id)
	} else {
		sys.Log.Info("completed deletion action; no records with id=%s removed; possible record D.N.E.", id)
	}

	data := cohortWBody{
		Message: fmt.Sprintf("success: deleted %d cohorts", deleteCount),
		Result:  []db.CohortRecord{},
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("cohortDelete: unable to marshal data: %s", err)
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(wBody)
}
