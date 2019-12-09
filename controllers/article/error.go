package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:25
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 10:56:27
 */

import "net/http"

// Error 错误显示
func Error(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}
