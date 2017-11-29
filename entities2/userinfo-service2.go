package entities2

import (
	_ "github.com/go-sql-driver/mysql"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Insert .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	var err error
	res, err := engine.Exec(userInfoInsertStmt, "xiaolun", "haha", "1") 
	if err != nil {
		return err
	}
	return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	res, err := engine.Exec(userInfoQueryAll)
	return res
}

var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"
// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	res, err := engine.Exec(userInfoQueryByID, "1")
	return res
}
