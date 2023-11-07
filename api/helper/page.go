package api_helper

import "duck-cook-recipe/entity"

func CreatePage(getCountItems func() int, limit, page int) entity.Pagination {
	tmpl := entity.Pagination{}
	countItems := getCountItems()
	total := (countItems / limit)

	remainder := (countItems % limit)
	if remainder == 0 {
		tmpl.TotalPage = total
	} else {
		tmpl.TotalPage = total + 1
	}

	tmpl.CurrentPage = page
	tmpl.RecordPerPage = limit

	if page <= 0 {
		tmpl.Next = page + 1
	} else if page < tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = page + 1
	} else if page == tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = 0
	}

	return tmpl
}
