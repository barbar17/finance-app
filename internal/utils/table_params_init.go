package utils

func TableParamsInit(allowedSort map[string]string, sort, order, search string) (string, string, string) {
	sortColumn := sort
	orderColumn := order
	searchPattern := "%" + search + "%"

	if v, ok := allowedSort[sort]; ok {
		sortColumn = v
	} else {
		sortColumn = "created_at"
	}

	if order != "asc" && order != "desc" {
		orderColumn = "desc"
	}

	return sortColumn, orderColumn, searchPattern
}
