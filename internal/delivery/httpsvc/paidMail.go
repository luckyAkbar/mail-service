package httpsvc

import (
	"mail-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

func (s *Service) registerUser() echo.HandlerFunc {
	type request struct {
		Email                string `json:"email"`
		Username             string `json:"username"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
		PhoneNumber          string `json:"phone_number"`
	}

	return func(c echo.Context) error {
		req := request{}
		if err := c.Bind(&req); err != nil {
			return ErrInvalidPayload
		}

		userRegisterHandler := usecase.NewUserRegitrator()
		userRegisterHandler.SetUserRegistrationData(
			req.Email,
			req.Username,
			req.PhoneNumber,
			req.Password,
			req.PasswordConfirmation,
		)

		if err := userRegisterHandler.Register(); err != nil {
			return err
		}

		return nil
	}
}
