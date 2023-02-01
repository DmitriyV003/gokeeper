package core

import (
	"gokeeper/internal/proto"
)

type SecretService struct {
}

func (s *SecretService) CreateLoginSecret(req *proto.CreateLoginSecretRequest) (string, error) {

}
