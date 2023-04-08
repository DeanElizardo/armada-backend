package databaseServices

type CohortRecord struct {
	ID   int
	Name string
}

// Create
func CreateCohortRecord(cohortName string) ([]CohortRecord, error) {
	var cohortRecord []CohortRecord
	var id int
	db, err := connectDB()
	if err != nil {
		return cohortRecord, err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO cohorts (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(sqlStatement, cohortName).Scan(&id)
	if err != nil {
		return cohortRecord, err
	}

	sqlStatement = `SELECT * FROM cohorts WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		return cohortRecord, err
	}
	for row.Next() {
		var id int
		var name string
		err = row.Scan(&id, &name)
		if err != nil {
			return cohortRecord, err
		}
		record := CohortRecord{id, name}
		cohortRecord = append(cohortRecord, record)
	}
	return cohortRecord, nil
}

// Read
func GetAllCohortRecords() ([]CohortRecord, error) {
	var cohortRecords []CohortRecord

	db, err := connectDB()
	if err != nil {
		return cohortRecords, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM cohorts`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return cohortRecords, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return cohortRecords, err
		}
		record := CohortRecord{id, name}
		cohortRecords = append(cohortRecords, record)
	}

	return cohortRecords, nil
}

func GetSingleCohort(id int) ([]CohortRecord, error) {
	var cohortRecord []CohortRecord
	db, err := connectDB()
	if err != nil {
		return cohortRecord, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM cohorts WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		return cohortRecord, err
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		err = row.Scan(&id, &name)
		if err != nil {
			return cohortRecord, err
		}
		record := CohortRecord{id, name}
		cohortRecord = append(cohortRecord, record)
	}

	return cohortRecord, nil
}

// Update
func UpdateCohort(id int, name string) ([]CohortRecord, error) {
	var cohortRecord []CohortRecord
	db, err := connectDB()
	if err != nil {
		return cohortRecord, err
	}
	defer db.Close()

	sqlStatement := `UPDATE cohorts SET name=$2 WHERE id=$1 RETURNING id, name`
	var updatedID int
	var updatedName string
	err = db.QueryRow(sqlStatement, id, name).Scan(&updatedID, &updatedName)
	if err != nil {
		return cohortRecord, err
	}

	record := CohortRecord{
		ID:   updatedID,
		Name: updatedName,
	}
	cohortRecord = append(cohortRecord, record)

	return cohortRecord, nil
}

// Delete
func DeleteCohort(id int) (int64, error) {
	var deleteCount int64
	db, err := connectDB()
	if err != nil {
		return deleteCount, err
	}
	defer db.Close()

	sqlStatement := `DELETE FROM cohorts WHERE id=$1`
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
