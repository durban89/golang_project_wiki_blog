package article

import (
	"net/http"

	"github.com/durban89/wiki/views"
)

// Tag 标签
func Tag(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "article/tag.html", 200)
}
