package service

import (
	"context"
	"todoserver/extension"
	"todoserver/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlansService struct {
}

// 添加
func (plansService *PlansService) AddPlan(ctx context.Context, plan model.Plan, userId string) error {
	plan.Id = primitive.NewObjectID()
	plan.UserId = userId
	_, err := extension.MongoCollection("plans").InsertOne(ctx, plan)
	return err
}

// 修改状态
func (plansService *PlansService) UpdatePlan(ctx context.Context, id string, plan model.Plan) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	selector := primitive.M{
		"_id": objId,
	}
	updater := primitive.M{
		"$set": primitive.M{
			"content":     plan.Content,
			"isCompleted": plan.IsCompleted,
		},
	}
	err := extension.MongoCollection("plans").UpdateOne(ctx, selector, updater)
	return err
}

// 删除一条
func (plansService *PlansService) DeleteOne(ctx context.Context, id string) error {
	//根据 ObjectId 删除
	objId, _ := primitive.ObjectIDFromHex(id)
	err := extension.MongoCollection("plans").Remove(ctx, primitive.M{"_id": objId})
	return err
}

// 查询所有
func (plansService *PlansService) FindAll(ctx context.Context, userId string) []model.Plan {
	var plans []model.Plan
	selector := primitive.M{
		"userId": userId,
	}
	extension.MongoCollection("plans").Find(ctx, selector).All(&plans)
	return plans
}

// 修改 plan.isCompleted 为 true 或 false
func (plansService *PlansService) UpdateAllStatus(ctx context.Context, activeCount int, userId string) error {
	var plans []model.Plan
	var err error
	extension.MongoCollection("plans").Find(ctx, primitive.M{"userId": userId}).All(&plans)
	var isCompleted bool
	//activeCount 为 ０ 时表示已全部完成，这时修改为未完成
	if activeCount == 0 {
		isCompleted = false
	} else {
		isCompleted = true
	}

	for i := 0; i < len(plans); i++ {
		err = extension.MongoCollection("plans").UpdateOne(
			ctx,
			primitive.M{"_id": plans[i].Id},
			primitive.M{"$set": primitive.M{"isCompleted": isCompleted}})
		if err != nil {
			return err
		}
	}
	return nil
}

// 删除已完成的 plan
func (plansService *PlansService) DeleteCompletedPlans(ctx context.Context, userId string) error {
	_, err := extension.MongoCollection("plans").RemoveAll(ctx, primitive.M{"userId": userId, "isCompleted": true})
	return err
}
