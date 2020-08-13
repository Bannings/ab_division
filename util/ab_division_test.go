package util

import (
	"ab_division/model"
	"strconv"
	"testing"
)

//测试1000个不同的udid分流区间段
func TestID2Bigint(t *testing.T) {
	AllocationMap := make(map[int]int)
	AllocationMap[20] = 0
	AllocationMap[40] = 0
	AllocationMap[100] = 0
	//udid := "86212311192417"
	udid := "09fd27abde9478fe"
	for i := 0; i < 1000; i++ {
		udid = udid + strconv.Itoa(i)
		intUdid := ID2Bigint(udid + "BGtyHN12U").Int64()
		index := intUdid % 100
		if index < 20 {
			count := AllocationMap[20]
			AllocationMap[20] = count + 1
		} else if index < 40 {
			count := AllocationMap[40]
			AllocationMap[40] = count + 1
		} else if index < 100 {
			count := AllocationMap[100]
			AllocationMap[100] = count + 1
		}
	}
	t.Logf("分流1总数：%d", AllocationMap[20])
	t.Logf("分流2总数：%d", AllocationMap[40])
	t.Logf("分流3总数：%d", AllocationMap[100])
}

func TestCalculateBucketMap(t *testing.T) {
	bucket1 := model.Bucket{BucketId: 1, BucketName: "test1", Assignment: "100/1", AllocationPercent: 20}
	bucket2 := model.Bucket{BucketId: 2, BucketName: "test2", Assignment: "100/2", AllocationPercent: 20}
	bucket3 := model.Bucket{BucketId: 3, BucketName: "test3", Assignment: "100/3", AllocationPercent: 60}
	buckets := []model.Bucket{bucket1, bucket2, bucket3}
	bucketMap, err := calculateBucketMap(buckets)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(bucketMap)
	}
}

func TestCalculationDivision(t *testing.T) {
	bucket1 := model.Bucket{BucketId: 1, BucketName: "test1", Assignment: "100/1", AllocationPercent: 20}
	bucket2 := model.Bucket{BucketId: 2, BucketName: "test2", Assignment: "100/2", AllocationPercent: 20}
	bucket3 := model.Bucket{BucketId: 3, BucketName: "test3", Assignment: "100/3", AllocationPercent: 60}
	buckets := []model.Bucket{bucket1, bucket2, bucket3}
	bucketMap, err := calculateBucketMap(buckets)
	if err != nil {
		t.Fatal(err)
	}
	udid := "865162311192325"
	intUdid := ID2Bigint(udid + "BGtyHN12U").Int64()
	index := calculationDivision(buckets, intUdid, bucketMap)
	if index == -1 {
		t.Fatal("dvision error")
	} else {
		t.Log(index)
	}
}
