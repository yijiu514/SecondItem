package models

import (
	"errors"
	"fmt"
)

var (
	// ErrUserNotExist 用户不存在
	ErrUserNotExist = errors.New("the user is not exist")
	// ErrUserISExist 用已户存在
	ErrUserISExist = errors.New("the user is exist")
	// ErrUserPermission 用户权限不足
	ErrUserPermission = errors.New("the user inadequate permissions")
)

// QueryMaxID 找到数据库当前最大id
func QueryMaxID() (int, error) {
	var max int
	err := DB.QueryRow("select max(id) from user").Scan(&max)

	if err != nil {
		return -1, fmt.Errorf("Query wrong :%w", err)
	}
	return max, nil
}

// UserQueryByEmail 根据email获取信息
func UserQueryByEmail(email string) (u User, err error) {

	DB.Get(&u, "SELECT * FROM user WHERE email = ? ", email)
	if err != nil {
		return u, fmt.Errorf("Query wrong %w", err)
	}
	return u, nil
}

// UserQueryByID 根据id获取信息
func UserQueryByID(id int) (u User, err error) {
	err = DB.Get(&u, "SELECT * FROM user WHERE id = ? ", id)
	if err != nil {
		return u, fmt.Errorf("Query wrong %w", err)
	}
	return u, nil
}

// IsUserExist 查询用户是否存在
func IsUserExist(email string) error {

	var num int
	err := DB.QueryRow("select count(*) from user where  email = ?", email).Scan(&num)
	if err != nil {
		return fmt.Errorf("Query wrong : %w", err)
	}

	if num == 1 {
		return ErrUserISExist
	}

	return nil
}

// IsUserNoExist 查询用户是否不存在
func IsUserNoExist(email string) error {

	var num int
	err := DB.QueryRow("select count(*) from user where  email = ?", email).Scan(&num)
	if err != nil {
		return fmt.Errorf("Query wrong : %w", err)
	}

	if num == 0 {
		return ErrUserNotExist
	}

	return nil
}

// Insert 数据插入（用户注册）
func (u User) Insert() (err error) {

	_, err = DB.NamedExec("insert into user (id,email,role,creatat,lockat,password,passwordsalt,sessionsalt) values (:id,:email,:role,:creatat,:lockat,:password,:passwordsalt,:sessionsalt)", &u)

	if err != nil {
		return fmt.Errorf("Insert wrong : %w", err)
	}

	return nil
}

// UpdateLockat 更新锁定时间
func (u *User) UpdateLockat(locktime int64) error {

	sqlStr := "update user set lockat = ? where id = ?"

	_, err := DB.Exec(sqlStr, locktime, u.ID)
	if err != nil {
		return fmt.Errorf("Update wrong : %w", err)
	}

	return nil
}

// UpdatePwdChange 修改密码
func (u User) UpdatePwdChange(pwd string, pwdsalt string) error {

	sqlStr := "update user set password = ?,passwordsalt = ?  where id = ?"

	_, err := DB.Exec(sqlStr, pwd, pwdsalt, u.ID)
	if err != nil {
		return fmt.Errorf("Update wrong : %w", err)
	}

	return nil
}

// UpdateRole 修改角色信息
func (u User) UpdateRole(role string) error {

	sqlStr := "update user set role = ? where id = ?"

	_, err := DB.Exec(sqlStr, role, u.ID)
	if err != nil {
		return fmt.Errorf("Update wrong : %w", err)
	}

	return nil
}

// UpdateSessionSalt 更新sesscion_salt
func (u User) UpdateSessionSalt(sessionsalt string) error {

	sqlStr := "update user set sessionsalt= ? where id = ?"

	_, err := DB.Exec(sqlStr, sessionsalt, u.ID)
	if err != nil {
		return fmt.Errorf("Update wrong : %w", err)
	}

	return nil
}

// QuerySessionSalt 根据id查询SessionSalt
func QuerySessionSalt(id int) (salt string, err error) {

	sqlStr := "select sessionsalt from user where id = ?"
	err = DB.QueryRow(sqlStr, id).Scan(&salt)
	if err != nil {
		return "", fmt.Errorf("Query wrong : %w", err)
	}

	return salt, nil
}

// EditorPermission 查询用户权限并判断
func (u User) EditorPermission() error {

	if u.Role == "manager" {
		return ErrUserPermission
	}

	return nil
}

// ManagerPermission 查询用户权限并判断
func (u User) ManagerPermission() error {

	if u.Role == "editor" {
		return ErrUserPermission
	}

	return nil
}
