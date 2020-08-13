package util

import (
	"ab_division/log"
	"ab_division/rpc/proto"
	"context"
	"errors"
	"runtime/debug"
)

var (
	ParameterError = errors.New("missing necessary parameter")
	InternalError  = errors.New("internal error")
)

type DivisionServer struct {
}

func (s *DivisionServer) GetABDivision(ctx context.Context, request *proto.User) (response *proto.ABResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.DataLogger.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			response = &proto.ABResponse{ResponseStatus: proto.Status_FAILED, Massage: "Internal error"}
			err = InternalError
			return
		}
	}()

	if request.Uid == 0 || request.Udid == "" {
		response = &proto.ABResponse{ResponseStatus: proto.Status_FAILED, Massage: "Missing necessary parameter"}
		return response, ParameterError
	}
	result := getABDivisions(request.Uid, request.Udid, request.BusinessId)
	res := proto.ABResponse{
		ResponseStatus: proto.Status_SUCCEED,
		Massage:        "Status_SUCCEED",
		Data:           result,
	}
	return &res, nil
}

func (s *DivisionServer) GetExperiment(ctx context.Context, request *proto.BusinessInfo) (response *proto.ExperimentInfoResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.DataLogger.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			response = &proto.ExperimentInfoResponse{ResponseStatus: proto.Status_FAILED, Massage: "Internal error"}
			err = InternalError
			return
		}
	}()

	if request.BusinessId == "" {
		response = &proto.ExperimentInfoResponse{ResponseStatus: proto.Status_FAILED, Massage: "Missing necessary parameter"}
		return response, ParameterError
	}
	result := getExperimentsByBusinessId(request.BusinessId)
	res := proto.ExperimentInfoResponse{
		ResponseStatus: proto.Status_SUCCEED,
		Massage:        "Status_SUCCEED",
		Data:           result,
	}
	return &res, nil

}
