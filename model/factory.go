package model

// Factory struct provides quick methods to create test content
// to be used in model tests.
type Factory struct{}

// DummyUsers or factory.DummyUsers creates a slice of valid Users
func (f *Factory) DummyUsers() []User {
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
func (f *Factory) DummyUser() User {
	return User{
		Name:     "unique_user",
		Email:    "unique@example.com",
		Password: "unique",
	}
}
