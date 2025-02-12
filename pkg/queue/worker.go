package queue

import (
	"github.com/riverqueue/river"
)

type WorkerDefaults[T river.JobArgs] = river.WorkerDefaults[T]

type Worker[T river.JobArgs] = river.Worker[T]
