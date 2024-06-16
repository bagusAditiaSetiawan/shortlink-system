package presenter

type PaginatePresenter struct {
	Limit int `json:"limit" form:"limit" validate:"required"`
	Page  int `json:"page" form:"page" validate:"required"`
}

func (presenter *PaginatePresenter) GetLimit() int {
	if presenter.Limit != 0 {
		return presenter.Limit
	}
	return 10
}

func (presenter *PaginatePresenter) GetOffset() int {
	if presenter.Page <= 0 {
		return 0
	}
	return (presenter.Page - 1) * presenter.Limit
}
