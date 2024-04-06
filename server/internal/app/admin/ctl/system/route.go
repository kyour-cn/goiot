package system

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoute(r chi.Router) {

	aCtl := AppCtl{}
	r.HandleFunc("/app/add", aCtl.Add)
	r.HandleFunc("/app/edit", aCtl.Edit)
	r.HandleFunc("/app/list", aCtl.List)
	r.HandleFunc("/app/delete", aCtl.Delete)
}
