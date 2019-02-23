package cache

import (
	"github.com/godcong/wego/cache"
	log "github.com/sirupsen/logrus"
)

// InitStateCache ...
func InitStateCache(c cache.Cache) {
	caches["state"] = c
}

// StateKey ...
func StateKey(s string) string {
	return "wego.spread.service.state." + s
}

// StateCache ...
func StateCache() cache.Cache {
	if v, b := caches["state"]; b {
		return v
	}
	return nil
}

// GetStateSign ...
func GetStateSign(key string) string {
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return ""
	}
	cfg := c.Get(StateKey(key))
	if c == nil {
		log.Debug("no config exists")
		return ""
	}
	if v, b := cfg.(string); b {
		return v
	}
	log.Debug("not bytes")
	return ""
}

// DeleteStateSign ...
func DeleteStateSign(sign string) {
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return
	}
	c.Delete(StateKey(sign))
}

// SetStateSign ...
func SetStateSign(key string, sign string) {
	c := PropertyCache()
	if c == nil {
		log.Debug("no property cache")
		return
	}
	c.SetWithTTL(StateKey(key), sign, 1800) //only keep 1800s
}
