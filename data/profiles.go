package data

/*
* data types and corresponding functions used by RESTful services
 */

import (
	"encoding/json"
	"io"
	"time"
)

// This struct represents a single profile:
// Fields need to be in capital letters! For camelcase json tags are added
type Profile struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	X           float32 `json:"x"`
	Y           float32 `json:"y"`
	Z           float32 `json:"z"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Custom type which wraps profile list
type Profiles []*Profile

// Marshal list into JSON
func (p *Profiles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w) // A bit faster than Marshal
	return e.Encode(p)
}

// Marshal list into JSON
func (p *Profile) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r) // A bit faster than Marshal
	return e.Decode(p)
}

// Return all saved profiles
func GetProfiles() Profiles {
	return productList
}

func AddProfile(p *Profile) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// A sample list of profiles for testing
var productList = []*Profile{
	&Profile{
		ID:          1,
		Name:        "TP1_NotValid",
		Description: "One Step at a time",
		X:           2.45,
		Y:           2.45,
		Z:           2.45,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Profile{
		ID:          2,
		Name:        "TP2_NotValid",
		Description: "all the things",
		X:           2.45,
		Y:           2.45,
		Z:           2.45,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
