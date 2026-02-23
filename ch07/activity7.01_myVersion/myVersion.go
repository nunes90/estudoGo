// activity7.01 - calculating pay and performance review
// myVersion
package main

import (
	"fmt"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]any
}

func (d Developer) FullName() string {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) FullName() string {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	fullName := d.FullName()
	yearPay := d.HourlyRate * d.HoursWorkedInYear
	return fullName, yearPay
}

func (m Manager) Pay() (string, float64) {
	fullName := m.FullName()
	yearPay := m.Salary + (m.Salary * m.CommissionRate)
	return fullName, yearPay
}

func payDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("%s got paid $%.2f for the year\n", fullName, yearPay)
}

func review(d *Developer, category string, rating any) {
	/*Valid ratings
	 * "Excellent" - 5
		"Good" - 4
		"Fair" - 3
		"Poor" - 2
		"Unsatisfactory" - 1
	*/
	switch rating.(type) {
	case string, int:
		d.Review[category] = rating
	}
}

func performanceReview(d Developer) float64 {
	sum := 0.0
	for _, v := range d.Review {
		if s, ok := v.(string); ok {
			switch s {
			case "Excellent":
				sum += 5
			case "Good":
				sum += 4
			case "Fair":
				sum += 3
			case "Poor":
				sum += 2
			case "Unsatisfactory":
				sum += 1
			}
		}
		if n, ok := v.(int); ok {
			sum += float64(n)
		}
	}
	averageRating := sum / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.FullName(), averageRating)
	return averageRating
}

func main() {
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: make(map[string]any)}

	review(&d, "WorkQuality", 5)
	review(&d, "TeamWork", 2)
	review(&d, "Communication", "Poor")
	review(&d, "Problem-solving", 4)
	review(&d, "Dependability", "Unsatisfactory")

	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	performanceReview(d)
	payDetails(d)
	payDetails(m)
}
