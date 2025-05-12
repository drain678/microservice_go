package service

import "github.com/drain678/microservice_go.git/internal/proxyproto"

// RespondError ...
func ConnectRespondError(code uint32, msg string) (*proxyproto.ConnectResponse, error) {
	return &proxyproto.ConnectResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: msg,
		},
	}, nil
}

func PublishResponseError(code uint32, msg string) (*proxyproto.PublishResponse, error) {
	return &proxyproto.PublishResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: msg,
		},
	}, nil
}

func SubscribeResponseError(code uint32, msg string) (*proxyproto.SubscribeResponse, error) {
	return &proxyproto.SubscribeResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: msg,
		},
	}, nil

}
