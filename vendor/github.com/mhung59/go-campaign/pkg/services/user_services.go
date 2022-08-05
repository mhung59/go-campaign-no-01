package services

import (
	"fmt"
	model "github.com/mhung59/go-campaign/pkg/model"
)

func GenAndExportUser(id int, name string, age int, pos string, address string) {
	u := new(model.User)
	u.Id = id
	u.Name = name
	u.Age = age
	u.Pos = pos
	u.Address = address

	fmt.Printf("ID: %d \n", u.Id)
	fmt.Printf("NAME: %s \n", u.Name)
	fmt.Printf("AGE: %d \n", u.Age)
	fmt.Printf("POS: %s \n", u.Pos)
	fmt.Printf("ADDRESS: %s \n", u.Address)
}
