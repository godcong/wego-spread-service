package model_test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/google/uuid"
	"testing"
)

func init() {
	cfg := config.InitConfig("")
	model.InitDB(cfg)
	model.DB().ShowSQL(true)
}

// TestPermissionUser_Relate ...
func TestPermissionUser_Relate(t *testing.T) {
	m := model.PermissionUser{}
	m.UserID = "208da1b6-2315-11e9-b046-00155d02660e"
	m.PermissionID = "208da1b6-2315-11e9-b046-00155d02660f"
	fmt.Printf("%+v", m)
	permission, user, err := m.Relate()
	t.Log(permission)
	t.Log(user)
	t.Log(err)
}

// TestNewUserProperty ...
func TestNewUserProperty(t *testing.T) {
	up := model.NewProperty(uuid.New().String())
	up.PublicKey = "key"
	up.UserID = "123"
	up.Scopes = []string{"one", "two"}
	i, e := model.DB().InsertOne(up)
	t.Log(i, e)
}

// TestNew ...
func TestNewUserCallback(t *testing.T) {

}
