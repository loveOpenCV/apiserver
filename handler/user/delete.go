package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
	}

	SendResponse(c, nil, nil)
}
