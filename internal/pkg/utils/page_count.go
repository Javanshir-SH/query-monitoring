package utils

// PageCount Calculator
func PageCount(totalCount int, perPage int) int {
	if totalCount == 0 {
		return 1
	}

	if perPage == 0 {
		return 0
	}

	pageCount := totalCount / perPage

	if totalCount%perPage > 0 {
		pageCount++
	}

	return pageCount
}
