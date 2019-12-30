package helpers

import (
	"log"

	"github.com/durban89/wiki/session"
)

/*
* @Author: durban.zhang
* @Date:   2019-12-30 17:47:15
* @Last Modified by:   durban.zhang
* @Last Modified time: 2019-12-30 17:53:11
 */

// InitSession 初始化Session
func InitSession(sessionManager *session.Manager) {
	var err error
	sessionManager, err = session.GetManager("memory", "sessionid", 3600)
	if err != nil {
		log.Println(err)
		return
	}

	go sessionManager.SessionGC()
}
