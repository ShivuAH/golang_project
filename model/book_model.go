package model

type BookModel struct {
	BookId      string `json:"bookid"`
	BookName    string `json:"bookname"`
	CategoryId  string `json:"categarid"`
	TotalCopies string `json:"totalcopies"`
}
