package platform

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoute(r chi.Router) {

	pCtl := ProductCtl{}
	r.HandleFunc("/product/add", pCtl.Add)
	r.HandleFunc("/product/edit", pCtl.Edit)
	r.HandleFunc("/product/list", pCtl.List)
	r.HandleFunc("/product/delete", pCtl.Delete)
}
