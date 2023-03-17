package v1

import (
	"fmt"
	"github.com/ChangQingAAS/Learing_GIN/common/alarm"
	"github.com/ChangQingAAS/Learing_GIN/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {
	// 获取Get参数
	name := c.Query("name")
	//price := c.DefaultQuery("price", "100")

	var res = entity.Result{}

	str, err := hello(name)
	if err != nil {
		res.SetCode(entity.CodeError)
		res.SetMessage(err.Error())
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	res.SetCode(entity.CodeSuccess)
	res.SetMessage(str)
	c.JSON(http.StatusOK, res)

}

func hello(name string) (str string, err error) {
	if name == "" {
		err = alarm.WeChat("name can't be null!")
		return
	}
	str = fmt.Sprintf("hello: %s", name)
	return
}
