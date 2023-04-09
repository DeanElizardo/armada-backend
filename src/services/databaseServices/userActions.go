package databaseServices

type UsersRecord struct {
	Uuid      string
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsAdmin   bool
}

// Create

// Read
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

func GetCourseUsers(courseId int) ([]UsersRecord, error) {
	var userRecords []UsersRecord

	db, err := connectDB()
	if err != nil {
		return userRecords, err
	}
	defer db.Close()

	sqlStatement := `
	SELECT uuid, username, firstname, lastname, isadmin FROM users
	JOIN users_courses ON users.uuid = users_courses.userId
	JOIN courses ON users_courses.courseId = courses.id
	WHERE courses.id = $1
	`
	rows, err := db.Query(sqlStatement, courseId)
	if err != nil {
		return userRecords, err
	}
	for rows.Next() {
		var uuid string
		var username string
		var email string
		var firstName string
		var lastName string
		var isAdmin bool
		err = rows.Scan(&uuid, &username, &email, &firstName, &lastName, &isAdmin)
		if err != nil {
			return userRecords, err
		}
		record := UsersRecord{
			Uuid:      uuid,
			Username:  username,
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			IsAdmin:   isAdmin,
		}
		userRecords = append(userRecords, record)
	}
	return userRecords, nil
}

// Update

// Delete
