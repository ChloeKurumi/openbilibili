package http

import (
	"net/http"

	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
)

func moPing(c *bm.Context) {
	var err error
	if err = pingSvc.Ping(c); err != nil {
		log.Error("app-resource service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
