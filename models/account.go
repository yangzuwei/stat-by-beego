package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Account struct {
	Date string `date`
	Num  int    `num`
}

var (
	Accounts []Account
	Actually int
)

/**
 * 每次动态调用 json文件发生更改则调用
**/
func Stat() {
	srcData, err := ioutil.ReadFile("./record.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(srcData, &Accounts)
	Actually = 0
	for _, v := range Accounts {
		Actually += v.Num
	}
}
