package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerValidator_ValidCandidateAxisCultBeliever_ValidateSuccess(t *testing.T) {

	var profile = &RoleProfile{
		UserName: "Kazuma",
		Religion: "none",
	}

	err := profile.validate()
	assert.NoError(t, err)
}

func TestCustomerValidator_ValidErisCultBeliever_ValidateFailed(t *testing.T) {

	var profile = &RoleProfile{
		UserName: "Darkness",
		Religion: "ErisCult",
	}

	err := profile.validate()
	assert.Error(t, err)
}
