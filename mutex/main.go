package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposite(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount

	fmt.Println("Deposited: ", amount)
}

func (b *BankAccount) WithDraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance -= amount

	fmt.Println("Withdrawed: ", amount)
}

func (b *BankAccount) Balance() {
	fmt.Println("balance: ", b.balance)
}

func main() {
	var wg sync.WaitGroup
	var account = BankAccount{balance: 100}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)

			account.Deposite(amount)
		}(i + 1)
	}

	wg.Wait()
	account.Balance()
}
