package model

import (
	"time"
)

// HeroPcdUoload [...]
type HeroPcdUoload struct {
	ID                       uint      `gorm:"primaryKey;column:id" json:"-"`
	FileUUID                 string    `gorm:"column:file_uuid" json:"fileUuid"`
	CreatedAt                time.Time `gorm:"column:created_at" json:"createdAt"`
	OriginalUploadedFilename string    `gorm:"column:original_uploaded_filename" json:"originalUploadedFilename"`
	SavedFilename            string    `gorm:"column:saved_filename" json:"savedFilename"`
	FileSize                 int64     `gorm:"column:file_size" json:"fileSize"`
	SaveDuration             int       `gorm:"column:save_duration" json:"saveDuration"`
	PcdFileType              string    `gorm:"column:pcd_file_type" json:"pcdFileType"` // 2d or 3d
}

// TableName get sql table name.获取数据库表名
func (m *HeroPcdUoload) TableName() string {
	return "t_hero_pcd_uoload"
}

// HeroPcdUoloadColumns get sql column name.获取数据库列名
var HeroPcdUoloadColumns = struct {
	ID                       string
	FileUUID                 string
	CreatedAt                string
	OriginalUploadedFilename string
	SavedFilename            string
	FileSize                 string
	SaveDuration             string
	PcdFileType              string
}{
	ID:                       "id",
	FileUUID:                 "file_uuid",
	CreatedAt:                "created_at",
	OriginalUploadedFilename: "original_uploaded_filename",
	SavedFilename:            "saved_filename",
	FileSize:                 "file_size",
	SaveDuration:             "save_duration",
	PcdFileType:              "pcd_file_type",
}

// ImCacheInbox [...]
type ImCacheInbox struct {
	ID            uint64 `gorm:"primaryKey;column:id" json:"-"`
	CacheMsgUUID  string `gorm:"column:cache_msg_uuid" json:"cacheMsgUuid"`
	InboxUserUUID string `gorm:"column:inbox_user_uuid" json:"inboxUserUuid"`
}

// TableName get sql table name.获取数据库表名
func (m *ImCacheInbox) TableName() string {
	return "t_im_cache_inbox"
}

// ImCacheInboxColumns get sql column name.获取数据库列名
var ImCacheInboxColumns = struct {
	ID            string
	CacheMsgUUID  string
	InboxUserUUID string
}{
	ID:            "id",
	CacheMsgUUID:  "cache_msg_uuid",
	InboxUserUUID: "inbox_user_uuid",
}

// ImGroupInfo [...]
type ImGroupInfo struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	GroupUUID string `gorm:"column:group_uuid" json:"groupUuid"`
	GroupName string `gorm:"column:group_name" json:"groupName"`
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`
	DeletedAt string `gorm:"column:deleted_at" json:"deletedAt"`
	Member    string `gorm:"column:member" json:"member"`
}

// TableName get sql table name.获取数据库表名
func (m *ImGroupInfo) TableName() string {
	return "t_im_group_info"
}

// ImGroupInfoColumns get sql column name.获取数据库列名
var ImGroupInfoColumns = struct {
	ID        string
	GroupUUID string
	GroupName string
	CreatedAt string
	DeletedAt string
	Member    string
}{
	ID:        "id",
	GroupUUID: "group_uuid",
	GroupName: "group_name",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
	Member:    "member",
}

// ImGroupMember [...]
type ImGroupMember struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"-"`
	GroupUUID  string `gorm:"column:group_uuid" json:"groupUuid"`
	MemberUUID string `gorm:"column:member_uuid" json:"memberUuid"`
}

// TableName get sql table name.获取数据库表名
func (m *ImGroupMember) TableName() string {
	return "t_im_group_member"
}

// ImGroupMemberColumns get sql column name.获取数据库列名
var ImGroupMemberColumns = struct {
	ID         string
	GroupUUID  string
	MemberUUID string
}{
	ID:         "id",
	GroupUUID:  "group_uuid",
	MemberUUID: "member_uuid",
}

// ImMessage [...]
type ImMessage struct {
	ID           uint64    `gorm:"primaryKey;column:id" json:"-"`
	UUID         string    `gorm:"column:uuid" json:"uuid"`
	MsgType      string    `gorm:"column:msg_type" json:"msgType"`
	TextTypeText string    `gorm:"column:text_type_text" json:"textTypeText"`
	FileTypeURI  string    `gorm:"column:file_type_uri" json:"fileTypeUri"`
	FromUserUUID string    `gorm:"column:from_user_uuid" json:"fromUserUuid"`
	ToEntityUUID string    `gorm:"column:to_entity_uuid" json:"toEntityUuid"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ImMessage) TableName() string {
	return "t_im_message"
}

// ImMessageColumns get sql column name.获取数据库列名
var ImMessageColumns = struct {
	ID           string
	UUID         string
	MsgType      string
	TextTypeText string
	FileTypeURI  string
	FromUserUUID string
	ToEntityUUID string
	CreatedAt    string
}{
	ID:           "id",
	UUID:         "uuid",
	MsgType:      "msg_type",
	TextTypeText: "text_type_text",
	FileTypeURI:  "file_type_uri",
	FromUserUUID: "from_user_uuid",
	ToEntityUUID: "to_entity_uuid",
	CreatedAt:    "created_at",
}

// Member [...]
type Member struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	LoginID   string    `gorm:"column:login_id" json:"loginId"`
	Password  string    `gorm:"column:password" json:"password"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Member) TableName() string {
	return "t_member"
}

// MemberColumns get sql column name.获取数据库列名
var MemberColumns = struct {
	ID        string
	UUID      string
	Nickname  string
	CreatedAt string
	LoginID   string
	Password  string
	DeletedAt string
}{
	ID:        "id",
	UUID:      "uuid",
	Nickname:  "nickname",
	CreatedAt: "created_at",
	LoginID:   "login_id",
	Password:  "password",
	DeletedAt: "deleted_at",
}

// SensorStat [...]
type SensorStat struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	Time      time.Time `gorm:"column:time" json:"time"`
	DeletedAt string    `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *SensorStat) TableName() string {
	return "t_sensor_stat"
}

// SensorStatColumns get sql column name.获取数据库列名
var SensorStatColumns = struct {
	ID        string
	Time      string
	DeletedAt string
}{
	ID:        "id",
	Time:      "time",
	DeletedAt: "deleted_at",
}

// VibrationInfo [...]
type VibrationInfo struct {
	ID              uint64    `gorm:"primaryKey;column:id" json:"-"`
	DetectedTime    time.Time `gorm:"column:detected_time" json:"detectedTime"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`
	Value           float64   `gorm:"column:value" json:"value"`
	LocationInMeter float64   `gorm:"column:location_in_meter" json:"locationInMeter"`
}

// TableName get sql table name.获取数据库表名
func (m *VibrationInfo) TableName() string {
	return "t_vibration_info"
}

// VibrationInfoColumns get sql column name.获取数据库列名
var VibrationInfoColumns = struct {
	ID              string
	DetectedTime    string
	CreatedAt       string
	Value           string
	LocationInMeter string
}{
	ID:              "id",
	DetectedTime:    "detected_time",
	CreatedAt:       "created_at",
	Value:           "value",
	LocationInMeter: "location_in_meter",
}
