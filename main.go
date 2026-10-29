package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"


"studentresult/result"
"studentresult/student"


)

func main() {
reader := bufio.NewReader(os.Stdin)


fmt.Println("===================================")
fmt.Println("   STUDENT RESULT PROCESSING")
fmt.Println("===================================")

// Read student name
fmt.Print("Enter Student Name: ")
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name)

// Read roll number
fmt.Print("Enter Roll Number: ")
rollNo, _ := reader.ReadString('\n')
rollNo = strings.TrimSpace(rollNo)

// Read marks for five subjects
var marks []float64

for i := 1; i <= 5; i++ {
	for {
		fmt.Printf("Enter mark for Subject %d: ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		mark, err := strconv.ParseFloat(input, 64)

		if err != nil || mark < 0 || mark > 100 {
			fmt.Println("Invalid mark! Enter a number between 0 and 100.")
			continue
		}

		marks = append(marks, mark)
		break
	}
}

// Create student record
s := student.Student{
	Name:   name,
	RollNo: rollNo,
	Marks:  marks,
}

// Calculate results
total := result.CalculateTotal(s.Marks)
average := result.CalculateAverage(s.Marks)
grade := result.CalculateGrade(average)

// Display result
fmt.Println("\n===================================")
fmt.Println("         STUDENT RESULT")
fmt.Println("===================================")
fmt.Println("Student Name :", s.Name)
fmt.Println("Roll Number  :", s.RollNo)
fmt.Printf("Total Marks  : %.2f / 500\n", total)
fmt.Printf("Average      : %.2f\n", average)
fmt.Println("Grade        :", grade)
fmt.Println("===================================")


}
