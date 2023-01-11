package request

import "errors"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthRequest) Validate() error {
	if a.Email == "" {
		return errors.New("email tidak boleh kosong")
	} else if a.Password == "" {
		return errors.New("password tidak boleh kosong")
	} else if len(a.Password) < 6 || len(a.Password) > 18 {
		return errors.New("password minimal 6 karakter dan maksimal 18 karakter")
	}

	return nil
}
