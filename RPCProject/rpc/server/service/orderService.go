package service

import (
	"context"
	"gomod/model"
	"gomod/rpc/dao"
	"gomod/rpc/server/proto"

	"fmt"
)

type Orderservice struct {
	UserDao dao.UserDao
}

func (p *Orderservice) SelectById(ctx context.Context, req *proto.QueryId) (*proto.OrderReq, error) {

	//type Order struct {
	//	Id       int64
	//	OrderNo  string `gorm:"column:orderNo" `
	//	UserName string `gorm:"column:userName" `
	//	Amount   float32
	//	Status   string
	//	FileUrl  string `gorm:"column:fileUrl" `
	//}

	var order = &model.Order{}
	order, _ = p.UserDao.SelectId(uint(req.Id))

	//db.DB.New().
	//	Table("orders").
	//	Select("*").
	//	Where("id = ?", req.Id).
	//	Find(&order)
fmt.Println(order,"hehehhe")

	 return &proto.OrderReq{
		Id:      int64(order.ID),
		OrderNo:  order.OrderNo,
		UserName: order.UserName,
		Amount:   order.Amount,
		Status:   order.Status,
		FileUrl:  order.FileUrl,
	}, nil
}
