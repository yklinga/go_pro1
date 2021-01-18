package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_pro1/route"
	"go_pro1/utils"
)



func main() {
	utils.BaseDB()
	//defer db.Close() todo

	fmt.Println("hello world")
	r := gin.Default()

	r = route.RouterCollect(r)

	panic(r.Run(":8080"))
}



