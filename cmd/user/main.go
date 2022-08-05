package main

import (
	"fmt"
	"go-campaign-no-02/internal/service"
)

func main(){
	u := service.CreateUser()
	fmt.Println(u)
	service.UpdateInfoUser(&u,"Hung")
	fmt.Println(u)

}