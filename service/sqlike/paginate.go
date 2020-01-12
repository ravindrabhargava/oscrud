package sqlike

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/si3nloong/sqlike/sql/expr"
	"github.com/si3nloong/sqlike/sqlike"
	"github.com/si3nloong/sqlike/sqlike/actions"
	"github.com/si3nloong/sqlike/sqlike/options"
	"github.com/si3nloong/sqlike/sqlike/primitive"
	"github.com/si3nloong/sqlike/types"
)

// Definition
var (
	OrderByDescending = "DESC"
)

// Paginator :
type Paginator struct {
	Cursor string
	Offset int
	Page   int
	Limit  int
	Order  map[string]string
	Select map[string]string
	Query  map[string]interface{}
}

// NewPaginator :
func NewPaginator() Paginator {
	return Paginator{}
}

// GetResult :
func (p *Paginator) GetResult(table *sqlike.Table, result interface{}) error {
	query := actions.Paginate().Limit(uint(p.Limit + 1))
	options := options.Paginate().SetDebug(true)
	selects := make([]interface{}, 0)

	if len(p.Select) > 0 {
		for key, value := range p.Select {
			if value != "" {
				selects = append(selects, expr.As(key, value))
			} else {
				selects = append(selects, expr.Column(key))
			}
		}
	} else {
		selects = append(selects, "*")
	}

	query = query.Select(selects...)
	query = query.Where(buildExprs(p.Query)...)
	for key, value := range p.Order {
		order := strings.ToLower(value)
		if order == OrderByDescending {
			query = query.OrderBy(expr.Desc(key))
		} else {
			query = query.OrderBy(expr.Asc(key))
		}
	}

	paginator, err := table.Paginate(query, options)
	if err != nil {
		return err
	}

	if p.Cursor != "" {
		decodeCursor, err := types.DecodeKey(p.Cursor)
		if err != nil {
			return err
		}

		if err = paginator.NextPage(decodeCursor); err != nil {
			return err
		}
	}

	slice := reflect.ValueOf(result)
	if err := paginator.All(slice); err != nil {
		return err
	}

	slice = slice.Elem()
	if v := slice.Len(); v > p.Limit {
		key := slice.Index(v - 1).Elem().FieldByName("Key")
		if key.CanInterface() {
			p.Cursor = key.Interface().(*types.Key).Encode()
			slice.Set(slice.Slice(0, v-1))
		}
	} else {
		p.Cursor = ""
	}

	return nil
}

func buildExprs(mapObject interface{}) []interface{} {
	queryExpr := make([]interface{}, 0)
	for key, value := range mapObject.(map[string]interface{}) {
		if key == "$AND" && reflect.TypeOf(value).Kind() == reflect.Map {
			queryExpr = append(queryExpr, expr.And(buildExprs(value)...))
		} else if key == "$OR" && reflect.TypeOf(value).Kind() == reflect.Map {
			queryExpr = append(queryExpr, expr.Or(buildExprs(value)...))
		} else {
			queryExpr = append(queryExpr, buildExpr(key, value)...)
		}
	}
	return queryExpr
}

func buildExpr(key string, value interface{}) []interface{} {
	query := make([]interface{}, 0)
	if yes := strings.HasSuffix(key, " LIKE"); yes {
		query = append(query,
			expr.Raw(
				fmt.Sprintf("%s LIKE '%%%s%%'", buildKey(strings.TrimSuffix(key, " LIKE")), value),
			),
		)
	} else if yes := strings.HasSuffix(key, " [OR]"); yes {
		// https://stackoverflow.com/questions/1127088/mysql-like-in
		vkey := strings.TrimSuffix(key, " [OR]")
		exprs := make([]*primitive.Raw, 0)
		for _, val := range strings.Split(value.(string), ",") {
			query = append(query,
				expr.Raw(
					fmt.Sprintf("%s LIKE '%%%v%%'", buildKey(vkey), val),
				),
			)
		}
		query = append(query, expr.Or(exprs))
	} else if yes := strings.HasSuffix(key, " [AND]"); yes {
		vkey := strings.TrimSuffix(key, " [AND]")
		exprs := make([]*primitive.Raw, 0)
		for _, val := range strings.Split(value.(string), ",") {
			query = append(query,
				expr.Raw(
					fmt.Sprintf("%s LIKE '%%%v%%'", buildKey(vkey), val),
				),
			)
		}
		query = append(query, expr.And(exprs))
	} else {
		query = append(query, expr.Equal(buildKey(key), value))
	}

	return query
}

func buildKey(key string) interface{} {
	if !strings.Contains(key, ".") {
		return key
	}
	keys := strings.Split(key, ".")
	return expr.JSONColumn(keys[0], keys[1:]...)
}
