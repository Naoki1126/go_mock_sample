package object

import "time"

type Book struct {
	Title       string    `json:title`
	Author      string    `json:author`
	Price       int       `json: price`
	ReleaseDate time.Time `json: release_date`
}

func NewBook(t string, a string, p int, r time.Time) *Book {
	return &Book{
		Title:       t,
		Author:      a,
		Price:       p,
		ReleaseDate: r,
	}
}

func (b *Book) Init() {
	b.Title = "sampleBook"
	b.Author = "著者がいません"
	b.Price = 100
	b.ReleaseDate = time.Date(2022, 4, 1, 1, 1, 1, 1, time.Local)
}
