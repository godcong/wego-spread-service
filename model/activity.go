package model

// Activity ...
type Activity struct {
	Model        `xorm:"extends" json:",inline"`
	UserID       string `xorm:"notnull unique default('') comment(用户ID) user_id" json:"user_id"`
	PropertyID   string `xorm:"notnull unique default('') comment(配置ID) property_id" json:"property_id"`
	ActivityCode string `xorm:"notnull unique default('') comment(活动码) activity_code" json:"activity_code"` //活动码
	Verify       bool   `xorm:"notnull default(false) comment(是否校验) verify" json:"verify"`                  //是否校验
	Mode         string `xorm:"notnull default('') comment(活动模式) mode" json:"mode"`                         //活动模式
}

// NewActivity ...
func NewActivity(id string) *Activity {
	return &Activity{
		Model: Model{ID: id},
	}
}
