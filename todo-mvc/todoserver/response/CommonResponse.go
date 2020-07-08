package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 1 //操作成功
	FAILED  int = 0 //操作失败

	BIND_PARAM_ERROR string = "参数绑定失败。"
	ADD_ERROR        string = "添加失败。"
	ADD_SUCCESS      string = "添加成功。"
	UPDATE_ERROR     string = "更新失败。"
	UPDATE_SUCCESS   string = "更新成功。"
	DELETE_ERROR     string = "删除失败。"
	DELETE_SUCCESS   string = "删除成功。"
	REGISTER_FAILED  string = "注册失败。"
	REGIDTER_SUCCESS string = "注册成功。"
	LOGIN_FAILED     string = "登录失败。"
	LOGIN_SUCCESS    string = "登录成功。"
)

//成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    SUCCESS,
		"message": "成功",
		"data":    v,
	})
}

//操作失败返回
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    FAILED,
		"message": v,
	})
}
