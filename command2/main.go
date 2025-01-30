package main

type BankAccount struct {
	Balance int
}

func Deposit(ba *BankAccount, amount int) {
	ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	if ba.Balance >= amount {
		ba.Balance -= amount
	}
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

func main() {
	ba := &BankAccount{}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 100)
	})

	for _, cmd := range commands {
		cmd()
	}
}
