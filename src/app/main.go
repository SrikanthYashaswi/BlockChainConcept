package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct {
	from   int
	to     int
	amount float32
	next   *Transaction
}
type Book struct {
	name     string
	capacity float32
	used     float32
	head     *Transaction
	genesis  Transaction
}

func main() {
	var ind Book
	initBook(&ind, "IND", 21000000.00)
	lotteryPayments(&ind)
	printBook(&ind)
}

// System Functions START
func initBook(book *Book, name string, capacity float32) {
	book.name = name
	book.capacity = capacity
	book.used = 0
	genesis(book)
}

func genesis(book *Book) {
	book.genesis = pay(0, 0, 0)
	book.head = &book.genesis
}

func pay(from int, to int, amount float32) Transaction {
	return Transaction{from, to, amount, nil}
}

func recordPay(book *Book, trans Transaction) {
	book.head.next = &trans
	book.head = &trans
	book.used += trans.amount
}

func printBook(book *Book) {
	fmt.Printf("--- %s [%f] [%f]---\n", book.name, book.capacity, book.used)
	var trans = &book.genesis
	for trans != nil {
		fmt.Printf("from: %d , to: %d, [%f]\n", trans.from, trans.to, trans.amount)
		trans = trans.next
	}
}

// System Function END

// Testing Function START

func lotteryPayments(book *Book) {
	var maxTransaction = 51
	var maxUsers = 4
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < maxTransaction; i++ {
		from, to := uniqCombination(r1, maxUsers)
		recordPay(book, pay(from, to, r1.Float32()*1000))
	}
}

func recipients(r *rand.Rand, users int) (int, int) {
	return r.Intn(users), r.Intn(users)
}

func uniqCombination(r *rand.Rand, users int) (int, int) {
	from, to := recipients(r, users)
	for from == to {
		from, to = recipients(r, users)
	}
	return from, to
}

// Testing Function END
