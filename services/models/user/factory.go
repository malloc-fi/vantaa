package user

// DummyUsers or factory.DummyUsers creates a slice of valid Users
func DummyUsers() []User {
	users := []User{
		User{
			Name:     "admin",
			Email:    "admin@example.com",
			Password: "adminpass",
		},
		User{
			Name:     "johndoe",
			Email:    "jhondoe@example.com",
			Password: "johndoepass",
		},
		User{
			Name:     "foo",
			Email:    "foo@example.com",
			Password: "foopass",
		},
		User{
			Name:     "bar",
			Email:    "bar@example.com",
			Password: "barpass",
		},
	}
	return users
}

// DummyUser or factory.DummyUser creates 1 valid User
func DummyUser() User {
	return User{
		Name:     "unique_user",
		Email:    "unique@example.com",
		Password: "unique",
	}
}

// CreateDummy adds a dummy user to the database
func CreateDummyUser() *User {
	u := DummyUser()
	u2, err := u.Save()
	if err != nil {
		panic(err)
	}
	return u2
}
