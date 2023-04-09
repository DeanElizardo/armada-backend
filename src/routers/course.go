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

type courseRBody struct {
	Name     string `json:"name"`
	CohortID int    `json:"cohortId"`
	CourseID int    `json:"courseId"`
}

// Create
func coursesCreate(w http.ResponseWriter, r *http.Request) {
	var rBody courseRBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		errorOut400(&w, "coursesCreate: unable to parse request body", err)
		return
	}

	course, err := db.CreateCourseRecord(rBody.Name, rBody.CohortID)
	if err != nil {
		errorOut500(&w, "coursesCreate: failed to create new course record", err)
		return
	}
	msg := fmt.Sprintf("successfully created new course: id=%d name=%s", course[0].ID, course[0].Name)
	sys.Log.Info(msg)

	data := courseWBody{
		Message: msg,
		Result:  course,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		errorOut500(&w, "coursesCreate: unable to marshal data", err)
		return
	}

	w.Write(wBody)
}

// Read
// !NB: The .ts version of this function returns all cohorts as well
func coursesGetAll(w http.ResponseWriter, r *http.Request) {
	courseRecords, err := db.GetAllCourses()
	if err != nil {
		errorOut500(&w, "coursesGetAll: unable to retrieve courses", err)
		return
	}
	msg := "successfully fetched all courses"
	sys.Log.Info(msg)

	data := courseWBody{
		Message: msg,
		Result:  courseRecords,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		errorOut500(&w, "coursesGetAll: unable to marshal data", err)
		return
	}

	w.Write(wBody)
}

func coursesGetForCohort(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		errorOut400(&w, "coursesGetForCohort: failed to parse id from request params", err)
		return
	}

	courseRecords, err := db.GetCohortCourses(id)
	if err != nil {
		errorOut500(&w, "coursesGetForCohort: unable to get courses for cohort", err)
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
		errorOut500(&w, "coursesGetForCohort: unable to marshal data", err)
		return
	}

	w.Write(wBody)
}

func coursesGetUsers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		errorOut400(&w, "coursesGetUsers: unable to parse id from request param", err)
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	userRecords, err := db.GetCourseUsers(id)
	if err != nil {
		errorOut500(&w, "coursesGetUsers: failed to fetch user records", err)
		return
	}
	msg := fmt.Sprintf("successfully fetched all user records for course id %d", id)
	sys.Log.Info(msg)

	data := userWBody{
		Message: msg,
		Result:  userRecords,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		errorOut500(&w, "courseGetUsers: failed to marshal data", err)
		return
	}

	w.Write(wBody)
}

// Update
func coursesUpdate(w http.ResponseWriter, r *http.Request) {
	var rBody courseRBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		errorOut500(&w, "coursesUpdate: failed to parse request body", err)
		return
	}

	course, err := db.UpdateCourse(rBody.CourseID, rBody.Name)
	if err != nil {
		errorOut500(&w, "coursesUpdate: uanable to update course", err)
		return
	}
	msg := fmt.Sprintf("successfully updated course id %d with name %s", course[0].ID, course[0].Name)
	sys.Log.Info(msg)

	data := courseWBody{
		Message: msg,
		Result:  course,
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		errorOut500(&w, "courseUpdate: unable to marshal data", err)
	}

	w.Write(wBody)
}

// Delete
func coursesDelete(w http.ResponseWriter, r *http.Request) {
	courseId, err := strconv.Atoi(getField(r, 0))
	if err != nil {
		errorOut400(&w, "coursesDelete: unable to parse request params", err)
		return
	}

	deleteCount, err := db.DeleteCourse(courseId)
	if err != nil {
		errorOut500(&w, "courseDelete: unable to delete course record", err)
		return
	}
	if deleteCount != 0 {
		sys.Log.Infof("coursesDelete: successfully deleted course (id=%s)", courseId)
	} else {
		sys.Log.Infof("coursesDelete: completed deletion action; no records with id=%s removed; possible record D.N.E.", courseId)
	}

	data := cohortWBody{
		Message: fmt.Sprintf("success: deleted %d courses", deleteCount),
		Result:  []db.CohortRecord{},
	}
	wBody, err := json.Marshal(data)
	if err != nil {
		errorOut500(&w, "coursesDelete: unable to marshal data", err)
		return
	}

	w.Write(wBody)
}
