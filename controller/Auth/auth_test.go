package auth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("inersin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)
}

func TestValidateToken(t *testing.T) {
	var token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImluZXJzaW4iLCJleHAiOjE2NTU5MTI2MDZ9.IEW5XG_h-4Xt4MW2rBHh62sA6z1fW9VDqY5LLNN8n-k"
	err := ValidateToken(token)
	fmt.Println(err)
}
