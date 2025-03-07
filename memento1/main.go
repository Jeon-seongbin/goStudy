package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = %d current = %d", b.balance, b.current)
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{Balance: balance})
	return b
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	return m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance -= m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.balance > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(10)
	ba.Deposit(20)
	ba.Deposit(30)

	ba.Undo()
	fmt.Println(ba)

	ba.Undo()
	fmt.Println(ba)

	ba.Redo()
	fmt.Println(ba)
}
