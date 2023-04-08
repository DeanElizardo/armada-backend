package databaseServices

type CourseRecord struct {
	ID       int
	Name     string
	CohortID int
}

func GetCohortCourses(cohortID int) ([]CourseRecord, error) {
	var courseRecords []CourseRecord

	db, err := connectDB()
	if err != nil {
		return courseRecords, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM courses WHERE cohortId=$1`
	rows, err := db.Query(sqlStatement, cohortID)
	if err != nil {
		return courseRecords, err
	}
	for rows.Next() {
		var id int
		var name string
		var cID int
		err = rows.Scan(&id, &name, &cID)
		if err != nil {
			return courseRecords, err
		}
		record := CourseRecord{id, name, cID}
		courseRecords = append(courseRecords, record)
	}

	return courseRecords, nil
}
