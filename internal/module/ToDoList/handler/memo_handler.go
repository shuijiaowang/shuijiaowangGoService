package handler

import (
	"SService/internal/module/ToDoList/dto"
	"SService/internal/module/ToDoList/service"
	"SService/internal/module/common/handler"
	"SService/pkg/util"

	"github.com/gin-gonic/gin"
)

type MemoHandler struct {
	handler.BaseHandler
}

func NewMemoHandler() *MemoHandler {
	return &MemoHandler{}
}

func (h *MemoHandler) AddMemo(c *gin.Context) {
	userUUID := h.GetUserUUID(c)
	var req dto.MemoCreateDTO
	h.Bind(c, &req)

	req.UserUUID = userUUID
	memo, err := service.AddMemo(req)
	if err != nil {
		util.Result(c, 500, "添加失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "添加成功", memo)
}
