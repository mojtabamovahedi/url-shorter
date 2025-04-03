package types

type Link struct {
	Id    uint   `gorm:"column:id;primary_key" json:"id"`
	Url   string `gorm:"column:url" json:"url"`
	Short string `gorm:"column:short" json:"short"`
}
