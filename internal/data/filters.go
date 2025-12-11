package data

import (
	"omnilight/internal/validator"
)

type Filters struct {
	Page         int
	Pagesize     int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {

	v.Check(f.Page > 0, "page", "must be greater than 0")
	v.Check(f.Page <= 10000000, "pagesize", "must be less than or equal to 10000000")
	v.Check(f.Pagesize > 0, "pagesize", "must be greater than 0")
	v.Check(f.Pagesize <= 100, "pagesize", "must be less than or equal to 100")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")

}
