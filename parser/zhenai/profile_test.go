package zhenai

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile("profile_test_data.html", contents, "profile_test_data")
	log.Printf("%s", result.Items)
}
