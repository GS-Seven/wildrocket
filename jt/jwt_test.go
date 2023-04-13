package jt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestJwtClaims_GenerateToken(t *testing.T) {
	ex := time.Minute * 10
	sec := "70fc4956fb1d4f0f963d9cc2a3bf86e934234223425"

	j := &JwtClaims{
		ID:             1,
		Username:       "test-01",
		StandardClaims: jwt.StandardClaims{},
	}
	instance := NewJwtInstance(ex, sec)
	token, err := instance.GenerateToken(j)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestJwtClaims_ParseToken(t *testing.T) {
	ex := time.Minute * 10
	sec := "70fc4956fb1d4f0f963d9cc2a3bf86e934234223425"

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0LTAxIiwiZXhwIjoxNjgxMzc4MTIxfQ.zhQ4AJ9ofkYXWTlWY8fLxrmR46DRvKC7EZ8FI5IWHOg"
	instance := NewJwtInstance(ex, sec)

	info, err := instance.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("======= id: %v\n", info.ID)
	fmt.Printf("======= username: %v\n", info.Username)
}
