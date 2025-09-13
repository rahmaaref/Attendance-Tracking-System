package reports

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// Monthly report: Present vs Absent
func MonthlyReport(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter student ID: ")
	var studentID int
	fmt.Scanln(&studentID)

	fmt.Print("Enter month (YYYY-MM): ")
	month, _ := reader.ReadString('\n')
	month = strings.TrimSpace(month)

	query := `
		SELECT COUNT(*) FILTER (WHERE status='Present'),
		       COUNT(*) FILTER (WHERE status='Absent')
		FROM attendance
		WHERE student_id=? AND date LIKE ?;
	`

	row := db.QueryRow(query, studentID, month+"-%")

	var presentCount, absentCount int
	err := row.Scan(&presentCount, &absentCount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n--- Monthly Report for Student ID %d (%s) ---\n", studentID, month)
	fmt.Printf("Present: %d days | Absent: %d days\n", presentCount, absentCount)
}
