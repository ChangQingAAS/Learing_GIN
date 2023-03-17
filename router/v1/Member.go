package v1

import (
	"github.com/ChangQingAAS/Learing_GIN/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context) {
	res := entity.Result{}
	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CodeError)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}
	res.SetCode(entity.CodeSuccess)
	res.SetData(data)
	c.JSON(http.StatusOK, res)
}
