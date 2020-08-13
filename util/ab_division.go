package util

import (
	"ab_division/model"
	"ab_division/rpc/proto"
	"crypto/md5"
	"encoding/hex"
	"math/big"
	"strconv"
	"time"
)

func ID2Bigint(id string) *big.Int {
	m5 := md5.New()
	m5.Write([]byte(id))
	st1 := m5.Sum(nil)
	low8Num := st1[len(st1)-4:]
	temp := hex.EncodeToString(low8Num)
	n := new(big.Int)
	n, _ = n.SetString(temp, 16)
	return n
}

func isNewUser(udid string) bool {
	//可替换为特殊用户甄别逻辑，如新用户，付费用户等
	return ID2Bigint(udid).Int64()%2 == 0
}

func getABDivisions(uid int64, udid, businessId string) []*proto.ABDivision {
	exps := model.Experiments
	var divisions []*proto.ABDivision
	for i := range exps {
		if exps[i].NewUserOnly && !isNewUser(udid) {
			continue
		}
		if businessId != exps[i].BusinessId {
			continue
		}
		intUdid := ID2Bigint(udid + exps[i].Salt).Int64()

		index := calculationDivision(exps[i].Buckets, intUdid, exps[i].BucketMap)
		assignment := exps[i].Buckets[index].Assignment
		passthrough := proto.Passthrough{
			Uid:        uid,
			Udid:       udid,
			ExpId:      strconv.Itoa(exps[i].ExpId),
			BucketName: exps[i].Buckets[index].BucketName,
			BucketId:   strconv.Itoa(exps[i].Buckets[index].BucketId),
			ServeTime:  time.Now().Unix(),
		}

		division := proto.ABDivision{Assignment: assignment, Passthrough: &passthrough}
		divisions = append(divisions, &division)
	}
	return divisions
}

func getExperimentsByBusinessId(businessId string) []*proto.ExperimentInfo {
	var experimentInfos []*proto.ExperimentInfo
	exps := model.Experiments
	for i := range exps {
		if businessId != exps[i].BusinessId {
			continue
		}
		experiment := proto.ExperimentInfo{
			ExpId:       strconv.Itoa(exps[i].ExpId),
			ExpName:     exps[i].ExpName,
			State:       exps[i].State,
			NewUserOnly: exps[i].NewUserOnly,
		}
		experimentInfos = append(experimentInfos, &experiment)
	}
	return experimentInfos
}

//分流思路，将分流比例转化为对100求余
//例如 intUdid % 100 =32
//bucketMap[BucketId1]=20,bucketMap[BucketId1]=40,bucketMap[BucketId1]=100
//则分流至BucketId1，返回index=1
func calculationDivision(buckets []model.Bucket, intUdid int64, bucketMap map[int]int) int {
	remainder := intUdid % 100
	for i, bucket := range buckets {
		if limit, ok := bucketMap[bucket.BucketId]; ok {
			if remainder < int64(limit) {
				return i
			}
		}
	}
	return -1
}
