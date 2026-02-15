package dto
type UserInfoReq struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
	Age int `form:"age" json:"age" binding:"required,gte=0,lte=100"`
	Sex string `form:"sex" json:"sex" binding:"required"`
}