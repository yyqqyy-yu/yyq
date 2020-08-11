package dao

import (

	"fmt"
	"github.com/jinzhu/gorm"
	"gomod/dao"
	"gomod/model"
)

type UserDao struct {
	db *gorm.DB
}

//根据ID查询对应model
func (d *UserDao) SelectId(id uint) (order *model.Order, err error) {
	d.db=dao.DB
	or:=new(model.Order)
	fmt.Println(id)
	err=d.db.DB().Ping()
	fmt.Println(err,"lianjieSQL")
	if err = d.db.Where("id=?", id).First(&or).Error; err != nil {
		fmt.Println(or,"hahhaha")
		fmt.Println(err)
		return or, err
	}
	fmt.Println(or,"hahhaha")
	return or, nil
}
