package ingress

import (
	"fmt"
	"strings"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/rules"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type SetHdr struct {
	name     string
	rules    *rules.Rules
	response bool
}

func NewReqSetHdr(n string, r *rules.Rules) *SetHdr {
	return &SetHdr{name: n, rules: r}
}

func NewResSetHdr(n string, r *rules.Rules) *SetHdr {
	return &SetHdr{name: n, rules: r, response: true}
}

func (a *SetHdr) GetName() string {
	return a.name
}

func (a *SetHdr) Process(k store.K8s, annotations ...map[string]string) (err error) {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		return
	}
	for _, param := range strings.Split(input, "\n") {
		if param == "" {
			continue
		}
		indexSpace := strings.IndexByte(param, ' ')
		if indexSpace == -1 {
			return fmt.Errorf("incorrect value '%s' in request-set-header annotation", param)
		}
		a.rules.Add(&rules.SetHdr{
			HdrName:   param[:indexSpace],
			HdrFormat: "\"" + param[indexSpace+1:] + "\"",
			Response:  a.response,
		})
	}
	return
}
