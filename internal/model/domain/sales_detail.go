package domain

type SalesDetail struct {
	ID        string  `gorm:"column:id;primaryKey"`
	SaleId    string  `gorm:"column:sale_id"`
	ProductId string  `gorm:"column:product_id"`
	Quantity  int     `gorm:"column:quantity"`
	Price     float64 `gorm:"column:price"`
	CreatedAt int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
