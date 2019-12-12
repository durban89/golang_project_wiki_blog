package helpers

/*
 * @Author: durban.zhang
 * @Date:   2019-12-09 11:22:03
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-12 16:52:46
 */

import (
	"log"
	"net/http"
	"net/url"
)

// RedirectWithMsg 回跳
func RedirectWithMsg(r *http.Request, msg string) string {
	u, err := url.Parse(r.Referer())
	if err != nil {
		log.Fatal("url Parse error:", err)
	}

	q := u.Query()

	q.Del("err_msg")

	nu := &url.URL{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     u.Path,
		RawQuery: q.Encode() + "&err_msg=" + msg,
	}

	return nu.String()
}
