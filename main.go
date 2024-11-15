package main

import (
	"fmt"
)

func main() {
	var temperature int
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temperature)
	fmt.Println(weather(temperature))
	var age int
	fmt.Print("Enter the age: ")
	fmt.Scan(&age)
	var department string
	fmt.Print("Enter the department name \n IT \n CSEAIML \n Non-TECHNICAL \n ------------- \n ")
	fmt.Scan(&department)
	ss := salary(age, department)
	if ss == "0" {
		fmt.Println("YOU ARE NOT APLICABLE FOR SALARY")
	} else {
		fmt.Println("YOUR SALARY IS ", ss)
	}
	createVehicle()
	var x int
	fmt.Println("Enter the number to check even or not")
	fmt.Scan(&x)
	fmt.Println(" IT IS ", isEven(x), "THAT THE NUMBER IS EVEN")

	// Example usage
	fmt.Println("Enter the Operator")
	var operator string
	fmt.Scan(&operator)
	fmt.Println("Enter the Two Numbers")
	var a, b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	ans, err := calculator(operator, a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", ans)
	}

	fmt.Println("Enter the number to get sum")
	var a1 int
	fmt.Scan(&a1)
	fmt.Println(Sum(a1))
	var one, two int
	fmt.Println("Enter the Two Numbers two swap")
	fmt.Scan(&one)
	fmt.Scan(&two)
	fmt.Println("Before swap a =", one, " b =", two)
	swapNo(&one, &two)
	fmt.Println("After swap a =", one, " b =", two)

	var v BankAccount
	fmt.Println("Enter the account number")
	fmt.Scan(&v.accountNumber)

	fmt.Println("Enter the bank name")
	fmt.Scan(&v.name)

	fmt.Println("Enter the balance ")
	fmt.Scan(&v.balance)
	fmt.Println("Enter the ifsc code ")
	fmt.Scan(&v.ifsc)
	display(v)
	
}
