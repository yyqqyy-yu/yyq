package protobuf

import (
	"dao/dao.go/model"
	"dao/dao.go/rpc/server/proto"
)

/*func (c *Condition) PbConditionToCondition(result *dao.Condition) (err error) {
	result.OrderNo = c.OrderNo
	result.UserID = c.UserId
	result.Symbol = c.Symbol
	result.Type = c.Type
	result.Side = c.Side
	result.Limit = c.Limit
	result.From = c.From
	result.Direct = c.Direct
	result.LastDealTime = c.LastDealTime
	return nil
}

func (i *MatchInfo) PbMatchToMatchCondition(result *dao.MatchCondition) (err error) {
	result.ID = uint(i.Id)
	result.Price = i.Price
	result.FieldAmount = i.FieldAmount
	result.FieldCashAmount = i.FieldCashAmount
	result.FieldFees = i.FieldFees
	result.Status = i.Status
	result.LastDealTime = i.FinishedTime
	return nil
}

func (o *OrderInfo) PbOrderToOrder(result *model.Order) (err error) {
	result.ID = uint(o.Id)
	result.Remark = o.Remark
	result.OrderNo = o.OrderNo
	result.UserID = uint(o.UserId)
	result.Symbol = o.Symbol
	result.Type = o.Type
	result.Side = o.Side
	result.Amount = o.Amount
	result.Price = o.Price
	result.FieldAmount = o.FieldAmount
	result.FieldCashAmount = o.FieldCashAmount
	result.FieldFees = o.FieldFees
	result.Status = o.Status
	result.Source = o.Source
	result.ClientOrderId = o.ClientOrderId
	result.CanceledTime = o.CanceledTime
	result.LastDealTime = o.LastDealTime
	return nil
}*/

/*func (s *OrderList) OrdersToPbList(orders []*model.Order) {
	for _, v := range orders {
		o := &OrderReq{}
		o.OrderToPbOrder(v)
		s.List = append(s.List, o)
	}
}*/

/*func (s *protobuf.OrderReq) OrderToPbOrder(v *model.Order) {
	//var createTime string
	//if v.CreatedAt != nil {
	//	createTime = v.CreatedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//var updateTime string
	//if v.UpdatedAt != nil {
	//	updateTime = v.UpdatedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//var deleteTime string
	//if v.DeletedAt != nil {
	//	deleteTime = v.DeletedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//s.CreatedAt = createTime
	//s.UpdatedAt = updateTime
	//s.DeletedAt = deleteTime
	s.Id = uint64(v.ID)
	s.CreatedAt=v.CreatedAt.Format("2006-01-02 15:04:05")
	s.UpdatedAt=v.UpdatedAt.Format("2006-01-02 15:04:05")
	s.DeletedAt = v.DeletedAt.Format("2006-01-02 15:04:05")
	s.OrderNo = v.OrderNo
	s.UserName=v.UserName
	s.Amount = v.Amount
	s.Status = v.Status
	s.FileUrl=v.FileUrl

}*/


func  PbOrderToOrder(s *model.Order,req *proto.OrderReq) {
	//var createTime string
	//if v.CreatedAt != nil {
	//	createTime = v.CreatedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//var updateTime string
	//if v.UpdatedAt != nil {
	//	updateTime = v.UpdatedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//var deleteTime string
	//if v.DeletedAt != nil {
	//	deleteTime = v.DeletedAt.Format("2006-01-02 15:04:05")
	//}
	//
	//s.CreatedAt = createTime
	//s.UpdatedAt = updateTime
	//s.DeletedAt = deleteTime
	s.ID=uint(req.Id)
	s.OrderNo = req.OrderNo
	s.UserName=req.UserName
	s.Amount = req.Amount
	s.Status = req.Status
	s.FileUrl=req.FileUrl

}
