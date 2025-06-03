package main

import (
	"fmt"
	"sync"
)

type Amount struct {
	balance int
	mu      sync.Mutex
}

func (a *Amount) Desposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.mu.Lock()
	a.balance += amount
	a.mu.Unlock()

}

func (a *Amount) withdrawal(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	if a.balance >= amount {
		a.balance -= amount
	} else {
		fmt.Println("Insufficient funds")
	}
	a.mu.Unlock()
}

func mutex() {
	var wg sync.WaitGroup
	amt := Amount{balance: 20000}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go amt.Desposit(0, &wg)
	}

	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go amt.withdrawal(20000, &wg)
	}

	wg.Wait()
	fmt.Println("Final balance:", amt.balance)

}
