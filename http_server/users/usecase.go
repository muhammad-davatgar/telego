package users

import (
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/utils"
)

func ValidateSignUp(c *echo.Context) (ValidatedSignUp, utils.Error) {
	new_user, validated_user := new(SignUpRequest), ValidatedSignUp{}
	var err error
	if err = (*c).Bind(new_user); err != nil {
		return validated_user, utils.NewError(err)
	}

	validated_user.Email = new_user.Email
	validated_user.Username = new_user.Username
	validated_user.Password = new_user.Password

	if err = (*c).Validate(validated_user); err != nil {
		// return validated_user, fmt.Errorf("validation error : &w", utils.NewValidatorError(err))
		return validated_user, utils.NewValidatorError(err)
	}
	return validated_user, utils.Error{}
}
