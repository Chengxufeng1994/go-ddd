package repository

import (
	"gorm.io/gorm"
)

type PaginationCriteria struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"offset"`
	TotalPages int   `json:"total_pages"`
	TotalRows  int64 `json:"total_rows"`
}

type PaginationResult struct {
	PaginationCriteria
}

func Pagination(criteria *PaginationCriteria) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if criteria.Page < 0 {
			criteria.Page = 1
		}

		if criteria.Limit < 0 {
			criteria.Limit = 10
		}

		offset := (criteria.Page - 1) * criteria.Limit
		limit := criteria.Limit

		return db.Offset(offset).Limit(limit)
	}
}

const (
	OrderByASC  string = "asc"
	OrderByDESC string = "desc"
)

type OrderByCriteria struct {
	SortBy  string
	OrderBy string
}
