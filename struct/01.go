package main

import "fmt"

type order struct {
	ID     int
	Status string
	Amount float64
}

func NewOrder(id int, status string, amount float64) order {
	return order{
		ID:     id,
		Status: status,
		Amount: amount,
	}
}

func main() {

	fmt.Println("func return struct ")
	NewOrder := NewOrder(123, "pending", 20.00)
	fmt.Println(NewOrder)
	fmt.Println(NewOrder.ID)
	fmt.Println(NewOrder.Amount)
	fmt.Println(NewOrder.Status)
}
