package models

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Mark attendance for a student
func MarkAttendance(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter student ID: ")
	var studentID int
	fmt.Scanln(&studentID)

	fmt.Print("Enter date (YYYY-MM-DD): ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	fmt.Print("Enter status (Present/Absent): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	_, err := db.Exec("INSERT INTO attendance (student_id, date, status) VALUES (?, ?, ?)", studentID, date, status)
	if err != nil {
		fmt.Println("❌ Error marking attendance:", err)
		return
	}

	fmt.Println("✅ Attendance marked successfully!")
}

// View all attendance records
func ViewAttendanceByDay(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date (YYYY-MM-DD): ")
	dateInput, _ := reader.ReadString('\n')
	dateInput = strings.TrimSpace(dateInput)

	rows, err := db.Query(`
		SELECT s.id, s.first_name, s.last_name, 
		       IFNULL(a.status, 'Absent') as status
		FROM students s
		LEFT JOIN attendance a ON s.id = a.student_id AND a.date = ?
	`, dateInput)
	if err != nil {
		fmt.Println("Error fetching attendance:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\nAttendance for %s:\n", dateInput)
	fmt.Println("ID\tFirst Name\tLast Name\tStatus")
	for rows.Next() {
		var id int
		var firstName, lastName, status string
		err := rows.Scan(&id, &firstName, &lastName, &status)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("%d\t%s\t\t%s\t\t%s\n", id, firstName, lastName, status)
	}
}
