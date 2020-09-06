package models

import "time"

type DiscussGroupMember struct {
	Model
	GroupId   int64 `gorm:"type:bigint;not null;comment:组id"`
	MemberId  int64 `gorm:"type:bigint;not null";comment:成员id`
	Timestamp int64 `gorm:"type:bigint;not null";comment:入群时间`
	Mute      int64 `gorm:"type:bigint;not null";comment:是否被禁言`
}

func AddDiscussGroupMember(groupId, memberId int64) error {
	discussGroupMember := DiscussGroupMember{
		GroupId:   groupId,
		MemberId:  memberId,
		Timestamp: time.Now().UnixNano() / 1e6,
		Mute:      0,
	}
	if err := db.Create(&discussGroupMember).Error; err != nil {
		return err
	}
	return nil
}
