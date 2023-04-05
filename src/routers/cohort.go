package routers

import (
	sys "armadabackend/logging"
	db "armadabackend/services/databaseServices"
	"encoding/json"
	"net/http"
)

type cohortWBody struct {
	Message string            `json:"message"`
	Result  []db.CohortRecord `json:"result"`
}

func cohortGetAll(w http.ResponseWriter, r *http.Request) {
	cohorts, err := db.GetAllCohortRecords()
	if err != nil {
		sys.Log.Errorf("failed to get all cohorts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		data := cohortWBody{
			Message: "Internal server error",
			Result:  []db.CohortRecord{},
		}
		body, err := json.Marshal(data)
		if err != nil {
			sys.Log.Errorf("failed to marshal data: %s", err)
		}
		w.Write(body)
	}
	sys.Log.Info("successfully fetched all cohorts")
	data := cohortWBody{
		Message: "Success: fetched all cohorts",
		Result:  cohorts,
	}
	body, err := json.Marshal(data)
	if err != nil {
		sys.Log.Errorf("failed to marshal data: %s", err)
	}
	w.Write(body)
}
