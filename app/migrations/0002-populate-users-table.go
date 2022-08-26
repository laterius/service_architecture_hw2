package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
)

type User0002 struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (u User0002) TableName() string {
	return "users"
}

func init() {

	users := []User0002{
		{Username: "user1", FirstName: "User", LastName: "One", Email: "1@user.com", Phone: "+1"},
		{Username: "user2", FirstName: "User", LastName: "Two", Email: "2@user.com", Phone: "+2"},
		{Username: "user3", FirstName: "User", LastName: "Three", Email: "3@user.com", Phone: "+3"},
	}

	mx := &gormigrate.Migration{
		ID:       "0002",
		Migrate:  mixture.CreateBatchM(users),
		Rollback: mixture.DeleteBatchR(users),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
