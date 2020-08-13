package util

import (
	"ab_division/log"
	"ab_division/model"
	"ab_division/util/db"
	"fmt"
)

func getBucketDetail(expId int) []model.Bucket {
	var buckets []model.Bucket
	err := db.ABExperimentDB().Model(&model.Bucket{}).Where("exp_id=? and is_default=0", expId).Find(&buckets).Error
	if err != nil {
		log.DataLogger.Errorf("Query experiment failed: %v", err)
		return nil
	}
	return buckets
}

func GetABExperiments() []model.Experiment {
	var experiments []model.Experiment
	err := db.ABExperimentDB().Model(&model.Experiment{}).Where("state='RUNNING' ").Find(&experiments).Error
	if err != nil {
		log.DataLogger.Errorf("Query experiment failed: %v", err)
		return nil
	}
	for i := range experiments {
		buckets := getBucketDetail(experiments[i].ExpId)
		if buckets != nil {
			experiments[i].Buckets = buckets
			bucketMap, err := calculateBucketMap(buckets)
			if err != nil {
				log.DataLogger.Error(err)
			} else {
				experiments[i].BucketMap = bucketMap
			}
			log.DataLogger.Infof("get experiment %v", experiments[i].ExpId)
		} else {
			return nil
		}
	}

	return experiments
}

//AllocationPercent为对应桶分流比例，生成桶时控制所有分流比例相加为100
func calculateBucketMap(buckets []model.Bucket) (map[int]int, error) {
	sum := 0
	bucketMap := make(map[int]int)
	for i := range buckets {
		sum += buckets[i].AllocationPercent
		bucketMap[buckets[i].BucketId] = sum
	}
	if sum != 100 {
		return nil, fmt.Errorf("bucket allocation exceed limit")
	}
	return bucketMap, nil
}

func LoadExperiments() {
	exps := GetABExperiments()
	if exps == nil {
		panic("Get Experiments or Bucket fail")
	}
	model.Experiments = exps
}
