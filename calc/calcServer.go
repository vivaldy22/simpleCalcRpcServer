package calc

import (
	"context"
	"errors"

	"github.com/edwardsuwirya/simpleCalcRpcServer/api"
)

type Server struct {
}

func (s *Server) AddNumber(ctx context.Context, in *api.ParameterMessage) (*api.ResultMessage, error) {
	num1 := in.Num1
	num2 := in.Num2
	var result = num1 + num2
	resultMessage := &api.ResultMessage{ResNum: result}

	return resultMessage, nil
}

func (s *Server) SubNumber(ctx context.Context, in *api.ParameterMessage) (*api.ResultMessage, error) {
	num1 := in.Num1
	num2 := in.Num2
	var result = num1 - num2
	resultMessage := &api.ResultMessage{ResNum: result}

	return resultMessage, nil
}

func (s *Server) MulNumber(ctx context.Context, in *api.ParameterMessage) (*api.ResultMessage, error) {
	num1 := in.Num1
	num2 := in.Num2
	var result = num1 * num2
	resultMessage := &api.ResultMessage{ResNum: result}

	return resultMessage, nil
}

func (s *Server) DivNumber(ctx context.Context, in *api.ParameterMessage) (*api.ResultMessage, error) {
	num1 := in.Num1
	num2 := in.Num2
	// resultMessage := new(api.ResultMessage)
	if num2 == 0 {
		return nil, errors.New("cannot divide by 0")
		// resultMessage = &api.ResultMessage{Error: &api.Error{Code: "123", Message: "cannot divide by 0"}}
	}
	var result = num1 / num2
	resultMessage := &api.ResultMessage{ResNum: result}

	return resultMessage, nil
}
