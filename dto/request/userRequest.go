package request

import (
	"errors"

	"github.com/badoux/checkmail"
)

type UserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *UserRequest) Validate() error {
	if u.FirstName == "" {
		return errors.New("firstname tidak boleh kosong")
	} else if u.LastName == "" {
		return errors.New("lastname tidak boleh kosong")
	} else if u.Email == "" {
		return errors.New("email tidak boleh kosong")
	} else if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("format email tidak sesuai")
	} else if u.Password == "" {
		return errors.New("password tidak boleh kosong")
	} else if len(u.Password) < 6 || len(u.Password) > 18 {
		return errors.New("password minimal 6 karakter dan maksimal 18 karakter")
	}

	return nil
}
