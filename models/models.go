package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id       int
	Username string
	Password string
}

type Student struct {
	Id        int
	Std_name  string
	Is_handle bool
}

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterDataBase("std", "mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8")
	orm.RegisterModel(new(Users), new(Student))
	o = orm.NewOrm()
	o.Using("std")
}

func TotalNum() int64 {

	std := new(Student)                 // 默认使用 std，你可以指定为其他数据库
	cnt, _ := o.QueryTable(std).Count() // SELECT COUNT(*) FROM STUDENT
	//fmt.Printf("std Count Num: %s, %s", cnt, err)
	return cnt
}

func HasHanle() int64 {

	std := new(Student)
	cnt, _ := o.QueryTable(std).Filter("is_handle", 1).Count() // SELECT COUNT(*) FROM STUDENT WHERE IS_HANDLE = 1
	//fmt.Printf("std Count Num: %s, %s", cnt, err)
	return cnt
}

func CheckUser(username string, password string) bool {
	var user Users
	err := o.Raw("select * from users where username = ?", username).QueryRow(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	fmt.Println("password is ", user)
	return password == user.Password
}
