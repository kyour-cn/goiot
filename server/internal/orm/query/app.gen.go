// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gourd/internal/orm/model"
)

func newApp(db *gorm.DB, opts ...gen.DOOption) app {
	_app := app{}

	_app.appDo.UseDB(db, opts...)
	_app.appDo.UseModel(&model.App{})

	tableName := _app.appDo.TableName()
	_app.ALL = field.NewAsterisk(tableName)
	_app.ID = field.NewInt32(tableName, "id")
	_app.Name = field.NewString(tableName, "name")
	_app.Key = field.NewString(tableName, "key")
	_app.Remark = field.NewString(tableName, "remark")
	_app.Status = field.NewInt32(tableName, "status")
	_app.Sort = field.NewInt32(tableName, "sort")

	_app.fillFieldMap()

	return _app
}

// app 应用列表
type app struct {
	appDo

	ALL    field.Asterisk
	ID     field.Int32
	Name   field.String // 应用名称
	Key    field.String // 应用KEY 别名
	Remark field.String // 备注
	Status field.Int32  // 状态
	Sort   field.Int32  // 排序 ASC

	fieldMap map[string]field.Expr
}

func (a app) Table(newTableName string) *app {
	a.appDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a app) As(alias string) *app {
	a.appDo.DO = *(a.appDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *app) updateTableName(table string) *app {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Name = field.NewString(table, "name")
	a.Key = field.NewString(table, "key")
	a.Remark = field.NewString(table, "remark")
	a.Status = field.NewInt32(table, "status")
	a.Sort = field.NewInt32(table, "sort")

	a.fillFieldMap()

	return a
}

func (a *app) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *app) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["key"] = a.Key
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["status"] = a.Status
	a.fieldMap["sort"] = a.Sort
}

func (a app) clone(db *gorm.DB) app {
	a.appDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a app) replaceDB(db *gorm.DB) app {
	a.appDo.ReplaceDB(db)
	return a
}

type appDo struct{ gen.DO }

func (a appDo) Debug() *appDo {
	return a.withDO(a.DO.Debug())
}

func (a appDo) WithContext(ctx context.Context) *appDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appDo) ReadDB() *appDo {
	return a.Clauses(dbresolver.Read)
}

func (a appDo) WriteDB() *appDo {
	return a.Clauses(dbresolver.Write)
}

func (a appDo) Session(config *gorm.Session) *appDo {
	return a.withDO(a.DO.Session(config))
}

func (a appDo) Clauses(conds ...clause.Expression) *appDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appDo) Returning(value interface{}, columns ...string) *appDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appDo) Not(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appDo) Or(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appDo) Select(conds ...field.Expr) *appDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appDo) Where(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appDo) Order(conds ...field.Expr) *appDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appDo) Distinct(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appDo) Omit(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appDo) Join(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appDo) LeftJoin(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appDo) RightJoin(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appDo) Group(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appDo) Having(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appDo) Limit(limit int) *appDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appDo) Offset(offset int) *appDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *appDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appDo) Unscoped() *appDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appDo) Create(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appDo) CreateInBatches(values []*model.App, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appDo) Save(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appDo) First() (*model.App, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Take() (*model.App, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Last() (*model.App, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Find() ([]*model.App, error) {
	result, err := a.DO.Find()
	return result.([]*model.App), err
}

func (a appDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.App, err error) {
	buf := make([]*model.App, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appDo) FindInBatches(result *[]*model.App, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appDo) Attrs(attrs ...field.AssignExpr) *appDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appDo) Assign(attrs ...field.AssignExpr) *appDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appDo) Joins(fields ...field.RelationField) *appDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appDo) Preload(fields ...field.RelationField) *appDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appDo) FirstOrInit() (*model.App, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FirstOrCreate() (*model.App, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FindByPage(offset int, limit int) (result []*model.App, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appDo) Delete(models ...*model.App) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appDo) withDO(do gen.Dao) *appDo {
	a.DO = *do.(*gen.DO)
	return a
}