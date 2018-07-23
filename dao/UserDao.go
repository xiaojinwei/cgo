package dao

import (
	"cgo/entity"
	"cgo/cgo"
	"log"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
	result,err := cgo.DB.Exec("INSERT INTO user(`username`,`password`,`create_time`) value(?,?,?)",user.Username,user.Password,user.CreateTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id,err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *UserDao) SelectUserByName(username string) []entity.User{
	rows,err := cgo.DB.Query("SELECT * FROM user WHERE username = ?",username)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID,&user.Username,&user.Password,&user.CreateTime)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users,user)
	}
	rows.Close()
	return users
}

func (p *UserDao)SelectAllUser() []entity.User {
	rows,err := cgo.DB.Query("SELECT * FROM user")
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID,&user.Username,&user.Password,&user.CreateTime)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users,user)
	}
	rows.Close()
	return users
}