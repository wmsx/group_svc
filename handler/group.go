package handler

import (
	"context"
	"github.com/wmsx/group_svc/models"
	proto "github.com/wmsx/group_svc/proto/group"
)

type GroupHandler struct{}

func (g GroupHandler) GetAllDiscussGroup(ctx context.Context,req *proto.GetAllDiscussGroupRequest, res *proto.GetAllDiscussGroupResponse) error {
	discussGroups, err := models.GetAllDiscussGroups()
	if err  != nil {
		return err
	}
	protoDiscussGroups :=  make([]*proto.DiscussGroup, 0)
	for _, discussGroup := range discussGroups {
		protoDiscussGroup := &proto.DiscussGroup{
			Id:     discussGroup.ID,
			Name:   discussGroup.Name,
			Cover: discussGroup.Cover,
			PostId: discussGroup.PostId,
		}
		protoDiscussGroups = append(protoDiscussGroups, protoDiscussGroup)
	}

	res.DiscussGroups = protoDiscussGroups
	return nil
}

func (g GroupHandler) JoinDiscussGroup(ctx context.Context, req *proto.JoinDiscussGroupRequest, res  *proto.JoinDiscussGroupResponse) error {
	err := models.AddDiscussGroupMember(req.GroupId, req.MengerId)
	if err != nil {
		return err
	}
	return nil
}

func (g GroupHandler) CreateDiscussGroup(ctx context.Context, req *proto.CreateDiscussGroupRequest, res *proto.CreateDiscussGroupResponse) error {
	err := models.CreateDiscussGroup(req.Name, req.Cover, req.PostId)
	if err != nil {
		return err
	}
	return nil
}

