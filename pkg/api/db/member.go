package db

import (
	model "github.com/StellarNotions/open-sesame/pkg/api/models"
)

var people []model.Person

// Insert allows populating database
func Insert(person model.Person) {
	people = append(people, person)
}

// Get returns the whole database
func Get() []model.Person {
	return people
}
