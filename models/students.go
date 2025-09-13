package models

import (
	"bufio"
	"database/sql"
	"fmt"
	"strings"
)

// Add new student
func AddStudent(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter student first name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter student last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	result, err := db.Exec("INSERT INTO students (first_name, last_name) VALUES (?, ?)", firstName, lastName)
	if err != nil {
		fmt.Println("❌ Error adding student:", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("✅ Student added successfully! (Could not fetch ID)")
		return
	}

	fmt.Printf("✅ Student added successfully!")
	fmt.Println()
	fmt.Printf("Student's id: %d ", id)
}

// Delete student
func DeleteStudent(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter student ID to delete: ")
	var studentID int
	fmt.Scanln(&studentID)

	_, err := db.Exec("DELETE FROM students WHERE id = ?", studentID)
	if err != nil {
		fmt.Println("❌ Error deleting student:", err)
		return
	}

	fmt.Println("✅ Student deleted successfully!")
}

// List all students
func ListStudents(db *sql.DB) {
	rows, err := db.Query("SELECT id, first_name, last_name FROM students")
	if err != nil {
		fmt.Println("❌ Error fetching students:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Students:")
	for rows.Next() {
		var id int
		var firstName, lastName string
		rows.Scan(&id, &firstName, &lastName)
		fmt.Printf("ID: %d, Name: %s %s\n", id, firstName, lastName)
	}
}
