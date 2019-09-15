package services

import "encoding/json"

type MarshallerService interface {
	Marshal(i interface{}) ([]byte, error)
	Unmarshal(data []byte, i interface{}) error
}

type marshallerService struct {

}

func (srv *marshallerService) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (srv *marshallerService) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewMarshallerService() MarshallerService {
	return &marshallerService{}
}

