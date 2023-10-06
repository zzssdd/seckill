// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductInfo = "productInfo"

// ProductInfo mapped from table <productInfo>
type ProductInfo struct {
	ID        int                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:å•†å“id" json:"id"` // å•†å“id
	Name      string                `gorm:"column:name;type:varchar(20);not null;comment:å•†å“åç§°" json:"name"`          // å•†å“åç§°
	Price     float32               `gorm:"column:price;type:float(8,2);not null;comment:å•†å“ä»·æ ¼" json:"price"`         // å•†å“ä»·æ ¼
	Pic       string               `gorm:"column:pic;type:varchar(30);comment:å•†å“å›¾ç‰‡" json:"pic"`                     // å•†å“å›¾ç‰‡
	Des       string               `gorm:"column:des;type:text;comment:å•†å“æè¿°" json:"des"`                            // å•†å“æè¿°
	Num       int32                `gorm:"column:num;type:int(11);comment:å•†å“æ•°é‡" json:"num"`                         // å•†å“æ•°é‡
	FreezeNum int32                `gorm:"column:freezeNum;type:int(11);comment:å†»ç»“å•†å“æ•°é‡" json:"freezeNum"`       // å†»ç»“å•†å“æ•°é‡
}

// TableName ProductInfo's table name
func (*ProductInfo) TableName() string {
	return TableNameProductInfo
}
