package middleware

import (
	"sync"

	"github.com/jonashiltl/geziyor/client"
	"github.com/jonashiltl/geziyor/internal"
)

// AllowedDomains checks for request host if it exists in AllowedDomains
type AllowedDomains struct {
	AllowedDomains []string
	logOnlyOnce    sync.Map
}

func (a *AllowedDomains) ProcessRequest(r *client.Request) {
	if len(a.AllowedDomains) != 0 && !internal.ContainsString(a.AllowedDomains, r.Host) {
		if _, logged := a.logOnlyOnce.LoadOrStore(r.Host, struct{}{}); !logged {
			internal.Logger.Printf("Domain not allowed: %s\n", r.Host)
		}
		r.Cancel()
		return
	}
}
