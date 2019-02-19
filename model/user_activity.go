package model

// UserActivity ...
type UserActivity struct {
	Model      `xorm:"extends" json:",inline"`
	ActivityID string `xorm:"notnull unique default('') comment(活动ID) activity_id" json:"activity_id"`
	UserID     string `xorm:"notnull unique default('') comment(用户ID) user_id" json:"user_id"`
	//PropertyID string `xorm:"notnull unique default('') comment(配置ID) property_id" json:"property_id"`
	SpreadCode string `xorm:"notnull unique default('') comment(用户推广码) spread_code"  json:"spread_code"`
}
