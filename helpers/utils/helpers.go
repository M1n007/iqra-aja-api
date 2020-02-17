package utils

func GenerateMeta(page int, size int, totalData int) map[string]interface{} {
	resMeta := map[string]interface{}{
		"page":       page,
		"size":       size,
		"total_page": totalData / size,
		"total_data": totalData,
	}

	return resMeta
}
