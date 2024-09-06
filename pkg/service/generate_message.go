package service

import (
	"encoding/json"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func GenerateMessage() []types.Message {

	countries := GenerateCountries()
	var list []types.Message

	for _, item := range countries {
		values, _ := json.Marshal(item)
		msm := types.Message{
			Topic: "bananinha",
			Value: values,
		}
		list = append(list, msm)
		msm = types.Message{
			Topic: "abobrinha",
			Value: values,
		}
		list = append(list, msm)
	}
	return list

}
