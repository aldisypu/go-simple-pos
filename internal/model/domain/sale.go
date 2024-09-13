package domain

type Sale struct {
	ID        string        `gorm:"column:id;primaryKey"`
	SaleDate  int64         `gorm:"column:sale_date;autoCreateTime:milli;autoUpdateTime:milli"`
	Total     float64       `gorm:"column:total"`
	CreatedAt int64         `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64         `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Details   []SalesDetail `gorm:"foreignKey:sale_id;references:id"`
}
