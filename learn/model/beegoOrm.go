package model

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type BeeUser struct {
	Id       int
	Name     string
	Profile  *Profile  `orm:"rel(one)"`      // OneToOne relation
	Post     []*Post   `orm:"reverse(many)"` // 设置一对多的反向关系
	CreateAt time.Time `orm:"null"`
}

type Profile struct {
	Id   int
	Age  int16
	User *BeeUser `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *BeeUser `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag   `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDataBase("default", "sqlite3", "/tmp/beegodb.db")

	orm.RegisterModel(new(BeeUser), new(Profile), new(Post), new(Tag))
	orm.RunSyncdb("default", false, true)
}
