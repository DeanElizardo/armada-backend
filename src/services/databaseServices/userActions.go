package databaseServices

type UsersRecord struct {
	Uuid      string
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsAdmin   bool
}

func GetAllUsersRecords() ([]UsersRecord, error) {
	var usersRecords []UsersRecord

	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uuid string
		var username string
		var email string
		var firstName string
		var lastName string
		var isAdmin bool
		err = rows.Scan(&uuid, &username, &email, &firstName, &lastName, &isAdmin)
		if err != nil {
			panic(err)
		}
		record := UsersRecord{uuid, username, email, firstName, lastName, isAdmin}
		usersRecords = append(usersRecords, record)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return usersRecords, nil
}
