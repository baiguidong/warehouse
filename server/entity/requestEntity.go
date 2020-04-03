package entity

import "time"

// 登录
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册
type Register struct {
	OnlineUsername string `json:"online_username"` // 当前登录用户名
	Username       string `json:"username"`        // 用户名
	Password       string `json:"password"`        // 密码
	Administrator  string `json:"administrator"`   // 超级管理员 Y | N
}

// 修改密码
type UpdatePass struct {
	OldPassword string `json:"oldPassword"` // 当前登录用户名
	Username    string `json:"username"`    // 用户名
	NewPassword string `json:"newPassword"` // 密码
}

// 删除ids
type DeleteIds struct {
	Ids []int64 `json:"ids"` // ids
}

// 添加商品种类
type GoodsType struct {
	GoodsName           string    `json:"goods_name"`            // 商品名称
	GoodsSpecs          string    `json:"goods_specs"`           // 商品规格 1.盒 2.瓶 3.支
	GoodsUnitPrince     float64   `json:"goods_unitprince"`      // 商品成本价
	GoodsPrince         float64   `json:"goods_prince"`          // 商品销售价
	GoodsDiscountPrince float64   `json:"goods_discount_prince"` // 商品折扣价
	GoodsImage          string    `json:"goods_image"`           // 商品图片
	GoodsBatchNumber    string    `json:"goods_batch_number"`    // 生产批号
	GoodsDate           time.Time `json:"goods_date"`            // 生产日期
}
