package cache

import (
	"encoding/json"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego/cache"
	log "github.com/sirupsen/logrus"
)

// InitPropertyCache ...
func InitPropertyCache(c cache.Cache) {
	caches["property"] = c
}

// PropertyKey ...
func PropertyKey(s string) string {
	return "wego.spread.service.property." + s
}

// PropertyCache ...
func PropertyCache() cache.Cache {
	if v, b := caches["property"]; b {
		return v
	}

	return nil
}

// GetSignConfig ...
func GetSignConfig(sign string) *wego.Config {
	var config wego.Config
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return nil
	}
	cfg := c.Get(PropertyKey(sign))
	if c == nil {
		log.Debug("no config exists")
		return nil
	}
	if v, b := cfg.([]byte); b {
		e := json.Unmarshal(v, &config)
		if e == nil {
			return &config
		}
		log.Error(e)
	}
	log.Debug("not bytes")
	return nil
}

// SetSignConfig ...
func SetSignConfig(sign string, config *wego.Config) {
	bytes, e := json.Marshal(config)
	if e != nil {
		log.Error(e)
		return
	}
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return
	}
	c.Set(PropertyKey(sign), string(bytes))
}

// DeleteSignConfig ...
func DeleteSignConfig(sign string) {
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return
	}
	c.Delete(PropertyKey(sign))
}

// CachedConfig ...
func CachedConfig(code string) (*wego.Config, error) {
	config := GetSignConfig(code)
	if config == nil {
		act := model.Activity{
			Code: code,
		}
		p, e := act.CodeProperty()
		if e != nil {
			return nil, e
		}
		config = p.Config()
		SetSignConfig(code, config)
	}
	return config, nil
}
