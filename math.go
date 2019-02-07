package main

import "errors"

func Sum(numbers ...interface{}) (error, float64) {
	sum := 0.0
	for _, number := range numbers {
		number64 := 0.0
		if value, ok := number.(int); ok {
			number64 = float64(value)
			sum += number64
		}else if number64, ok := number.(float64); ok {
			sum += number64
		} else {
			return errors.New("invalid sum: this function only acccept integers and float"), 0
		}
	}
	return nil, sum
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