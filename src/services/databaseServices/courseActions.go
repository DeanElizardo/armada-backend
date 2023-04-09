package databaseServices

type CourseRecord struct {
	ID       int
	Name     string
	CohortID int
}

// Create
func CreateCourseRecord(name string, cohortID int) ([]CourseRecord, error) {
	var courseRecord []CourseRecord
	var id int
	db, err := connectDB()
	if err != nil {
		return courseRecord, err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO courses (name, cohortId) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(sqlStatement, name, cohortID).Scan(&id)
	if err != nil {
		return courseRecord, err
	}

	sqlStatement = `SELECT * FROM courses WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		return courseRecord, err
	}
	for row.Next() {
		var i int
		var n string
		var c int
		err = row.Scan(&i, &n, &c)
		if err != nil {
			return courseRecord, err
		}
		record := CourseRecord{i, n, c}
		courseRecord = append(courseRecord, record)
	}
	return courseRecord, nil
}

// Read
func GetAllCourses() ([]CourseRecord, error) {
	var courseRecords []CourseRecord
	db, err := connectDB()
	if err != nil {
		return courseRecords, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM courses`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return courseRecords, err
	}
	for rows.Next() {
		var i int
		var n string
		var c int
		err = rows.Scan(&i, &n, &c)
		if err != nil {
			return courseRecords, err
		}
		record := CourseRecord{i, n, c}
		courseRecords = append(courseRecords, record)
	}
	return courseRecords, nil
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

func GetSingleCourse(id int) ([]CourseRecord, error) {
	var courseRecord []CourseRecord

	db, err := connectDB()
	if err != nil {
		return courseRecord, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM courses WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		return courseRecord, err
	}
	for row.Next() {
		var i int
		var n string
		var c int
		err = row.Scan(&i, &n, &c)
		if err != nil {
			return courseRecord, nil
		}
		record := CourseRecord{i, n, c}
		courseRecord = append(courseRecord, record)
	}
	return courseRecord, nil
}

// Update
func UpdateCourse(id int, name string) ([]CourseRecord, error) {
	var courseRecord []CourseRecord

	db, err := connectDB()
	if err != nil {
		return courseRecord, err
	}
	defer db.Close()

	sqlStatement := `UPDATE courses SET name=$2 WHERE id=$1 RETURNING id, name, cohortId`
	var updatedID int
	var updatedName string
	var cohortID int
	err = db.QueryRow(sqlStatement, id, name).Scan(&updatedID, &updatedName, &cohortID)
	if err != nil {
		return courseRecord, err
	}

	record := CourseRecord{
		ID:       updatedID,
		Name:     updatedName,
		CohortID: cohortID,
	}
	courseRecord = append(courseRecord, record)
	return courseRecord, nil
}

// Delete
func DeleteCourse(id int) (int64, error) {
	var deleteCount int64
	db, err := connectDB()
	if err != nil {
		return deleteCount, err
	}
	defer db.Close()

	sqlStatement := `DELETE FROM courses WHERE id=$1`
	response, err := db.Exec(sqlStatement, id)
	if err != nil {
		return deleteCount, err
	}
	deleteCount, err = response.RowsAffected()
	if err != nil {
		return deleteCount, err
	}
	return deleteCount, nil
}
