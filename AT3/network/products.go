package network

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//validator temp spot
//var v = validator.New()

//Item is the struct that contains the json fields. Simple Name, Description.
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}

//echo validate, use custom interface
type SampleValidator struct {
	validator *validator.Validate
}

//Validate validates sample request body
func (s *SampleValidator) Validate(i interface{}) error {
	return s.validator.Struct(i)
}

//Set up some "samples"
type SampleHandler struct {
	samples map[int]Item
	e       *echo.Echo
}

func NewSampleHandle() (*SampleHandler, *echo.Echo) {
	var s = &SampleHandler{
		samples: map[int]Item{
			1: Item{
				Name:        "Taco",
				Description: "A food item",
				ID:          1,
			},
			2: Item{
				Name:        "Paper Plane",
				Description: "A folded piece of paper, in the shape of a plane",
				ID:          2,
			},
		},
		e: echo.New(),
	}
	return s, s.e
}

func (sh *SampleHandler) FindValidID() int {
	i := len(sh.samples) + 1
	_, ok := sh.samples[i]
	for ok {
		i++
		_, ok = sh.samples[i]
	}
	return i

}

//get all samples
func (sh *SampleHandler) GetSamples(c echo.Context) error {
	return c.JSON(http.StatusOK, sh.samples)
}

//get from samples, with ID
func (sh *SampleHandler) GetSample(c echo.Context) error {
	//capture id parameter
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	//check if id is in the map
	p, ok := sh.samples[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "sample not found")
	}

	//must of found it, print time
	return c.JSON(http.StatusOK, p)
}

//Post method for appending to the end of the samples map
func (sh *SampleHandler) CreateSample(c echo.Context) error {
	//make the struct for capturing the body of request, also maps out expected json input
	type body struct {
		//these tags in validate says, this is a reuired field and must be a minimum of 4 characters
		Name        string `json:"name" validate:"required,min=3"`
		Description string `json:"description" validate:"required"`
	}
	var v = validator.New()
	//bind it into your reqBody, if there is err yeet it
	var reqBody body
	//Other way to do validator, using echo
	sh.e.Validator = &SampleValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		fmt.Println("1")
		return err
	}

	//Alternate validator this is using echo's validator
	if err := c.Validate(reqBody); err != nil {
		fmt.Println("2")
		return err
	}

	//add to end of list, make sure not to overwrite anything
	i := sh.FindValidID()
	sample := Item{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		ID:          i,
	}

	sh.samples[i] = sample

	return c.JSON(http.StatusOK, sample)
}

//put method, update a sample
func (sh *SampleHandler) UpdateSample(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	//make sure ID is in map
	_, ok := sh.samples[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "sample not found")
	}

	//Since it's been found, validate new data is ok.
	type body struct {
		//these tags in validate says, this is a reuired field and must be a minimum of 4 characters
		Name        string `json:"name" validate:"omitempty,min=3"`
		Description string `json:"description" validate:"omitempty,min=3"`
	}
	var v = validator.New()

	//bind it into your reqBody, if there is err yeet it then validate
	var reqBody body
	sh.e.Validator = &SampleValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}

	//Extra validation, only override what was put in
	var n string
	if reqBody.Name == "" {
		n = sh.samples[pID].Name
	} else {
		n = reqBody.Name
	}

	var d string
	if reqBody.Description == "" {
		d = sh.samples[pID].Description
	} else {
		d = reqBody.Description
	}

	//if we are this far, pop the item into the map and override old values
	sample := Item{
		Name:        n,
		Description: d,
		ID:          pID,
	}

	sh.samples[pID] = sample
	return c.JSON(http.StatusOK, sample)
}

//delete method, remove a sample
func (sh *SampleHandler) DeleteSample(c echo.Context) error {
	//make sure we can capture the ID
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	//make sure ID is in map
	_, ok := sh.samples[pID]
	if !ok {
		return c.JSON(http.StatusNotFound, "sample not found")
	}

	delete(sh.samples, pID)
	return c.JSON(http.StatusOK, sh.samples)
}
