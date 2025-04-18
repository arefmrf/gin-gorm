package responses

import (
	"fmt"
	"strconv"
	userModels "web/internal/modules/user/models"
)

type User struct {
	ID    string
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModels.User) User {
	return User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
