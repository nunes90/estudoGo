package author

import "fmt"

// Author represents an author of a book.
type Author struct {
	Name    string
	Contact string
}

func NewAuthor(name, contact string) *Author {
	return &Author{Name: name, Contact: contact}
}

func (a *Author) WriteChapter(chapterTitle string, content string) {
	fmt.Printf("Author %s is writing a chapter titled    '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) ReviewChapter(chapterTitle string, content string) {
	fmt.Printf("Author %s is reviewing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Author %s has finalized the chapter titled '%s'\n", a.Name, chapterTitle)
}
