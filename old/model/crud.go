package model

import (
	"fmt"
	"golang-app/view"
)

// AddEmployee - Add employee details to database
func AddEmployee(id string, name string) error {
	rows, err := dbConnection.Query("INSERT INTO EMPLOYEE VALUES($1, $2)", id, name)
	defer rows.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	return nil
}

// DeleteEmployee  - Delete Employe data
func DeleteEmployee(id string) error {
	rows, err := dbConnection.Query("DELETE FROM EMPLOYEE WHERE id=$1", id)
	defer rows.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	return nil
}

// GetAllEmployees - Get All Employees
func GetAllEmployees() ([]view.Employee, error) {
	rows, err := dbConnection.Query("SELECT * FROM EMPLOYEE")
	defer rows.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	employees := []view.Employee{}
	for rows.Next() {
		employee := view.Employee{}
		rows.Scan(&employee.ID, &employee.Name)
		employees = append(employees, employee)
	}
	fmt.Printf("Employees %v\n", employees)
	return employees, nil
}

// GetEmployeeByID  - Get employee detail by ID
func GetEmployeeByID(id string) ([]view.Employee, error) {
	rows, err := dbConnection.Query("SELECT * FROM EMPLOYEE WHERE id=$1", id)
	defer rows.Close()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	employees := []view.Employee{}
	for rows.Next() {
		employee := view.Employee{}
		rows.Scan(&employee.ID, &employee.Name)
		employees = append(employees, employee)
	}
	fmt.Printf("Employees %v\n", employees)
	return employees, nil
}
