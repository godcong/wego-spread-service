package model

// UserActivity ...
type UserActivity struct {
	Model        `xorm:"extends" json:",inline"`
	ActivityID   string `xorm:"notnull unique default('') comment(活动ID) activity_id" json:"activity_id"`
	UserID       string `xorm:"notnull unique default('') comment(参加活动的用户ID) user_id" json:"user_id"`
	SpreadCode   string `xorm:"notnull unique default('') comment(参加活动的用户推广码) spread_code"  json:"spread_code"`
	SpreadNumber int64  `xorm:"notnull default('') comment(推广数) spread_code" json:"spread_number"`
}
