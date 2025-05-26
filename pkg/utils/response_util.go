package utils

import (
	"net/http"

	"github.com/kweusuf/url-shortner/pkg/model"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

func ConstructResponse(res interface{}, err error) (interface{}, error) {
	if err != nil {
		log.Error(err.Error())
		resp := model.GetConfResponse{
			HttpStatus: http.StatusBadRequest,
			Body:       err.Error(),
		}
		return resp, nil
	}
	resp := model.GetConfResponse{
		HttpStatus: http.StatusOK,
		Body:       res,
	}
	return resp, nil
}
