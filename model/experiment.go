package model

type Bucket struct {
	BucketId          int    `json:"bucket_id" gorm:"column:id"`
	BucketName        string `json:"bucket_name" gorm:"column:bucket_name"`
	ExpId             int    `json:"exp_id" gorm:"column:exp_id"`
	Assignment        string `json:"assignment" gorm:"column:assignment"`
	AllocationPercent int    `json:"allocation_percent" gorm:"column:allocation_percent"`
	IsDefault         int    `json:"is_default" gorm:"column:is_default"`
}

func (bucketId *Bucket) TableName() string {
	return "strategy_bucket"
}

type Experiment struct {
	ExpId       int         `json:"exp_id" gorm:"column:id"`
	ExpName     string      `json:"exp_name" gorm:"column:exp_name"`
	ComicId     int         `json:"comic_id" gorm:"column:comic_id"`
	Salt        string      `json:"salt" gorm:"column:salt"`
	State       string      `json:"state" gorm:"column:state"`
	BusinessId  string      `json:"business_id" gorm:"column:business_id"`
	NewUserOnly bool        `json:"new_user_only" gorm:"column:new_user_only"`
	Buckets     []Bucket    `json:"buckets"`
	BucketMap   map[int]int `json:"_"`
}

func (experiment *Experiment) TableName() string {
	return "strategy_experiment"
}
