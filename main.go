package main

import (
	"go.uber.org/zap"
	"mail-service/common"
	"mail-service/common/envoriment"
)

func main() {
	logger := common.Logger("false")
	zap.ReplaceGlobals(logger)

	env := envoriment.GetEnvoriment()
	kafka := common.KafkaEvent{
		Network: env.Network,
		GroupId: env.GroupId,
		Address: env.Address,
		Topic:   env.Topic,
		Partion: env.Partion,
	}

	kafka.Consumer()

}
