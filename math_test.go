package main

import "testing"

func TestSum(t *testing.T){
	err,total := Sum(5,7.8)
	if err !=nil{
		t.Error(err)
	}
	if total != 12.8{
		t.Errorf("Sum was incorrect, got %f, expected: %f",total,12.8)
	}
}