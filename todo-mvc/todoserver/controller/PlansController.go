package controller

import (
	"todoserver/model"
	"todoserver/response"
	"todoserver/service"

	"github.com/gin-gonic/gin"
)

type PlansController struct {
}

var plansService = new(service.PlansService)

func (plansController *PlansController) Router(router *gin.RouterGroup) {
	//查询所有
	router.GET("/plans", plansController.findAll)
	//添加一条 plan
	router.POST("/plans", plansController.addPlan)
	//修改一条 plan 的状态或内容
	router.PUT("/plans/:id", plansController.updatePlan)
	//删除一条 plan
	router.DELETE("/plans/:id", plansController.deleteOne)
	//修改所有 plan 状态
	router.POST("/plans/batchUpdate", plansController.updateAllStatus)
	//删除所有已完成的
	router.POST("/plans/batchDelete", plansController.deleteCompletedPlans)
}

func (plansController *PlansController) findAll(context *gin.Context) {
	userId := context.GetString("userId")
	if userId == "" {
		response.Failed(context, "获取 userId 失败。")
		return
	}
	plans := plansService.FindAll(context, userId)
	response.Success(context, plans)
}

func (plansController *PlansController) addPlan(context *gin.Context) {
	userId := context.GetString("userId")
	var plan model.Plan
	err := context.BindJSON(&plan)
	if err != nil {
		response.Failed(context, response.BIND_PARAM_ERROR)
		return
	}
	err = plansService.AddPlan(context, plan, userId)
	if err != nil {
		response.Failed(context, response.ADD_ERROR)
		return
	}
	response.Success(context, response.ADD_SUCCESS)
}

func (plansController *PlansController) updatePlan(context *gin.Context) {
	id := context.Param("id")
	var plan model.Plan
	err := context.BindJSON(&plan)
	if err != nil {
		response.Failed(context, response.BIND_PARAM_ERROR)
		return
	}
	err = plansService.UpdatePlan(context, id, plan)
	if err != nil {
		response.Failed(context, response.UPDATE_ERROR)
		return
	}
	response.Success(context, response.UPDATE_SUCCESS)
}

func (plansController *PlansController) deleteOne(context *gin.Context) {
	id := context.Param("id")
	err := plansService.DeleteOne(context, id)
	if err != nil {
		response.Failed(context, response.DELETE_ERROR)
		return
	}
	response.Success(context, response.DELETE_SUCCESS)
}

func (plansController *PlansController) updateAllStatus(context *gin.Context) {
	userId := context.GetString("userId")
	var activeCount map[string]int
	err := context.BindJSON(&activeCount)
	if err != nil {
		response.Failed(context, response.BIND_PARAM_ERROR)
		return
	}
	err = plansService.UpdateAllStatus(context, activeCount["activeCount"], userId)
	if err != nil {
		response.Failed(context, response.UPDATE_ERROR)
	}
	response.Success(context, response.UPDATE_SUCCESS)
}

func (plansController *PlansController) deleteCompletedPlans(context *gin.Context) {
	userId := context.GetString("userId")
	err := plansService.DeleteCompletedPlans(context, userId)
	if err != nil {
		response.Failed(context, response.DELETE_ERROR)
		return
	}
	response.Success(context, response.DELETE_SUCCESS)
}
