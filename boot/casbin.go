package boot

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

var syncedEnforcer *casbin.SyncedEnforcer

func CasBinSetup() {
	Apter, err := gormadapter.NewAdapterByDB(Db)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	syncedEnforcer, err = casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
}

func CasBin() (*casbin.SyncedEnforcer, error) {
	if err := syncedEnforcer.LoadPolicy(); err == nil {
		return syncedEnforcer, err
	} else {
		log.Println("casbin rbac_model or policy init error, message: ", err.Error())
		return nil, err
	}
}
