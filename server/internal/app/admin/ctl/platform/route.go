package platform

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoute(r chi.Router) {

	// 产品
	pCtl := ProductCtl{}
	r.HandleFunc("/product/add", pCtl.Add)
	r.HandleFunc("/product/edit", pCtl.Edit)
	r.HandleFunc("/product/list", pCtl.List)
	r.HandleFunc("/product/delete", pCtl.Delete)

	// 设备
	dCtl := DeviceCtl{}
	r.HandleFunc("/device/add", dCtl.Add)
	r.HandleFunc("/device/edit", dCtl.Edit)
	r.HandleFunc("/device/list", dCtl.List)
	r.HandleFunc("/device/delete", dCtl.Delete)
}
