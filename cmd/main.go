package main

import (
	repositoryMap "github.com/zelta-7/cache/pkg/repository/map"
	repositoryQueue "github.com/zelta-7/cache/pkg/repository/queue"
	serviceMap "github.com/zelta-7/cache/pkg/service/map"
	serviceQueue "github.com/zelta-7/cache/pkg/service/queue"
)

func main() {
	queueRepository := repositoryQueue.NewQueueRepo()
	mapRepository := repositoryMap.NewMapRepo()

	queueService := serviceQueue.NewQueueService(queueRepository)
	mapService := serviceMap.NewMapService(mapRepository)

}
