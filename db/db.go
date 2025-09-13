package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./attendance.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create Students table
	createStudentTable := `
    CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);`
	_, err = db.Exec(createStudentTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create Attendance table
	createAttendanceTable := `
	CREATE TABLE IF NOT EXISTS attendance (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    student_id INTEGER,
    date TEXT,
    status TEXT,
    FOREIGN KEY(student_id) REFERENCES students(id)
);`

	_, err = db.Exec(createAttendanceTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
