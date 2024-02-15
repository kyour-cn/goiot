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

func newDevice(db *gorm.DB, opts ...gen.DOOption) device {
	_device := device{}

	_device.deviceDo.UseDB(db, opts...)
	_device.deviceDo.UseModel(&model.Device{})

	tableName := _device.deviceDo.TableName()
	_device.ALL = field.NewAsterisk(tableName)
	_device.ID = field.NewInt32(tableName, "id")
	_device.DeviceKey = field.NewString(tableName, "device_key")
	_device.Mac = field.NewString(tableName, "mac")
	_device.ProductID = field.NewInt32(tableName, "product_id")
	_device.UserID = field.NewInt32(tableName, "user_id")
	_device.Secret = field.NewString(tableName, "secret")
	_device.OnlineTime = field.NewInt32(tableName, "online_time")
	_device.CreateTime = field.NewUint(tableName, "create_time")
	_device.UpdateTime = field.NewUint(tableName, "update_time")
	_device.DeleteTime = field.NewField(tableName, "delete_time")

	_device.fillFieldMap()

	return _device
}

type device struct {
	deviceDo

	ALL        field.Asterisk
	ID         field.Int32
	DeviceKey  field.String // 设备标识id
	Mac        field.String // 设备mac地址
	ProductID  field.Int32  // 产品id
	UserID     field.Int32  // 所属用户
	Secret     field.String // 密钥
	OnlineTime field.Int32  // 最后上线时间
	CreateTime field.Uint   // 创建时间
	UpdateTime field.Uint   // 更新时间
	DeleteTime field.Field  // 删除时间 0表示未删除

	fieldMap map[string]field.Expr
}

func (d device) Table(newTableName string) *device {
	d.deviceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d device) As(alias string) *device {
	d.deviceDo.DO = *(d.deviceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *device) updateTableName(table string) *device {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.DeviceKey = field.NewString(table, "device_key")
	d.Mac = field.NewString(table, "mac")
	d.ProductID = field.NewInt32(table, "product_id")
	d.UserID = field.NewInt32(table, "user_id")
	d.Secret = field.NewString(table, "secret")
	d.OnlineTime = field.NewInt32(table, "online_time")
	d.CreateTime = field.NewUint(table, "create_time")
	d.UpdateTime = field.NewUint(table, "update_time")
	d.DeleteTime = field.NewField(table, "delete_time")

	d.fillFieldMap()

	return d
}

func (d *device) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *device) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["id"] = d.ID
	d.fieldMap["device_key"] = d.DeviceKey
	d.fieldMap["mac"] = d.Mac
	d.fieldMap["product_id"] = d.ProductID
	d.fieldMap["user_id"] = d.UserID
	d.fieldMap["secret"] = d.Secret
	d.fieldMap["online_time"] = d.OnlineTime
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["update_time"] = d.UpdateTime
	d.fieldMap["delete_time"] = d.DeleteTime
}

func (d device) clone(db *gorm.DB) device {
	d.deviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d device) replaceDB(db *gorm.DB) device {
	d.deviceDo.ReplaceDB(db)
	return d
}

type deviceDo struct{ gen.DO }

func (d deviceDo) Debug() *deviceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceDo) WithContext(ctx context.Context) *deviceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceDo) ReadDB() *deviceDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceDo) WriteDB() *deviceDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceDo) Session(config *gorm.Session) *deviceDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceDo) Clauses(conds ...clause.Expression) *deviceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceDo) Returning(value interface{}, columns ...string) *deviceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceDo) Not(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceDo) Or(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceDo) Select(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceDo) Where(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceDo) Order(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceDo) Distinct(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceDo) Omit(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceDo) Join(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceDo) Group(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceDo) Having(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceDo) Limit(limit int) *deviceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceDo) Offset(offset int) *deviceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *deviceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceDo) Unscoped() *deviceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceDo) Create(values ...*model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceDo) CreateInBatches(values []*model.Device, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceDo) Save(values ...*model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceDo) First() (*model.Device, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Take() (*model.Device, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Last() (*model.Device, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Find() ([]*model.Device, error) {
	result, err := d.DO.Find()
	return result.([]*model.Device), err
}

func (d deviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Device, err error) {
	buf := make([]*model.Device, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceDo) FindInBatches(result *[]*model.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceDo) Attrs(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceDo) Assign(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceDo) Joins(fields ...field.RelationField) *deviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceDo) Preload(fields ...field.RelationField) *deviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceDo) FirstOrInit() (*model.Device, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) FirstOrCreate() (*model.Device, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) FindByPage(offset int, limit int) (result []*model.Device, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceDo) Delete(models ...*model.Device) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceDo) withDO(do gen.Dao) *deviceDo {
	d.DO = *do.(*gen.DO)
	return d
}
