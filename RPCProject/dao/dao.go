package dao

import (

	"fmt"
	"github.com/google/wire"
	"gomod/model"
)

type Dao interface {
	Close()
	InsertOrder(or *model.Order) (err error)
	DeleteById(id uint) (err error)
	SelectAll() (orList []*model.Order, err error)
	SelectByCreateTime() (orList []*model.Order, err error)
	SelectByMoney() (orList []*model.Order, err error)
	SelectByLike(username string) (orList []*model.Order, err error)
	SelectToUpdate(or *model.Order) (err error)
	SelectById(id uint) (bor *model.Order, err error)
	UpdateByOrder(or *model.Order) (err error)
}

var Provider = wire.NewSet(New, Init)

// dao dao.
type dao struct {
	db *Orm
}

func New(db *Orm) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *Orm) (d *dao, cf func(), err error) {
	// todo maybe need use happy water get config
	d = &dao{
		db: db,
	}

	cf = d.Close
	return
}

func (d *dao) InsertOrder(or *model.Order) (err error) {

	d.db.AutoMigrate(&model.Order{})

	// Migrate the schema
	//d.db.AutoMigrate(&model.Order{})
	//d.db.SingularTable(true)
	err = d.db.Create(&or).Error
	if err != nil {
		return err
	} else {
		return nil
	}
	// 查询
	//var u = new(model.Userllllll)
	//d.db.First(u)
	//fmt.Printf("%#v\n", u)

	//var uu model.Userllllll
	//d.db.Find(&uu, "sex=?", 1)
	//fmt.Printf("%#v\n", uu)

	// 更新
	//d.db.Model(&u).Update("name","yangyuquan" )
	// 删除
}

func (d *dao) DeleteById(id uint) (err error) {
	if err = d.db.Where("id=?", id).Delete(model.Order{}).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (d *dao) SelectAll() (orList []*model.Order, err error) {
	d.db.AutoMigrate(&model.Order{})
	if err := d.db.Find(&orList).Error; err != nil {
		return orList, err
	} else {
		return orList, nil
	}
}
func (d *dao) SelectByCreateTime() (orList []*model.Order, err error) {
	d.db.AutoMigrate(&model.Order{})
	er := d.db.Model(&model.Order{}).Order("created_at ASC").Find(&orList)

	if err != nil {
		return orList, er.Error
	} else {
		return orList, nil
	}
}
func (d *dao) SelectByMoney() (orList []*model.Order, err error) {
	d.db.AutoMigrate(&model.Order{})
	er := d.db.Model(&model.Order{}).Order("amount ASC").Find(&orList)

	if er != nil {
		return orList, er.Error
	} else {
		return orList, nil
	}
}
func (d *dao) SelectByLike(username string) (orList []*model.Order, err error) {

	er := d.db.Where("user_name like ?", "%"+username+"%").Find(&orList)
	if err != nil {
		return orList, er.Error
	} else {
		return orList, nil
	}
}

func (d *dao) SelectToUpdate(or *model.Order) (err error) {
	if err = d.db.Where("id=?", or.ID).First(&or).Error; err != nil {
		return err
	}
	return nil

}

func (d *dao) SelectById(id uint) (bor *model.Order, err error) {
	or:=new(model.Order)
	if err = d.db.Where("id=?", id).First(&or).Error; err != nil {
		fmt.Println(or,"hahahhahahah")
		return or, err
	}
	fmt.Println(or,"hahahhahahah")
	return or, nil

}

func (d *dao) UpdateByOrder(or *model.Order) (err error) {
	d.db.AutoMigrate(&model.Order{})
	/*if err=  DB.Where("id=?",id).First(&todo).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}*/
	if err = d.db.Save(&or).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (d *dao) Close() {
}
