package utils

func PaginateCalculate(pageReq int, limitReq int, totalReq int) (offset int, limit int, totalPage int) {
	page := pageReq
	if page <= 0 {
		page = 1
	}
	limit = limitReq
	if limit <= 0 {
		limit = 10
	}

	offset = (page - 1) * limit

	totalPage = 0
	if totalReq > 0 {
		totalPage = int((int64(totalReq) + int64(limit) - 1) / int64(limit))
	}

	return offset, limit, totalPage
}
