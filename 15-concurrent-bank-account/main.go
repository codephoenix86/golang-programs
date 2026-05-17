package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (account *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	account.mu.Lock()
	defer account.mu.Unlock()
	defer wg.Done()
	account.balance += amount
}

func main() {
	var account BankAccount
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go account.Deposit(1, &wg)
	}
	wg.Wait()
	fmt.Printf("Final Balance: %d\n", account.balance)
}
