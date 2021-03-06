// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"LGIS/model"
)

func newKcdj(db *gorm.DB) kcdj {
	_kcdj := kcdj{}

	_kcdj.kcdjDo.UseDB(db)
	_kcdj.kcdjDo.UseModel(&model.Kcdj{})

	tableName := _kcdj.kcdjDo.TableName()
	_kcdj.ALL = field.NewField(tableName, "*")
	_kcdj.KCAAA = field.NewString(tableName, "KCAAA")
	_kcdj.KCC = field.NewString(tableName, "KCC")
	_kcdj.JJDAJ = field.NewString(tableName, "JJDAJ")
	_kcdj.JJGLA = field.NewString(tableName, "JJGLA")
	_kcdj.DWAAC = field.NewString(tableName, "DWAAC")
	_kcdj.DWAAD = field.NewString(tableName, "DWAAD")
	_kcdj.KCBA = field.NewString(tableName, "KCBA")
	_kcdj.KCCA = field.NewString(tableName, "KCCA")
	_kcdj.KCCB = field.NewString(tableName, "KCCB")
	_kcdj.PKGKB = field.NewString(tableName, "PKGKB")
	_kcdj.KCAOC = field.NewString(tableName, "KCAOC")
	_kcdj.KCAOAE = field.NewString(tableName, "KCAOAE")
	_kcdj.KCAOAF = field.NewString(tableName, "KCAOAF")
	_kcdj.PKD = field.NewString(tableName, "PKD")
	_kcdj.JJDCBF = field.NewString(tableName, "JJDCBF")
	_kcdj.数据 = field.NewString(tableName, "数据")

	_kcdj.fillFieldMap()

	return _kcdj
}

type kcdj struct {
	kcdjDo kcdjDo

	ALL    field.Field
	KCAAA  field.String
	KCC    field.String
	JJDAJ  field.String
	JJGLA  field.String
	DWAAC  field.String
	DWAAD  field.String
	KCBA   field.String
	KCCA   field.String
	KCCB   field.String
	PKGKB  field.String
	KCAOC  field.String
	KCAOAE field.String
	KCAOAF field.String
	PKD    field.String
	JJDCBF field.String
	数据     field.String

	fieldMap map[string]field.Expr
}

func (k kcdj) Table(newTableName string) *kcdj {
	k.kcdjDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kcdj) As(alias string) *kcdj {
	k.kcdjDo.DO = *(k.kcdjDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kcdj) updateTableName(table string) *kcdj {
	k.ALL = field.NewField(table, "*")
	k.KCAAA = field.NewString(table, "KCAAA")
	k.KCC = field.NewString(table, "KCC")
	k.JJDAJ = field.NewString(table, "JJDAJ")
	k.JJGLA = field.NewString(table, "JJGLA")
	k.DWAAC = field.NewString(table, "DWAAC")
	k.DWAAD = field.NewString(table, "DWAAD")
	k.KCBA = field.NewString(table, "KCBA")
	k.KCCA = field.NewString(table, "KCCA")
	k.KCCB = field.NewString(table, "KCCB")
	k.PKGKB = field.NewString(table, "PKGKB")
	k.KCAOC = field.NewString(table, "KCAOC")
	k.KCAOAE = field.NewString(table, "KCAOAE")
	k.KCAOAF = field.NewString(table, "KCAOAF")
	k.PKD = field.NewString(table, "PKD")
	k.JJDCBF = field.NewString(table, "JJDCBF")
	k.数据 = field.NewString(table, "数据")

	k.fillFieldMap()

	return k
}

func (k *kcdj) WithContext(ctx context.Context) *kcdjDo { return k.kcdjDo.WithContext(ctx) }

func (k kcdj) TableName() string { return k.kcdjDo.TableName() }

func (k kcdj) Alias() string { return k.kcdjDo.Alias() }

func (k *kcdj) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kcdj) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 16)
	k.fieldMap["KCAAA"] = k.KCAAA
	k.fieldMap["KCC"] = k.KCC
	k.fieldMap["JJDAJ"] = k.JJDAJ
	k.fieldMap["JJGLA"] = k.JJGLA
	k.fieldMap["DWAAC"] = k.DWAAC
	k.fieldMap["DWAAD"] = k.DWAAD
	k.fieldMap["KCBA"] = k.KCBA
	k.fieldMap["KCCA"] = k.KCCA
	k.fieldMap["KCCB"] = k.KCCB
	k.fieldMap["PKGKB"] = k.PKGKB
	k.fieldMap["KCAOC"] = k.KCAOC
	k.fieldMap["KCAOAE"] = k.KCAOAE
	k.fieldMap["KCAOAF"] = k.KCAOAF
	k.fieldMap["PKD"] = k.PKD
	k.fieldMap["JJDCBF"] = k.JJDCBF
	k.fieldMap["数据"] = k.数据
}

func (k kcdj) clone(db *gorm.DB) kcdj {
	k.kcdjDo.ReplaceDB(db)
	return k
}

type kcdjDo struct{ gen.DO }

func (k kcdjDo) Debug() *kcdjDo {
	return k.withDO(k.DO.Debug())
}

func (k kcdjDo) WithContext(ctx context.Context) *kcdjDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kcdjDo) ReadDB(ctx context.Context) *kcdjDo {
	return k.WithContext(ctx).Clauses(dbresolver.Read)
}

func (k kcdjDo) WriteDB(ctx context.Context) *kcdjDo {
	return k.WithContext(ctx).Clauses(dbresolver.Write)
}

func (k kcdjDo) Clauses(conds ...clause.Expression) *kcdjDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kcdjDo) Returning(value interface{}, columns ...string) *kcdjDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kcdjDo) Not(conds ...gen.Condition) *kcdjDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kcdjDo) Or(conds ...gen.Condition) *kcdjDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kcdjDo) Select(conds ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kcdjDo) Where(conds ...gen.Condition) *kcdjDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kcdjDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *kcdjDo {
	return k.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (k kcdjDo) Order(conds ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kcdjDo) Distinct(cols ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kcdjDo) Omit(cols ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kcdjDo) Join(table schema.Tabler, on ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kcdjDo) LeftJoin(table schema.Tabler, on ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kcdjDo) RightJoin(table schema.Tabler, on ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kcdjDo) Group(cols ...field.Expr) *kcdjDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kcdjDo) Having(conds ...gen.Condition) *kcdjDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kcdjDo) Limit(limit int) *kcdjDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kcdjDo) Offset(offset int) *kcdjDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kcdjDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *kcdjDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kcdjDo) Unscoped() *kcdjDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kcdjDo) Create(values ...*model.Kcdj) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kcdjDo) CreateInBatches(values []*model.Kcdj, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kcdjDo) Save(values ...*model.Kcdj) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kcdjDo) First() (*model.Kcdj, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcdj), nil
	}
}

func (k kcdjDo) Take() (*model.Kcdj, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcdj), nil
	}
}

func (k kcdjDo) Last() (*model.Kcdj, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcdj), nil
	}
}

func (k kcdjDo) Find() ([]*model.Kcdj, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kcdj), err
}

func (k kcdjDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kcdj, err error) {
	buf := make([]*model.Kcdj, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kcdjDo) FindInBatches(result *[]*model.Kcdj, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kcdjDo) Attrs(attrs ...field.AssignExpr) *kcdjDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kcdjDo) Assign(attrs ...field.AssignExpr) *kcdjDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kcdjDo) Joins(fields ...field.RelationField) *kcdjDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kcdjDo) Preload(fields ...field.RelationField) *kcdjDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kcdjDo) FirstOrInit() (*model.Kcdj, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcdj), nil
	}
}

func (k kcdjDo) FirstOrCreate() (*model.Kcdj, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcdj), nil
	}
}

func (k kcdjDo) FindByPage(offset int, limit int) (result []*model.Kcdj, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kcdjDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k *kcdjDo) withDO(do gen.Dao) *kcdjDo {
	k.DO = *do.(*gen.DO)
	return k
}
