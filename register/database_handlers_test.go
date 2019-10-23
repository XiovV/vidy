package main

import "testing"

func TestInsertUserAndDoesUserExist(t *testing.T) {
	// First we insert a user into the database

	InsertUser(User{"testuser@user.com", "testpassword123"})

	// Then we call DoesUserExist() to check if the user we just inserted is inside the database
	doesUserExist := DoesUserExist("testuser@user.com")

	if doesUserExist == false {
		t.Errorf("User isn't in the databse")
	} else {
		RemoveUser("testuser@user.com")
	}
}
