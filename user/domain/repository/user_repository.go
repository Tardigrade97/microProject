// @author: Tardigrade
// @since: 2022/4/9
// @desc: //TODO

package repository

import "github.com/microProject/user/domain/model"

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	//根据id查找用户信息
	FindUserByID(int64) (*model.User, error)
	//创建用户
	CreateUser(*model.User) (int64, error)
	//根据ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有有用
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) {

}