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
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      int    `json:"status"`
		DueDate     int64  `json:"due_date"`
		CreatedAt   int64  `json:"created_at"`
		UpdatedAt   int64  `json:"updated_at"`
	}
)
