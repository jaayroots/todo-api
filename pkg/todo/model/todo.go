package model

type (
	Todo struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      int    `json:"status"`
		DueDate     string `json:"due_date"`
	}

	TodoReq struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Status      int    `json:"status" validate:"required"`
		DueDate     int64  `json:"due_date"`
	}

	TodoRes struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Status      int     `json:"status"`
		DueDate     int64   `json:"due_date"`
		CreatedAt   int64   `json:"created_at"`
		UpdatedAt   int64   `json:"updated_at"`
		CreatedBy   *string `json:"created_by"`
		UpdatedBy   *string `json:"updated_by"`
	}

	TodoSearchReq struct {
		Page   int           `json:"page" validate:"required"`
		Limit  int           `json:"limit" validate:"required"`
		Filter TodoFilterReq `json:"filter" validate:"required"`
	}

	TodoFilterReq struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *int    `json:"status"`
	}

	TodoSearchRes struct {
		Item     []*TodoRes     `json:"item"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
		Total     int64 `json:"total"`
	}
)
