package policies

import (
	"github.com/TeaWeb/code/teaweb/actions/default/proxy/proxyutils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
)

type Helper struct {
}

// 相关Helper
func (this *Helper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method == http.MethodGet {
		proxyutils.AddServerMenu(action, false)

		action.Data["selectedMenu"] = "list"
	}
}
