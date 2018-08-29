package main

type Transaction struct {
	from   int32
	to     int32
	amount float32
	next   *Transaction
}
type Book struct {
	name string
	head *Transaction
}

func main() {
	var ind Book
	initBook(&ind, "IND")
}

func initBook(book *Book, name string) {
	book.name = name
	genesis(book)
}

func genesis(book *Book) {
	book.head = pay(0, 0, 0)

}

func pay(from int32, to int32, amount float32) *Transaction {
	return &Transaction{from, to, amount, nil}
}
