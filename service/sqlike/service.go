package sqlike

import (
	"oscrud"
	"strings"
	"time"

	sql "github.com/si3nloong/sqlike/sqlike"
	"github.com/si3nloong/sqlike/types"
)

// Sqlike :
type Sqlike struct {
	client   *sql.Client
	database *sql.Database
}

// NewService :
func NewService(client *sql.Client) *Sqlike {
	return &Sqlike{client: client}
}

// Database :
func (service *Sqlike) Database(db string) *Sqlike {
	service.database = service.client.Database(db)
	return service
}

// ToService :
func (service *Sqlike) ToService(table string) Service {
	if service.database == nil {
		panic("You set database by `Database()` before transform to service.")
	}
	return Service{
		service.client,
		service.database,
		service.database.Table(table),
	}
}

// Service :
type Service struct {
	client   *sql.Client
	database *sql.Database
	table    *sql.Table
}

// StoreProfile :
type StoreProfile struct {
	Key             *types.Key `json:"id"`
	StoreRefID      string     `json:"storeId"`
	MerchantRefID   string     `json:"merchantId"`
	CategoryID      []string   `json:"categoryId"`
	CardID          []string   `json:"cardId"`
	CoverImageURL   string     `json:"coverImg"`
	LogoURL         string     `json:"logo"`
	Description     []string   `json:"desc"`
	TnC             []string   `json:"tnc"`
	Status          string     `json:"status"`
	Operation       []string   `json:"operation"`
	CreatedDateTime time.Time  `json:"createdAt"`
	UpdatedDateTime time.Time  `json:"updatedAt"`
}

// Find :
func (service Service) Find(ctx oscrud.Context) oscrud.Context {

	var i struct {
		Cursor string `query:"$cursor"`
		Offset int    `query:"$offset"`
		Page   int    `query:"$page"`
		Limit  int    `query:"$limit"`
		Order  string `query:"$order"`
		Select string `query:"$select"`
	}

	if err := ctx.Bind(&i); err != nil {
		return ctx.Stack(500, err).End()
	}

	order := make(map[string]string)
	if i.Order != "" {
		orders := strings.Split(i.Order, ",")
		lastKey := ""
		for _, key := range orders {
			if key == "desc" || key == "des" || key == "1-0" {
				order[lastKey] = OrderByDescending
				continue
			}
			order[key] = ""
		}
	}

	fields := make(map[string]string)
	if i.Select != "" {
		keys := strings.Split(i.Select, ",")
		for _, key := range keys {
			fields[key] = ""
		}
	}

	paginate := Paginator{
		Cursor: i.Cursor,
		Offset: i.Offset,
		Page:   i.Page,
		Limit:  i.Limit,
		Order:  order,
		Select: fields,
	}

	results := make([]*StoreProfile, 0)
	if err := paginate.GetResult(service.table, &results); err != nil {
		return ctx.Stack(500, err).End()
	}

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"cursor": paginate.Cursor,
			"offset": paginate.Offset,
			"limit":  paginate.Limit,
			"page":   paginate.Page,
		},
		"result": results,
	}
	return ctx.JSON(200, response).End()
}
