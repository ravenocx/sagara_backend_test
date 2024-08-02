package utils

import (
	"strings"

	"github.com/ravenocx/clothes-store/domain/dto"
)

const (
	LOW_STOCK   = "low"
	EMPTY_STOCK = "empty"
)

func GetClothesConstructWhereQuery(query dto.GetClothesQuery) (string, []interface{}) {
	whereQuery := []string{}
	var args []interface{}

	if query.Color != "" {
		whereQuery = append(whereQuery, "color = ?")
		args = append(args, query.Color)
	}

	if query.Size != "" {
		whereQuery = append(whereQuery, "size = ?")
		args = append(args, query.Size)
	}

	switch query.Stock {
	case LOW_STOCK:
		whereQuery = append(whereQuery, "stock < 5")
	case EMPTY_STOCK:
		whereQuery = append(whereQuery, "stock = 0")
	}

	if len(whereQuery) > 0 {
		return " WHERE " + strings.Join(whereQuery, " AND "), args
	}

	return "", nil
}
