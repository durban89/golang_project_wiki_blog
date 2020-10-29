package auth

/*
 * @Author: durban.zhang
 * @Date:   2019-12-30 18:05:42
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-30 18:11:34
 */

import (
	"log"

	"wiki/session"
)

// SessionManager Session管理器
var SessionManager *session.Manager

func init() {
	var err error
	SessionManager, err = session.GetManager("memory", "sessionid", 3600)
	if err != nil {
		log.Println(err)
		return
	}

	go SessionManager.SessionGC()
}
