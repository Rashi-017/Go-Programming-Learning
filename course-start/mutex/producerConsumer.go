package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalPizza = 10

var PizzaMade, PizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func MakePizza(Pizzanumber int) *PizzaOrder {
	Pizzanumber++
	if Pizzanumber <= totalPizza {
		delay := rand.Intn(5) + 1
		fmt.Printf("recieved order %d\n", Pizzanumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			PizzaFailed++
		} else {
			PizzaMade++
		}
		total++
		fmt.Printf("making pizza %d and delay %d ", Pizzanumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd < 2 {
			fmt.Printf("we ran out of ingredients for pizza %d", Pizzanumber)
		} else if rnd <= 4 {
			fmt.Printf("cook quit %d", Pizzanumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order %d is ready", Pizzanumber)
		}
		p := PizzaOrder{
			pizzaNumber: Pizzanumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: Pizzanumber,
	}
}
func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentpizza := MakePizza(i)
		if currentpizza != nil {
			i = currentpizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentpizza:
			case quitchan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitchan)
				return
			}
		}
	}
}
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
func producerConsumer() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("pizzarie is open ")
	fmt.Println("--------------------")

	pizzajob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	//producer run in background
	go pizzeria(pizzajob)

	for i := range pizzajob.data {
		NumberOfPizzas := 0
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				fmt.Println(i.message)
				fmt.Printf("order is of out of delivery%d\n", i.pizzaNumber)
			} else {
				fmt.Println(i.message)
				fmt.Println("customer gone made ")
			}
		} else {
			fmt.Println("done making pizzas ")
		}
	}
	fmt.Println("done for the day")

}
