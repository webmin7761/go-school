package main

type User struct {
	Name string
	Sex  string
}

type UserDao struct {
}

//根据sql只能得到一个用户
func (u *UserDao) Single(sql string) (user *User, err error) {
	return nil, nil
}

//根据sql只能得到一个用户或查询不出任何用户
func (u *UserDao) SingleOrDefault(sql string) (user *User, err error) {
	return nil, nil
}

//返回符合sql的所有用户
func (u *UserDao) Query(sql string) (users []User, err error) {
	return nil, nil
}
