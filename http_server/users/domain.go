package users

import (
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/utils"
)

type SignUpRequest struct {
	Username string
	Password string
}

type ValidatedUser struct {
	Username string `json:"username" validate:"required,min=4,max=10"`
	Password string `json:"password" validate:"required,min=8"`
}

func ValidateSignUp(c *echo.Context) (ValidatedUser, utils.Error) {
	new_user, validated_user := new(SignUpRequest), ValidatedUser{}
	var err error
	if err = (*c).Bind(new_user); err != nil {
		return validated_user, utils.NewError(err)
	}

	validated_user.Username = new_user.Username
	validated_user.Password = new_user.Password

	if err = (*c).Validate(validated_user); err != nil {
		// return validated_user, fmt.Errorf("validation error : &w", utils.NewValidatorError(err))
		return validated_user, utils.NewValidatorError(err)
	}
	return validated_user, utils.Error{}
}
