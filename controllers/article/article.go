package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-30 17:49:57
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-30 18:05:18
 */

import (
	"log"

	"github.com/durban89/wiki/session"
)

// SessionManager 初始化session
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
