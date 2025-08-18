package postgreSQL

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint      `gorm:"primarykey;index" json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"index"`

	//Pagination and Filter
	Pagination Pagination `json:"-" form:"-" gorm:"-"`
	//Practice
	IsTest bool `json:"-" form:"tutorial" gorm:"-"`

	//ILike cond
	ILike   []string `json:"-" gorm:"-"`
	Preload []string `json:"-" gorm:"-"`

	DISTINCT string `json:"-" form:"-" gorm:"-"`

	//Connection
	conn       *gorm.DB        `json:"-" form:"-" gorm:"-"`
	sql        *sql.DB         `json:"-" form:"-" gorm:"-"`
	customConn bool            `json:"-" form:"-" gorm:"-"`
	debug      bool            `json:"-" form:"-" gorm:"-"`
	joins      [][]interface{} `json:"-" form:"-" gorm:"-"`
}
