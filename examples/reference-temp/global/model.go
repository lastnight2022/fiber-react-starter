package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type ADTPL_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type MongoDB struct {
	Client  *mongo.Client
	Db      *mongo.Database
	Session mongo.Session
}
