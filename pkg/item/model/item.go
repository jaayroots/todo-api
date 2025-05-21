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
)
