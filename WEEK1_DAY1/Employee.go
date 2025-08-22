package main

import "fmt"

type Employee interface {
	CalculateSalary() float64
	GetName() string
}

type FullTimeEmployee struct {
	Name          string
	MonthlySalary float64
}

func (fte FullTimeEmployee) CalculateSalary() float64 {
	return fte.MonthlySalary
}

func (fte FullTimeEmployee) GetName() string {
	return fte.Name
}

type Contractor struct {
	Name      string
	DailyRate float64
	WorkDays  int
}

func (c Contractor) CalculateSalary() float64 {
	return c.DailyRate * float64(c.WorkDays)
}

func (c Contractor) GetName() string {
	return c.Name
}

type Freelancer struct {
	Name        string
	HourlyRate  float64
	HoursWorked int
}

func (f Freelancer) CalculateSalary() float64 {
	return f.HourlyRate * float64(f.HoursWorked)
}

func (f Freelancer) GetName() string {
	return f.Name
}

func main() {
	ft := FullTimeEmployee{Name: "Alice", MonthlySalary: 15000}
	cont := Contractor{Name: "Bob", DailyRate: 150, WorkDays: 20}
	free := Freelancer{Name: "Charlie", HourlyRate: 100, HoursWorked: 20}

	employees := []Employee{ft, cont, free}

	fmt.Println("Employee Salaries:")
	for _, emp := range employees {
		fmt.Printf("%s (Type: %T): Salary = %.2f\n", emp.GetName(), emp, emp.CalculateSalary())
	}
}
