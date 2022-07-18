package helper

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
	"user/model/web"
)

func Paginate(req web.PaginationReq) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		if req.PageSize > 0 {
			query = query.Limit(req.PageSize)
			if req.Page > 0 {
				offset := ((req.Page - 1) * req.PageSize)
				if req.Page == 1 {
					offset = 0
				}
				query = query.Offset(offset)
			}
		}
		return query
	}
}

func ToNullBool(b bool, valid bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: valid,
	}
}

func FromNullBool(nullBool sql.NullBool) bool {
	if nullBool.Valid {
		return nullBool.Bool
	} else {
		return false
	}
}

func FromNullString(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	} else {
		return ""
	}
}

func ToNullTime(time time.Time, valid bool) sql.NullTime {
	return sql.NullTime{
		Time:  time,
		Valid: valid,
	}
}
func FromNullTime(nullTime sql.NullTime) time.Time {
	if nullTime.Valid {
		return nullTime.Time
	}
	return time.Time{}
}
