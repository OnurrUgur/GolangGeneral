package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string`json:"first_name"`
	LastName string`json:"last_name"`
	HairColor string`json:"hair_color"`
	HasDog bool`json:"has_dog"`

}

func second() {

	myJson := `
[
	{
		"first_name" : "Clark",
		"last_name" : "Json",
		"hair_color" : "black",
		"has_dog" : true
	},
	{
		"first_name" : "Bruce",
		"last_name" : "Wayne",
		"hair_color" : "brown",
		"has_dog" : false
	}

]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println(err)
	}

	log.Printf("unmarshalled: %v" , unmarshalled )

	var mySlice []Person
	var m1 Person

	m1.FirstName = "Emily"
	m1.LastName = "Last"
	m1.HairColor = "Black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Onur"
	m2.LastName = "Uğur"
	m2.HairColor = "Brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice,"","   ")
	if err != nil {
		log.Println("erorr marshalling is",err)
	}

	fmt.Println(string(newJson))



}