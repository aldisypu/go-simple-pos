package domain

type Product struct {
	ID          string  `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Stock       int     `gorm:"column:stock"`
	CategoryId  string  `gorm:"column:category_id"`
	CreatedAt   int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
