package model

import (
	"github.com/godcong/wego/core"
)

// RegTypeH5 ...
const RegTypeH5 = "0"

// RegTypeProgram ...
const RegTypeProgram = "1"

// WechatUser ...
type WechatUser struct {
	Model                `xorm:"extends" json:",inline"`
	UserID               string `xorm:"notnull unique default('') user_id"`
	Block                bool   `xorm:"notnull default(false) comment(禁止访问)"`                      //禁止访问
	AppID                string `xorm:"notnull default('') comment(appid)" json:"appid,omitempty"` //appid
	Type                 int    `xorm:"notnull default(0) comment(微信or小程序用户标识)" json:",omitempty"` //type
	*core.WechatUserInfo `xorm:"extends" json:",inline"`
}

// UserFromHook ...
func UserFromHook(info *core.WechatUserInfo, id string, typ int) *WechatUser {
	return &WechatUser{
		Block:          false,
		AppID:          id,
		Type:           typ,
		WechatUserInfo: info,
	}
}
