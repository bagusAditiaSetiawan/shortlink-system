package helper

import "shortlink-system/api/presenter"

func ToWebResponse(data interface{}) presenter.WebResponse {
	return presenter.WebResponse{
		Data: data,
	}
}
