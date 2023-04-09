package routers

import (
	db "armadabackend/services/databaseServices"
)

type userWBody struct {
	Message string           `json:"message"`
	Result  []db.UsersRecord `json:"result"`
}
