package shorted_link

type CreateShortedLink struct {
	Url string `json:"url" form:"url" validate:"required"`
}
