package model

type LocalizedString map[string]string

type LocalizedContent struct {
	Title       string
	Description string
}

type (
	Item struct {
		ID int `json:"id"`
	}

	ItemReq struct {
		Title       LocalizedString `json:"title" validate:"required"`
		Description LocalizedString `json:"description,omitempty"`
	}

	ItemRes struct {
		ID          int              `json:"id"`
		Title       LocalizedString  `json:"title" validate:"required"`
		Description *LocalizedString `json:"description,omitempty"`
		CreatedAt   int64            `json:"created_at"`
		UpdatedAt   int64            `json:"updated_at"`
		CreatedBy   *string          `json:"created_by"`
		UpdatedBy   *string          `json:"updated_by"`
	}

	ItemWithLangRes struct {
		ID          int     `json:"id"`
		Title       *string `json:"title"`
		Description *string `json:"description"`
		CreatedAt   int64   `json:"created_at"`
		UpdatedAt   int64   `json:"updated_at"`
		CreatedBy   *string `json:"created_by"`
		UpdatedBy   *string `json:"updated_by"`
	}

	ItemSearchReq struct {
		Page   int           `json:"page" validate:"required"`
		Limit  int           `json:"limit" validate:"required"`
		Filter ItemFilterReq `json:"filter" validate:"required"`
	}

	ItemFilterReq struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *int    `json:"status"`
	}

	ItemSearchRes struct {
		Item     []*ItemRes     `json:"item"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
		Total     int64 `json:"total"`
	}
)
