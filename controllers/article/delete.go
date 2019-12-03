package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 10:56:03
 */

import "net/http"

// Delete 删除操作
func Delete(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}
