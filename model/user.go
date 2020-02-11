package model

import (
	"encoding/json"
	"errors"
	"log"

	errno "myapi/pkg/err"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (user *User) SelectUserByName(name string) error {
	stmt, err := DB.Prepare("SELECT user_name,password FROM user WHERE user_name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	defer rows.Close()
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&user.UserName, &user.Password)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

// Validate the fields.
func (user *User) Validate() error {
	if user.UserName == "" || user.Password == "" {
		return errors.New(errno.ErrValidation.Message)
	}
	return nil
}

func (user *User) Create() (int64, error) {
	id, err := Insert("INSERT INTO  user(user_name,password) values (?,?)", user.UserName, user.Password)
	if err != nil {
		return 1, err
	}

	return id, nil
}

func (user *User) UserToJson() string {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}

func (user *User) JsonToUser(jsonBlob string) error {
	err := json.Unmarshal([]byte(jsonBlob), &user)
	if err != nil {
		return err
	}
	return nil
}
