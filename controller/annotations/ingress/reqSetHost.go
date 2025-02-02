package ingress

import (
	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/rules"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type ReqSetHost struct {
	name  string
	rules *rules.Rules
}

func NewReqSetHost(n string, r *rules.Rules) *ReqSetHost {
	return &ReqSetHost{name: n, rules: r}
}

func (a *ReqSetHost) GetName() string {
	return a.name
}

func (a *ReqSetHost) Process(k store.K8s, annotations ...map[string]string) (err error) {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		return
	}
	a.rules.Add(&rules.SetHdr{
		HdrName:   "Host",
		HdrFormat: input,
	})
	return
}
