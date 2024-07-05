package shorted_link

import (
	"shortlink-system/api/presenter"
)

type CreateShortedLink struct {
	Url string `json:"url" form:"url" validate:"required"`
}

type PaginateShortedLink struct {
	presenter.PaginatePresenter
	Url    string `json:"url" form:"url" validate:"required"`
	UserID uint
}
