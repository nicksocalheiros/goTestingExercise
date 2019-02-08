package utils_test

import (
	"github.com/uasouz/goTestingExercise/utils"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestRemoveNonNumeric(t *testing.T){
	document := "014.400.032-62"
	expected := "01440003262"
	if value, err := utils.RemoveNonNumeric(document); true {
		if err!= nil {
			t.Error(err)
		} else if value != expected {
			t.Errorf("Incorrect Value,found %s, expected %s",value,expected)
		}
	}
}

func TestHashToBcrypt(t *testing.T){
	expected := "$2y$12$z4NcAPlZuwqNFd7Q61Q1Q.jnGMLdgkhKSI9RH5WQCmzVr/aMtuROe"
	originalValue := "12345678"
	result,err := utils.HashToBCrypt(originalValue)
	if err!= nil {
		t.Error(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(expected),[]byte(originalValue))
	if err!= nil {
		t.Error("Error on implementation checking - Value<->Expected",err)
	}
	err = bcrypt.CompareHashAndPassword(result,[]byte(originalValue))
	if err!= nil {
		t.Error("Error on implementation checking - Value<->Result",err)
	}
}