package pointers_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Everything struct {
	Library
	Catalog
}

type Library []*Book

type Catalog []*Book

type Book struct {
	Title  string
	Author *Author
}

type Author struct {
	Name string
}

func TestSerializtion(t *testing.T) {
	a1 := &Author{"Enrique Vila-Matas"}
	b1 := &Book{"El mal de Montano", a1}
	b2 := &Book{"Paris no se acaba nunca", a1}
	l1 := Library{b1, b2}
	c1 := Catalog{b1}
	before := &Everything{l1, c1}

	encoded, _ := json.Marshal(before)

	after := &Everything{}
	json.Unmarshal(encoded, after)

	assert.Equal(t, after, before)
	assert.Equal(t, after.Library[0].Title, before.Library[0].Title)

	// note that attention, although the before author is equal to the after author, in reality there are now two different structs in memory
	var beforeAuthorPtr *Author = before.Library[0].Author
	var afterAuthorPtr *Author = after.Library[0].Author

	var beforeAuthor Author = *beforeAuthorPtr
	var afterAuthor Author = *afterAuthorPtr

	assert.Equal(t, afterAuthor, beforeAuthor)
	assert.Equal(t, afterAuthor.Name, beforeAuthor.Name)

	afterAuthor.Name = "E. Vila-Matas"
	assert.NotEqual(t, afterAuthor.Name, beforeAuthor.Name)
}
