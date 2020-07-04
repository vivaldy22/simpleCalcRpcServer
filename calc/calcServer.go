package calc

import (
	"context"
	"github.com/edwardsuwirya/simpleCalcRpcServer/api"
)

type Server struct {
}

func (s *Server) AddNumber(ctx context.Context, in *api.AdditionMessage) (*api.AdditionResultMessage, error) {
	num1 := in.Num1
	num2 := in.Num2
	var result = num1 + num2
	resultMessage := &api.AdditionResultMessage{ResNum: result}

	return resultMessage, nil
}
