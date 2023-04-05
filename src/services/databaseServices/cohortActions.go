package databaseServices

type CohortRecord struct {
	ID   int64
	Name string
}

func GetAllCohortRecords() ([]CohortRecord, error) {
	var cohortRecords []CohortRecord

	db, err := connectDB()
	if err != nil {
		return cohortRecords, err
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM cohort`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return cohortRecords, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
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
