package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/uasouz/goTestingExercise/math"
)

var (

)

func main(){
err , sum:= math.Sum(0.1,0.3)
if err != nil {
	log.Error(err.Error())
	return
}
fmt.Println("Sum",sum)
	err , sub:= math.Sub(1,2.14241,2,1,23,2,2,1,1)
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("Sub",sub)
}