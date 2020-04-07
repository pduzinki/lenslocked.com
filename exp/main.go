package main

import (
	"fmt"

	"lenslocked.com/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	// Create a user
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// Update a user
	user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	// NOTE: You may need to update the query code a bit as well
	foundUser, err := us.ByEmail("michael@dundermifflin.com")
	if err != nil {
		panic(err)
	}
	fmt.Print("--------\n")
	fmt.Println(foundUser)

	// Delete a user
	if err := us.Delete(foundUser.ID); err != nil {
		panic(err)
	}
	// Verify the user is deleted
	_, err = us.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}
}
