package router

import (
	"github.com/ChangQingAAS/Learing_GIN/middleware/logger"
	"github.com/ChangQingAAS/Learing_GIN/middleware/sign"
	v1 "github.com/ChangQingAAS/Learing_GIN/router/v1"
	v2 "github.com/ChangQingAAS/Learing_GIN/router/v2"
	"github.com/ChangQingAAS/Learing_GIN/validator/member"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

func InitRouter(r *gin.Engine) {
	r.Use(logger.LogToFile())

	// router/v1 表示 v1 版本的文件。v1 不需签名验证
	GroupV1 := r.Group("/v1")
	{
		// .Any 表示支持多种请求方式。
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	//router/v2 表示 v2 版本的文件。v2 需要签名验证
	GroupV2 := r.Group("/v2").Use(sign.Sign())
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", member.NameValid)
	}
}
