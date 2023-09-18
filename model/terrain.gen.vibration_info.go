package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _VibrationInfoMgr struct {
	*_BaseMgr
}

// VibrationInfoMgr open func
func VibrationInfoMgr(db *gorm.DB) *_VibrationInfoMgr {
	if db == nil {
		panic(fmt.Errorf("VibrationInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VibrationInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_vibration_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VibrationInfoMgr) GetTableName() string {
	return "t_vibration_info"
}

// Reset 重置gorm会话
func (obj *_VibrationInfoMgr) Reset() *_VibrationInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VibrationInfoMgr) Get() (result VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_VibrationInfoMgr) Gets() (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VibrationInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VibrationInfoMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDetectedTime detected_time获取
func (obj *_VibrationInfoMgr) WithDetectedTime(detectedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["detected_time"] = detectedTime })
}

// WithCreatedAt created_at获取
func (obj *_VibrationInfoMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithValue value获取
func (obj *_VibrationInfoMgr) WithValue(value float64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithLocationInMeter location_in_meter获取
func (obj *_VibrationInfoMgr) WithLocationInMeter(locationInMeter float64) Option {
	return optionFunc(func(o *options) { o.query["location_in_meter"] = locationInMeter })
}

// GetByOption 功能选项模式获取
func (obj *_VibrationInfoMgr) GetByOption(opts ...Option) (result VibrationInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VibrationInfoMgr) GetByOptions(opts ...Option) (results []*VibrationInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_VibrationInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]VibrationInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_VibrationInfoMgr) GetFromID(id uint64) (result VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_VibrationInfoMgr) GetBatchFromID(ids []uint64) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDetectedTime 通过detected_time获取内容
func (obj *_VibrationInfoMgr) GetFromDetectedTime(detectedTime time.Time) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`detected_time` = ?", detectedTime).Find(&results).Error

	return
}

// GetBatchFromDetectedTime 批量查找
func (obj *_VibrationInfoMgr) GetBatchFromDetectedTime(detectedTimes []time.Time) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`detected_time` IN (?)", detectedTimes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_VibrationInfoMgr) GetFromCreatedAt(createdAt time.Time) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_VibrationInfoMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_VibrationInfoMgr) GetFromValue(value float64) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_VibrationInfoMgr) GetBatchFromValue(values []float64) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromLocationInMeter 通过location_in_meter获取内容
func (obj *_VibrationInfoMgr) GetFromLocationInMeter(locationInMeter float64) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`location_in_meter` = ?", locationInMeter).Find(&results).Error

	return
}

// GetBatchFromLocationInMeter 批量查找
func (obj *_VibrationInfoMgr) GetBatchFromLocationInMeter(locationInMeters []float64) (results []*VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`location_in_meter` IN (?)", locationInMeters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_VibrationInfoMgr) FetchByPrimaryKey(id uint64) (result VibrationInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VibrationInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
