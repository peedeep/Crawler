package model

import "encoding/json"

type Movie struct {
	Name string
	Time string
}

func fromJsonObj(o interface{}) (Movie, error) {
	var movie Movie
	bytes, err := json.Marshal(o)
	if err != nil {
		return movie, err
	}
	err = json.Unmarshal(bytes, &movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}
