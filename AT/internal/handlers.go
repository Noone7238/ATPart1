package network

import (
	"gitstuff/AT/internal/structures"

	"github.com/labstack/echo/v4"
)

//Set up some sample items
type ItemHandler struct {
	items map[int]structures.Item
	e     *echo.Echo
}

func NewItemHandle() (*ItemHandler, *echo.Echo) {
	var s = &ItemHandler{
		items: map[int]structures.Item{
			1: structures.Item{
				Name:        "Taco",
				Description: "A food item",
				ID:          1,
			},
			2: structures.Item{
				Name:        "Paper Plane",
				Description: "A folded piece of paper, in the shape of a plane",
				ID:          2,
			},
		},
		e: echo.New(),
	}
	return s, s.e
}

func (sh *ItemHandler) FindValidID() int {
	i := len(sh.items) + 1
	_, ok := sh.items[i]
	for ok {
		i++
		_, ok = sh.items[i]
	}
	return i

}
