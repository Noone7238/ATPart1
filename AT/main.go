package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//Item is the struct that contains the json fields. Simple Name, Description and ID for now.
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

//SampleHandler contains the internal data structure. Went with map, although may need to swap to array as requested.
type SampleHandler struct {
	items map[string]Item
}

//NewSampleHandle returns a pointer to a handler that has some sample items to populate at start
func NewSampleHandle() *SampleHandler {
	return &SampleHandler{
		items: map[string]Item{
			"1": Item{
				Name:        "Taco",
				Description: "A food item",
				ID:          "1",
			},
			"2": Item{
				Name:        "Paper Plane",
				Description: "A folded piece of paper, in the shape of a plane",
				ID:          "2",
			},
		},
	}
}

//selector a simple funtion that takes in the request and decides how the request should be handled
func (sh *SampleHandler) selector(w http.ResponseWriter, r *http.Request) {
	//A selector. For now I default with get all, maybe in future I can default a blurb to the screen on how to use. Like a sample help printout.
	switch r.Method {
	case "GET":
		sh.get(w, r)
	case "POST":
		sh.add(w, r)
	case "DELETE":
		sh.delete(w, r)
	default:
		sh.getall(w, r)
	}

}

//get either get's the passed in ID or falls back to calling getall
func (sh *SampleHandler) get(w http.ResponseWriter, r *http.Request) {
	//split up request, isolate ID
	parts := strings.Split(r.URL.String(), "/")

	//they either formatted it incorrectly or didn't add a slash at end, just give all for now.
	if len(parts) != 3 {
		sh.getall(w, r)
		return
	}

	//there is a slash at the end but nothing was specified. so give all.
	if len(parts[2]) == 0 {
		sh.getall(w, r)
		return
	}

	//see if the item can be found, if not write out error and return
	_, ok := sh.items[parts[2]]
	if !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("ID not found"))
		return
	}

	//Marshall happy format, happy life
	jsonByte, err := json.Marshal(sh.items[parts[2]])
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	//Time to write out the json bytes
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

//getall returns all stored items currently.
func (sh *SampleHandler) getall(w http.ResponseWriter, r *http.Request) {
	items := make([]Item, len(sh.items))

	i := 0
	for _, item := range sh.items {
		items[i] = item
		i++
	}

	//Marshall
	jsonByte, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

// delete a function that removes an item from the internal list.
func (sh *SampleHandler) delete(w http.ResponseWriter, r *http.Request) {
	//cut the passed in request into pieces, isolating the ID
	parts := strings.Split(r.URL.String(), "/")

	//make sure it is formated correctly, can potientally expand error checking
	if len(parts) != 3 {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Format incorrect"))
		return
	}

	//see if the requested ID is in the list. If not let them know it isn't. May want to let them know which ones are in the list in the future.
	_, ok := sh.items[parts[2]]
	if !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("ID Not Found"))
		return
	}

	//We've gotten this far, so item can be deleted
	delete(sh.items, parts[2])
	w.WriteHeader(http.StatusOK)
}

func (sh *SampleHandler) add(w http.ResponseWriter, r *http.Request) {
	//split up request, isolate ID
	parts := strings.Split(r.URL.String(), "/")

	//they either formatted it incorrectly or didn't add a slash at end, just give all for now.
	if len(parts) != 3 {
		sh.addnew(w, r)
		return
	}

	var ID = parts[2]

	//there is a slash at the end but nothing was specified. so give all.
	if len(ID) == 0 {
		sh.addnew(w, r)
		return
	}

	//see if the item can be found, if not write out error and return
	_, ok := sh.items[ID]
	if !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("ID not found"))
		return
	}

	//The ID was found, see what fields are there and update them
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	//store in temp item
	var i Item
	err = json.Unmarshal(bodyBytes, &i)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	// First we get a "copy" of the entry
	if entry, ok := sh.items[ID]; ok {
		// Then we modify the copy
		if len(i.Name) != 0 {
			entry.Name = i.Name
		}
		if len(i.Description) != 0 {
			entry.Description = i.Description
		}
		// Then we reassign map entry
		sh.items[ID] = entry
	}

	w.WriteHeader(http.StatusOK)
}

//add a function for adding an item to the internal list.
func (sh *SampleHandler) addnew(w http.ResponseWriter, r *http.Request) {
	//grab the json bytes and make sure there are no errors
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	//make a temp item to store the unmarshaled json bytes into
	var i Item
	err = json.Unmarshal(bodyBytes, &i)
	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	//if all goes well, add the item in the map.
	sh.items[i.ID] = i
	w.WriteHeader(http.StatusOK)
}

func main() {
	samhandler := NewSampleHandle()

	http.HandleFunc("/sample/", samhandler.selector)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
