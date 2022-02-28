package network

import (
	"gitstuff/AT/internal/structures"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//get method, all items
func (ih *ItemHandler) GetItems(c echo.Context) error {
	return c.JSON(http.StatusOK, ih.items)
}

//get method, from items, with ID
func (ih *ItemHandler) GetItem(c echo.Context) error {
	//capture id parameter
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	//check if id is in the map
	p, ok := ih.items[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	//must of found it, print time
	return c.JSON(http.StatusOK, p)
}

//post method, partial update, with ID
func (ih *ItemHandler) CreateItem(c echo.Context) error {
	//make the struct for capturing the body of request, also maps out expected json input
	type body struct {
		//these tags in validate says, this is a reuired field and must be a minimum of 4 characters
		Name        string `json:"name" validate:"required,min=3"`
		Description string `json:"description" validate:"required"`
	}

	//bind it into your reqBody, if there is err yeet it
	var reqBody body
	//Other way to do validator, using echo
	ih.e.Validator = structures.NewSampValidator()
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	//Alternate validator this is using echo's validator
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	//add to end of list, make sure not to overwrite anything
	i := ih.FindValidID()
	item := structures.Item{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		ID:          i,
	}

	ih.items[i] = item

	return c.JSON(http.StatusOK, item)
}

//put method, add a new item
func (ih *ItemHandler) UpdateItem(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	//make sure ID is in map
	_, ok := ih.items[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	//Since it's been found, validate new data is ok.
	type body struct {
		//these tags in validate says, this is a reuired field and must be a minimum of 4 characters
		Name        string `json:"name" validate:"omitempty,min=3"`
		Description string `json:"description" validate:"omitempty,min=3"`
	}

	//bind it into your reqBody, if there is err yeet it
	var reqBody body
	//Other way to do validator, using echo
	ih.e.Validator = structures.NewSampValidator()
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	//Extra validation, only override what was put in
	var n string
	if reqBody.Name == "" {
		n = ih.items[pID].Name
	} else {
		n = reqBody.Name
	}

	var d string
	if reqBody.Description == "" {
		d = ih.items[pID].Description
	} else {
		d = reqBody.Description
	}

	//if we are this far, pop the item into the map and override old values
	item := structures.Item{
		Name:        n,
		Description: d,
		ID:          pID,
	}

	ih.items[pID] = item
	return c.JSON(http.StatusOK, item)
}

//delete method, remove a item
func (ih *ItemHandler) DeleteItem(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	//make sure ID is in map
	_, ok := ih.items[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	delete(ih.items, pID)
	return c.JSON(http.StatusOK, ih.items)
}
