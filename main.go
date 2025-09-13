package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"attendance/db"
	"attendance/models"
	"attendance/reports"
)

func main() {
	// DB connection
	database := db.InitDB()
	defer database.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Student Attendance System ---")
		fmt.Println("1) Add Student")
		fmt.Println("2) List Students")
		fmt.Println("3) Mark Attendance of a student")
		fmt.Println("4) View Attendance by day")
		fmt.Println("5) Monthly Report")
		fmt.Println("6) Delete Student")
		fmt.Println("7) Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			models.AddStudent(database, reader)
		case "2":
			models.ListStudents(database)
		case "3":
			models.MarkAttendance(database, reader)
		case "4":
			models.ViewAttendanceByDay(database)
		case "5":
			reports.MonthlyReport(database, reader)
		case "6":
			models.DeleteStudent(database, reader)
		case "7":
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
