package validators

import (
	"strings"
	"todo-api-go/models"
)

func ParseUser(user *models.User) {
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}
