package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string    `orm:"size(20)"`
	Avater   string    `orm:"size(500)"`
	Phone    string    `orm:"size(20)"`
	Password string    `orm:"size(32)"`
	AddTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

type UserForm struct {
	UserName string `valid:"Required;MinSize(4);MaxSize(20)"` // Name     不能为空并且最小长度是4 最大长度是20
	PassWord string `valid:"Required;MinSize(6);MaxSize(20)"` // PassWord 不能为空并且最小长度是4 最大长度是20
}

func (m *User) GetAllUser() []*User {
	info := User{}
	list := make([]*User, 0)
	info.Query().All(&list)
	return list
}

func (m *User) GetUserById(uid int64) (*User, error) {
	info := &User{}
	err := info.Query().Filter("Id", uid).One(info)
	return info, err
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *User) Read(fields ...string) (int64, error) {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return 0, err
	}
	return m.Id, nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
