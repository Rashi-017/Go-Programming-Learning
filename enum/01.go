package main

import "fmt"

type orderError struct {
	status string
}

func (e *orderError) Error() string {
	return "invalid order status " + e.status
}

type orderType string

const (
	Delivered  orderType = "Delivered"
	Pending    orderType = "Pending"
	Incomplete orderType = "Incomplete"
)

func IsValidOrderType(o orderType) bool {
	switch o {
	case Delivered:
		return true
	case Pending:
		return true
	case Incomplete:
		return true
	default:
		return false
	}

}
func validOrder(s orderType) (bool error) {
	if s != "Delivered" && s != "Pending" && s != "Incomplete" {
		return &orderError{status: string(s)}
	}
	return nil
}

func orderstatus(s orderType) {
	if !IsValidOrderType(s) {
		fmt.Println("enter valid type orderstatus")
	} else {
		fmt.Println("thank for orderstatus", s)
	}
}
func main() {
	var order orderType = "Pending"
	orderstatus(order)
	err := validOrder(order)

	if err != nil {
		fmt.Println("Error", err)
	}

}
