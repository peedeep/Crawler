package model

import "encoding/json"

type Profile struct {
	Name string
	Gender string
	Age string
	Height string 
	Weight int
	Income string
	Marriage string
	Education string
	Occupation string
	Hokou string
	Xinzou string
	House string
	Car string
}

func FromJsonOjb(o interface{}) (Profile, error) {
	var profile Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)
	return profile, err
}
