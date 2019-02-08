package math

import (
	"errors"
	"fmt"
	"github.com/cockroachdb/apd"
	"strconv"
)

func Sum(numbers ...interface{}) (error, float64) {
	apdCtx := apd.Context{MaxExponent:10,Precision:10}
	sum := apd.New(0.0,10)
	for _, number := range numbers {
		//number64 := apd.New(0.0,10)
		if value, ok := number.(int); ok {
			number64,condition,err := apd.NewFromString(strconv.FormatFloat(float64(value),'f',-1,64))
			if err!=nil {
				return err,0
			}
			fmt.Println(condition)
			//var result = apd.New(0,10)
			condition,err = apdCtx.Add(sum,sum,number64)
			if err!=nil {
				return err,0
			}
		}else if number64, ok := number.(float64); ok {
			number64b,condition,err := apd.NewFromString(strconv.FormatFloat(number64,'f',-1,64))
			if err!=nil {
				return err,0
			}
			fmt.Println(condition)
			//var result = apd.New(0,10)
			condition,err = apdCtx.Add(sum,sum,number64b)
			if err!=nil {
				return err,0
			}
		} else {
			return errors.New("invalid sum: this function only acccept integers and float"), 0
		}
	}
	sumValue,err :=sum.Float64()
	if err!=nil {
		return err,0
	}
	return nil, sumValue
}

func Sub(numbers ...interface{}) (error, float64) {
	sum := 0.0
	for _, number := range numbers {
		number64 := 0.0
		if value, ok := number.(int); ok {
			number64 = float64(value)
			sum -= number64
		}else if number64, ok := number.(float64); ok {
			sum -= number64
		} else {
			return errors.New("invalid sum: this function only acccept integers and float"), 0
		}
	}
	return nil, sum
}