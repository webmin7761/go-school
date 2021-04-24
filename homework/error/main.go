package main

import (
	"github.com/prometheus/common/log"
)

func main() {

	//dao中的操作是按具体对象来操作的，sql.ErrNoRows是通用说法，并不是表意的，例如：
	//如果是按用户名查找用户的场景下，查找不到用户应该把sql.ErrNoRows Wrap到“用户不存的”错误里
	//这样在上层如果要进行某种通用型处理也可以兼顾的到
	userDao := NewUserDao(db)

	defer db.Close()

	_, err := userDao.FindByName("foo")
	if err != nil {
		log.Errorf("%v", err)
	}

	_, err = userDao.FindByName("fo1o")
	if err != nil {
		log.Errorf("%v", err)
	}

	var users []User
	users, err = userDao.QueryBySex("W")
	if err != nil {
		log.Errorf("%v", err)
	}

	log.Infof("%v", users)
}
