package handlers

type StandardResponse struct {
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}

const (
	StatusOk         = 200
	StatusBadRequest = 400
	StatusNotFound   = 404
)

var StatusDescriptions = map[int]string{
	StatusOk:         "Ok",
	StatusBadRequest: "Bad Request",
	StatusNotFound:   "Not Found",
}

func BadRequestResponse() StandardResponse {
	return StandardResponse{
		StatusCode:  StatusBadRequest,
		Description: StatusDescriptions[StatusBadRequest],
	}
}

func NotFoundResponse() StandardResponse {
	return StandardResponse{
		StatusCode:  StatusNotFound,
		Description: StatusDescriptions[StatusNotFound],
	}
}
