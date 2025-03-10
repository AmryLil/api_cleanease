package helpers

import "time"

type ResponseGetAllSuccess struct {
	Status     bool        `json:"status" example:"true"`
	Message    string      `json:"message" example:"success message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type ResponseGetDetailSuccess struct {
	Data    interface{} `json:"data"`
	Status  bool        `json:"status" example:"true"`
	Message string      `json:"message" example:"success message"`
}

type ResponseCUDSuccess struct {
	Message string `json:"message" example:"Success"`
	Status  bool   `json:"status" example:"true"`
}

type ResponseError struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"error message"`
}

type Pagination struct {
	TotalData    int `json:"total_data" example:"50"`
	CurrentPage  int `json:"current_page" example:"1"`
	NextPage     int `json:"next_page" example:"2"`
	PreviousPage int `json:"previous_page" example:"0"`
	PageSize     int `json:"page_size" example:"5"`
	TotalPage    int `json:"total_page" example:"10"`
}

type ResponseAuth struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"error message"`
	Data    any    `json:"data" `
}

func Response(message string, datas ...map[string]any) map[string]any {

	var res = map[string]any{
		"message": message,
	}

	if len(datas) > 0 {
		for _, data := range datas {
			for key, value := range data {
				res[key] = value
			}
		}
	}

	return res
}

func BuildErrorResponse(message string, datas ...map[string]any) map[string]any {
	var res = map[string]any{
		"status":  false,
		"message": message,
	}

	if len(datas) > 0 {
		for _, data := range datas {
			for key, value := range data {
				res[key] = value
			}
		}
	}
	return res
}

func PaginationResponse(page int, pageSize int, totalData int) Pagination {
	var pagination Pagination

	if pageSize >= totalData {
		pagination.PreviousPage = 0
		pagination.NextPage = 0
	} else {
		pagination.PreviousPage = max(page-1, -1)
		pagination.NextPage = min(page+1, (totalData+pageSize-1)/pageSize)
	}

	pagination.TotalData = totalData
	pagination.CurrentPage = page
	pagination.TotalPage = (totalData + pageSize - 1) / pageSize
	pagination.PageSize = pageSize

	return pagination
}

// new
type ResponseStatusOK struct {
	Status     bool        `json:"status" example:"true"`
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"Operation successful"`
	Timestamp  string      `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string      `json:"path" example:"/api/v1/resource"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
	RequestID  interface{} `json:"request_id"`
}

type ResponseStatusCreated struct {
	Status     bool        `json:"status" example:"true"`
	StatusCode int         `json:"status_code" example:"201"`
	Message    string      `json:"message" example:"Operation successful"`
	Timestamp  string      `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string      `json:"path" example:"/api/v1/resource"`
	Data       interface{} `json:"data"`
	RequestID  interface{} `json:"request_id"`
}

type ResponseStatusBadRequest struct {
	Status     bool   `json:"status" example:"false"`
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"Bad Request"`
	Timestamp  string `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string `json:"path" example:"/api/v1/resource"`
	Error      struct {
		Code       string      `json:"code" example:"BAD_REQUEST"`
		Details    interface{} `json:"details"`
		Suggestion string      `json:"suggestion" example:"Try checking the API documentation or verifying your input."`
	} `json:"error"`
	RequestID interface{} `json:"request_id"`
}

type ResponseStatusUnauthorized struct {
	Status     bool   `json:"status" example:"false"`
	StatusCode int    `json:"status_code" example:"401"`
	Message    string `json:"message" example:"Unauthorized"`
	Timestamp  string `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string `json:"path" example:"/api/v1/resource"`
	Error      struct {
		Code       string      `json:"code" example:"UNAUTHORIZED"`
		Details    interface{} `json:"details"`
		Suggestion string      `json:"suggestion" example:"Try checking the API documentation or verifying your input."`
	} `json:"error"`
	RequestID interface{} `json:"request_id"`
}

type ResponseStatusNotFound struct {
	Status     bool   `json:"status" example:"false"`
	StatusCode int    `json:"status_code" example:"404"`
	Message    string `json:"message" example:"Not Found"`
	Timestamp  string `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string `json:"path" example:"/api/v1/resource"`
	Error      struct {
		Code       string      `json:"code" example:"NOT_FOUND"`
		Details    interface{} `json:"details"`
		Suggestion string      `json:"suggestion" example:"Try checking the API documentation or verifying your input."`
	} `json:"error"`
	RequestID interface{} `json:"request_id"`
}

type ResponseStatusInternalServerError struct {
	Status     bool   `json:"status" example:"false"`
	StatusCode int    `json:"status_code" example:"500"`
	Message    string `json:"message" example:"Internal Server Error"`
	Timestamp  string `json:"timestamp" example:"2024-12-10T14:00:00Z"`
	Path       string `json:"path" example:"/api/v1/resource"`
	Error      struct {
		Code       string      `json:"code" example:"INTERNAL_SERVER_ERROR"`
		Details    interface{} `json:"details"`
		Suggestion string      `json:"suggestion" example:"Try again later or contact the administrator."`
	} `json:"error"`
	RequestID interface{} `json:"request_id"`
}

type ErrorParam struct {
	Message    string      `json:"message"`
	Details    interface{} `json:"details"`
	Path       string      `json:"path"`
	Suggestion string      `json:"suggestion"`
	RequestID  interface{} `json:"request_id"`
}

type SuccessParam struct {
	Message    string      `json:"message"`
	Path       string      `json:"path"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
	RequestID  interface{} `json:"request_id"`
}

type NotifData struct {
	Nama       string
	NomorTiket string
	Tanggal    string
	Email      string
	Phone      string
}

type PayloadEmail struct {
	TypeNotif string
	Subject   string
	Header    string
	Title     string
	Message   string
}

func BuildResponseStatusOK(p SuccessParam) ResponseStatusOK {
	return ResponseStatusOK{
		Status:     true,
		StatusCode: 200,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		Data:       p.Data,
		Pagination: p.Pagination,
		RequestID:  p.RequestID,
	}
}

func BuildResponseStatusCreated(p SuccessParam) ResponseStatusCreated {
	return ResponseStatusCreated{
		Status:     true,
		StatusCode: 201,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		Data:       p.Data,
		RequestID:  p.RequestID,
	}
}

func BuildResponseStatusBadRequest(p ErrorParam) ResponseStatusBadRequest {
	response := ResponseStatusBadRequest{
		Status:     false,
		StatusCode: 400,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		RequestID:  p.RequestID,
	}
	response.Error.Code = "BAD_REQUEST"
	response.Error.Details = p.Details
	response.Error.Suggestion = p.Suggestion
	return response
}

func BuildResponseStatusUnauthorized(p ErrorParam) ResponseStatusUnauthorized {
	response := ResponseStatusUnauthorized{
		Status:     false,
		StatusCode: 401,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		RequestID:  p.RequestID,
	}
	response.Error.Code = "UNAUTHORIZED"
	response.Error.Details = p.Details
	response.Error.Suggestion = p.Suggestion
	return response
}

func BuildResponseStatusNotFound(p ErrorParam) ResponseStatusNotFound {
	response := ResponseStatusNotFound{
		Status:     false,
		StatusCode: 404,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		RequestID:  p.RequestID,
	}
	response.Error.Code = "NOT_FOUND"
	response.Error.Details = p.Details
	response.Error.Suggestion = p.Suggestion
	return response
}

func BuildResponseStatusInternalServerError(p ErrorParam) ResponseStatusInternalServerError {
	response := ResponseStatusInternalServerError{
		Status:     false,
		StatusCode: 500,
		Message:    p.Message,
		Timestamp:  time.Now().String(),
		Path:       p.Path,
		RequestID:  p.RequestID,
	}
	response.Error.Code = "INTERNAL_SERVER_ERROR"
	response.Error.Details = p.Details
	response.Error.Suggestion = p.Suggestion
	return response
}

