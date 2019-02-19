package model

// Activity ...
type Activity struct {
	Model        `xorm:"extends" json:",inline"`
	UserID       string `xorm:"notnull unique default('') user_id" json:"user_id"`
	PropertyID   string `xorm:"notnull unique default('') property_id" json:"property_id"`
	ActivityCode string `xorm:"notnull unique default('') activity_code" json:"activity_code"` //活动码唯一ID
	Mode         string `xorm:"notnull default('')" json:"mode"`                               //活动模式
}
