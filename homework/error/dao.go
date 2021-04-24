package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Name   string //主键
	Passwd string
	Sex    string
}

type UserDao struct {
	db *sql.DB
}

var (
	ErrUserNotFind = fmt.Errorf("dao: 用户不存在 %w", sql.ErrNoRows)
	errRow2User    = errors.New("dao: row to user error")
)

//创建UserDao
func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

//登录验证
func (u *UserDao) FindByName(name string) (user *User, err error) {

	rows, err := db.Query("select * from user where name = ?", name)

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("dao: FindByName失败 %w", err)
	}

	if !rows.Next() {
		return nil, ErrUserNotFind
	}

	user, err = u.rowToUser(rows)
	if err != nil {
		return nil, fmt.Errorf("%v %w", errRow2User, err)
	}

	return user, nil
}

//返回所有指定性别用户
func (u *UserDao) QueryBySex(sex string) (users []User, err error) {

	rows, err := db.Query("select * from user where sex = ?", sex)

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("dao: QueryBySex失败 %w", err)
	}

	users = []User{}
	if !rows.Next() {
		return users, nil
	}

	for rows.Next() {

		user, err := u.rowToUser(rows)
		if err != nil {
			return nil, fmt.Errorf("%v %w", errRow2User, err)
		}

		users = append(users, *user)
	}

	return users, nil
}

//row转换为User
func (u *UserDao) rowToUser(rows *sql.Rows) (user *User, err error) {

	user = &User{}

	if err := rows.Scan(&user.Name, &user.Passwd, &user.Sex); err != nil {
		return nil, err
	}

	return user, nil
}
