package network

import (
	"gitstuff/AT/internal/structures"

	"github.com/labstack/echo/v4"
)

//Set up some sample items
type ItemHandler struct {
	items []structures.Item
	e     *echo.Echo
}

func NewItemHandle() (*ItemHandler, *echo.Echo) {
	var s = &ItemHandler{
		e:     echo.New(),
		items: make([]structures.Item, 0),
	}

	var i1 = structures.Item{
		Name:        "Taco",
		Description: "A food item",
	}
	var i2 = structures.Item{
		Name:        "Paper Plane",
		Description: "A folded piece of paper, in the shape of a plane",
	}

	s.items = append(s.items, i1, i2)
	return s, s.e
}

func (ih *ItemHandler) IsValidID(id int) bool {

	return (id < 0) || (id > len(ih.items)-1)
}

func (ih *ItemHandler) Remove(i int) {
	ih.items = append(ih.items[:i], ih.items[i+1:]...)
}
