package models

type DiscussGroupMember struct {
	Model
	GroupId  int64 `gorm:"type:bigint;not null"`
	MemberId int64 `gorm:"type:bigint;not null"`
}

func AddDiscussGroupMember(groupId, memberId int64) error {
	discussGroupMember := DiscussGroupMember{
		GroupId:  groupId,
		MemberId: memberId,
	}
	if err := db.Create(&discussGroupMember).Error; err != nil {
		return err
	}
	return nil
}

