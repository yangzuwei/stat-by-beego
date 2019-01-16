package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

type Student struct {
	Id        int
	Std_name  string
	Is_handle bool
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterDataBase("std", "mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8")
	orm.RegisterModel(new(User), new(Student))
}

func TotalNum() int64 {

	o := orm.NewOrm()
	o.Using("std")

	std := new(Student)                 // 默认使用 std，你可以指定为其他数据库
	cnt, _ := o.QueryTable(std).Count() // SELECT COUNT(*) FROM STUDENT
	//fmt.Printf("std Count Num: %s, %s", cnt, err)
	return cnt
}

func HasHanle() int64 {
	o := orm.NewOrm()
	o.Using("std")

	std := new(Student)
	cnt, _ := o.QueryTable(std).Filter("is_handle", 1).Count() // SELECT COUNT(*) FROM STUDENT WHERE IS_HANDLE = 1
	//fmt.Printf("std Count Num: %s, %s", cnt, err)
	return cnt
}
