package platform

import (
	"gorm.io/gen"
	"gourd/internal/app/admin"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"net/http"
	"strconv"
	"strings"
)

type ProductCtl struct {
	admin.BaseController //继承基础控制器
}

func (c *ProductCtl) Add(w http.ResponseWriter, r *http.Request) {

	status, _ := strconv.Atoi(r.FormValue("status"))

	product := &model.Product{
		Name:   r.FormValue("name"),
		Key:    r.FormValue("key"),
		Remark: r.FormValue("remark"),
		Status: int32(status),
	}
	var qA = query.Product

	err := qA.Create(product)
	if err != nil {
		_ = c.Fail(w, http.StatusInternalServerError, "添加失败", err.Error())
		return
	}

	_ = c.Success(w, "添加成功", nil)
}

func (c *ProductCtl) Delete(w http.ResponseWriter, r *http.Request) {
	var qA = query.Product

	var idArr []int32
	for _, s := range strings.Split(r.FormValue("id"), ",") {
		id, _ := strconv.Atoi(s)
		idArr = append(idArr, int32(id))
	}

	_, err := qA.Where(qA.ID.In(idArr...)).Delete()
	if err != nil {
		_ = c.Fail(w, http.StatusInternalServerError, "删除失败", nil)
		return
	}
	_ = c.Success(w, "删除成功", nil)
}

func (c *ProductCtl) Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	status, _ := strconv.Atoi(r.FormValue("status"))

	var qA = query.Product
	_, err := qA.Where(qA.ID.Eq(int32(id))).UpdateSimple(
		qA.Name.Value(r.FormValue("name")),
		qA.Key.Value(r.FormValue("key")),
		qA.Status.Value(int32(status)),
		qA.Remark.Value(r.FormValue("remark")),
	)
	if err != nil {
		_ = c.Fail(w, http.StatusInternalServerError, "修改失败", nil)
		return
	}
	_ = c.Success(w, "修改成功", nil)
}

func (c *ProductCtl) List(w http.ResponseWriter, r *http.Request) {

	rq := r.URL.Query()

	page := 1
	if rq.Has("page") {
		page, _ = strconv.Atoi(rq.Get("page"))
	}

	pageSize := 10
	if rq.Has("page_size") {
		pageSize, _ = strconv.Atoi(rq.Get("page_size"))
	}

	var q = query.Product

	var condition []gen.Condition
	if rq.Has("name") {
		condition = append(condition, q.Name.Eq(rq.Get("name")))
	}
	if rq.Has("key") {
		condition = append(condition, q.Key.Eq(rq.Get("key")))
	}
	if rq.Get("status") != "" {
		status, _ := strconv.Atoi(rq.Get("status"))
		condition = append(condition, q.Status.Eq(int32(status)))
	}

	_q := q.Limit(pageSize).
		Offset((page - 1) * pageSize)

	if len(condition) > 0 {
		_q = _q.Where(condition...)
	}

	list, err := _q.Find()
	if err != nil {
		_ = c.Fail(w, http.StatusInternalServerError, "获取失败", nil)
		return
	}
	total, _ := q.Where(condition...).Count()
	data := map[string]any{
		"data":  list,
		"total": total,
	}
	_ = c.Success(w, "获取成功", data)
}
