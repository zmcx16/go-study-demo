package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestBaseValidator_ValidSameConfirmPassword_ValidateSuccess(t *testing.T) {

	var req = RegisterAxisClut{
		Username:       "Megumi",
		PasswordNew:    "Nice Explosion!",
		PasswordRepeat: "Nice Explosion!",
		Email:          "megumi@explosion.konosuba",
	}

	validate := validator.New()
	err := validate.Struct(req)

	assert.NoError(t, err)
}

func TestBaseValidator_ValidDifferentConfirmPassword_ValidateFailed(t *testing.T) {

	var req = RegisterAxisClut{
		Username:       "Aqua",
		PasswordNew:    "Join Axis Cult Now!",
		PasswordRepeat: "Join Eris Cult Now!",
		Email:          "aqua_axiscult.konosuba",
	}

	validate := validator.New()
	err := validate.Struct(req)

	assert.Error(t, err)
}
