package service

import (
	"todoserver/model"
	"todoserver/utils"

	"github.com/globalsign/mgo/bson"
)

type PlansService struct {
}

var mongoUtil = new(utils.MongoUtil)

//添加
func (plansService *PlansService) AddPlan(plan model.Plan, userKey string) error {
	session, c := mongoUtil.ConnectMongo("plans")
	conn := redisUtil.ConnectRedis()
	defer session.Close()
	defer conn.Close()

	userId := redisUtil.GetUserIdFromRedis(userKey)
	plan.Id = bson.NewObjectId()
	plan.UserId = userId
	err := c.Insert(plan)
	return err
}

//修改状态
func (plansService *PlansService) UpdatePlan(id string, plan model.Plan) error {
	session, c := mongoUtil.ConnectMongo("plans")
	defer session.Close()

	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"content": plan.Content, "isCompleted": plan.IsCompleted}})
	return err
}

//删除一条
func (plansService *PlansService) DeleteOne(id string) error {
	session, c := mongoUtil.ConnectMongo("plans")
	defer session.Close()
	//根据 ObjectId 删除
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//查询所有
func (plansService *PlansService) FindAll(userKey string) []model.Plan {
	session, c := mongoUtil.ConnectMongo("plans")
	defer session.Close()

	var plans []model.Plan
	userId := redisUtil.GetUserIdFromRedis(userKey)
	c.Find(bson.M{"userId": userId}).All(&plans)
	return plans
}

//修改 plan.isCompleted 为 true 或 false
func (plansService *PlansService) UpdateAllStatus(activeCount int, userKey string) error {
	session, c := mongoUtil.ConnectMongo("plans")
	defer session.Close()

	var plans []model.Plan
	var err error
	userId := redisUtil.GetUserIdFromRedis(userKey)
	c.Find(bson.M{"userId": userId}).All(&plans)
	var isCompleted bool
	//activeCount 为 ０ 时表示已全部完成，这时修改为未完成
	if activeCount == 0 {
		isCompleted = false
	} else {
		isCompleted = true
	}

	for i := 0; i < len(plans); i++ {
		err = c.Update(bson.M{"_id": plans[i].Id}, bson.M{"$set": bson.M{"isCompleted": isCompleted}})
		if err != nil {
			return err
		}
	}
	return nil
}

//删除已完成的 plan
func (plansService *PlansService) DeleteCompletedPlans(userKey string) error {
	session, c := mongoUtil.ConnectMongo("plans")
	defer session.Close()

	userId := redisUtil.GetUserIdFromRedis(userKey)
	_, err := c.RemoveAll(bson.M{"userId": userId, "isCompleted": true})
	return err
}
