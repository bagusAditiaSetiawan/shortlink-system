package presenter

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/pkg/entities"
)

func ToShortedLinkResponse(shortedLinks []entities.ShortedLink, total int) fiber.Map {
	data := []fiber.Map{}
	for _, shortedLink := range shortedLinks {
		data = append(data, fiber.Map{
			"id":                shortedLink.ID,
			"link_shorted_full": shortedLink.LinkShortedFull,
			"link_original":     shortedLink.LinkOriginal,
			"num_of_accessed":   shortedLink.NumOfAccessed,
			"owner":             shortedLink.User.Username,
		})
	}
	return fiber.Map{
		"total": total,
		"list":  data,
	}
}
