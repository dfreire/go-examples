package pointers_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Book struct {
	Title     string
	AuthorPtr *Author
}

type Author struct {
	Name string
}

func Test(t *testing.T) {
	author1 := &Author{"Enrique Vila-Matas"}
	book1 := &Book{"El mal de Montano", author1}
	assert.Equal(t, book1.AuthorPtr, author1)

	encoded, _ := json.Marshal(book1)

	book2 := &Book{}
	json.Unmarshal(encoded, book2)

	assert.Equal(t, author1, book2.AuthorPtr)
	assert.Equal(t, book1.AuthorPtr, book2.AuthorPtr)
	assert.Equal(t, book1.Title, book2.Title)
	assert.Equal(t, book1, book2)
}
