package domain

import "testing"

func TestUser_Validaiton(t *testing.T) {
	_, err := NewUser(&User{
		FirstName:   "Asad",
		LastName:    "sohel",
		Email:       "asad@me.com",
		PhoneNumber: "01777",
		Password:    "123",
	})

	if err != nil {
		t.Fatalf("expeced error: %+v", err)
	}

	_, err = NewUser(&User{
		FirstName:   "a",
		LastName:    "sohel",
		Email:       "asad@me.com",
		PhoneNumber: "01777",
		Password:    "123",
	})

	if err != nil {
		t.Fatalf("expeced error: %+v", err)
	}
}
