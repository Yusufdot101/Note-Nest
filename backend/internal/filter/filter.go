package filter

import (
	"math"
	"slices"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Filter struct {
	Page         int
	PageSize     int
	Order, Sort  string
	SafeSortList []string
}

func ValidateFilter(v *validator.Validator, f *Filter) {
	v.CheckAddError(f.Page > 0, "page", "must be positive")
	v.CheckAddError(f.PageSize > 0, "page_size", "must be positive")
	v.CheckAddError(slices.Contains(f.SafeSortList, f.Sort), "sort", "invalid")
}

func (f *Filter) Limit() int {
	return f.PageSize
}

func (f *Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f *Filter) SortColumn() string {
	if slices.Contains(f.SafeSortList, f.Sort) {
		return f.Sort
	}
	panic("invalid sort column")
}

func (f *Filter) SortDirection() string {
	if f.Order == "descending" {
		return "DESC"
	}
	return "ASC"
}

type Metadata struct {
	TotalResources int
	PageSize       int
	TotalPages     int
	PageNumber     int
}

func GenerateMetadata(pageNumber, pageSize, totalResources int) *Metadata {
	totalPages := int(math.Ceil(float64(totalResources) / float64(pageSize)))
	return &Metadata{
		TotalResources: totalResources,
		PageSize:       pageSize,
		TotalPages:     totalPages,
		PageNumber:     pageNumber,
	}
}
