package system

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoute(r chi.Router) {

	lCtl := AppCtl{}
	r.HandleFunc("/app/add", lCtl.Add)
	r.HandleFunc("/app/edit", lCtl.Edit)
	r.HandleFunc("/app/list", lCtl.List)
	r.HandleFunc("/app/delete", lCtl.Delete)
}
