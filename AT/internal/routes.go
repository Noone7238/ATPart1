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
		//log
		return err
	}

	if ih.IsValidID(pID) {
		return c.JSON(http.StatusNotFound, "item out of range")
	}

	//must of found it, print time
	return c.JSON(http.StatusOK, ih.items[pID])
}

//post method, partial update
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
		//log
		return err
	}

	//Alternate validator this is using echo's validator
	if err := c.Validate(reqBody); err != nil {
		//log
		return err
	}

	//add to end of list, make sure not to overwrite anything
	item := structures.Item{
		Name:        reqBody.Name,
		Description: reqBody.Description,
	}

	ih.items = append(ih.items, item)

	return c.JSON(http.StatusOK, item)
}

//put method, add a new item
func (ih *ItemHandler) UpdateItem(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//log
		return err
	}

	if ih.IsValidID(pID) {
		//log
		return c.JSON(http.StatusNotFound, "item out of range")
	}

	//Since it's been found, validate new data is ok.
	type body struct {
		//these tags in validate says, this is a reuired field and must be a minimum of 4 characters
		Name        string `json:"name" validate:"omitempty,min=3"`
		Description string `json:"description" validate:"omitempty,min=3"`
	}

	//bind it into your reqBody
	var reqBody body
	//Other way to do validator, using echo
	ih.e.Validator = structures.NewSampValidator()
	if err := c.Bind(&reqBody); err != nil {
		//log
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		//log
		return err
	}

	//Extra validation, only override what was put in
	if reqBody.Name != "" {
		ih.items[pID].Name = reqBody.Name
	}

	if reqBody.Description != "" {
		ih.items[pID].Description = reqBody.Description
	}

	return c.JSON(http.StatusOK, ih.items[pID])
}

//delete method, remove a item
func (ih *ItemHandler) DeleteItem(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//log
		return err
	}

	//make sure ID is in slice
	if ih.IsValidID(pID) {
		//log
		return c.JSON(http.StatusNotFound, "item out of range")
	}

	ih.Remove(pID)
	return c.JSON(http.StatusOK, ih.items)
}
