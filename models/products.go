package models

type Product struct {
	ID          int     `gorm:"primaryKey;column:id"`
	Name        string  `gorm:"column:name"`
	Price       float32 `gorm:"column:price"`
	Description string  `gorm:"column:description"`
}

func (p Product) TableName() string {
	return "products"

}
