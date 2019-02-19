package model

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-spread-service/cache"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// OAuth ...
type OAuth struct {
	Scopes      []string `xorm:"oauth.scopes"`
	RedirectURL string   `xorm:"oauth.redirect_url"`
}

// Property ...
type Property struct {
	Model       `xorm:"extends" json:",inline"`
	UserID      string   `xorm:"user_id" json:"user_id"`
	Sign        string   `xorm:"sign" json:"sign"` //配置唯一识别码
	AppID       string   `xorm:"unique app_id " json:"app_id"`
	MchID       string   `xorm:"mch_id" json:"mch_id"`
	MchKey      string   `xorm:"mch_key" json:"mch_key"`
	PemCert     string   `xorm:"pem_cert" json:"pem_cert"`
	PemKEY      string   `xorm:"pem_key" json:"pem_key"`
	RootCA      string   `xorm:"root_ca" json:"root_ca"`
	NotifyURL   string   `xorm:"notify_url" json:"notify_url"`
	RefundURL   string   `xorm:"refund_url" json:"refund_url"`
	Kind        string   `xorm:"kind" json:"kind"`
	Sandbox     bool     `xorm:"sandbox" json:"sandbox" `
	AppSecret   string   `xorm:"app_secret" json:"app_secret"`
	Token       string   `xorm:"token" json:"token"`
	AesKey      string   `xorm:"aes_key" json:"aes_key"`
	PublicKey   string   `xorm:"public_key" json:"public_key"`
	PrivateKey  string   `xorm:"private_key" json:"private_key"`
	Scopes      []string `xorm:"scopes" json:"scopes"`
	RedirectURI string   `xorm:"redirect_uri" json:"redirect_uri"`
}

// NewProperty ...
func NewProperty(id string) *Property {
	return &Property{Model: Model{
		ID: id,
	}}
}

// Properties ...
func (obj *Property) Properties() ([]*Property, error) {
	var properties []*Property
	err := DB().Table(obj).Find(&properties)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return properties, nil
}

// Config ...
func (obj *Property) Config() *wego.Config {
	log.Debug(*obj)
	var config wego.Config
	config.Payment = &wego.PaymentConfig{
		AppID:     obj.AppID,
		AppSecret: obj.AppSecret,
		MchID:     obj.MchID,
		Key:       obj.MchKey,
		SafeCert: &wego.SafeCertConfig{
			Cert:   []byte(obj.PemCert),
			Key:    []byte(obj.PemKEY),
			RootCA: []byte(obj.RootCA),
		},
	}
	config.OAuth = &wego.OAuthConfig{
		Scopes:      obj.Scopes,
		RedirectURI: obj.RedirectURI,
	}
	config.OfficialAccount = &wego.OfficialAccountConfig{
		AppID:       obj.AppID,
		AppSecret:   obj.AppSecret,
		Token:       obj.Token,
		AesKey:      obj.AesKey,
		AccessToken: nil,
		OAuth:       config.OAuth,
	}
	return &config
}

// CachedConfig ...
func CachedConfig(sign string) (*wego.Config, error) {
	config := cache.GetSignConfig(sign)
	if config == nil {
		p := Property{
			Sign: sign,
		}
		b, e := model.Get(nil, &p)
		if e != nil {
			return nil, e
		}
		if !b {
			return nil, xerrors.New("no found")
		}
		config = p.Config()
		cache.SetSignConfig(sign, config)
	}
	return config, nil
}
