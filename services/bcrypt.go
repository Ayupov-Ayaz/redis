package services

import "golang.org/x/crypto/bcrypt"

type BcryptService interface {
	HashPassword(str string) (hash string, err error)
	CompareHashPassword(password, hash string) bool
}

type bcryptService struct {
	cost int
}

func NewBcryptService(cost int) BcryptService {
	return &bcryptService{
		cost:cost,
	}
}

func (srv *bcryptService) HashPassword(str string) (hash string, err error) {
	var bHash []byte
	bHash, err = bcrypt.GenerateFromPassword([]byte(str), srv.cost)
	if err != nil {
		return "", err
	}
	hash = string(bHash)
	return
}

func (srv *bcryptService) CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


