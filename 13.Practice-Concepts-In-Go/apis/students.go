package apis

import (
	"math/rand"
	"strconv"
	"time"
)


type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	DOB string `json:"date_of_birth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentProfile struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Group string `json:"group"`
	Hobbies []string `json:"hobbies"`
	Student Student `json:"student"`
}

var students [] StudentProfile 

func InitializeWithStaticData() *StudentProfile {
	data := StudentProfile {
		ID: strconv.Itoa(rand.Intn(10000)),
		Email: "example@gmail.com",
		Group: "engineering",
		Hobbies: []string{"Coding", "Hiking", "Music"},
		Student: Student{
			Name:      "John Doe",
			Age:       22,
			Gender:    "Male",
			DOB:       "2002-01-01",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, 
	}

	// push into global slice
	students = append(students, data)
	return &data
}

