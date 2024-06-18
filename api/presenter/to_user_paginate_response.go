package presenter

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/pkg/entities"
)

func ToUserPaginateResponse(users []entities.User, total int) fiber.Map {
	dataUser := []fiber.Map{}
	for _, user := range users {
		dataUser = append(dataUser, fiber.Map{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		})
	}
	return fiber.Map{
		"total": total,
		"list":  users,
	}
}
