package queue

import (
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

type Job[T river.JobArgs] = river.Job[T]

type JobInsertResult = *rivertype.JobInsertResult

type JobInsertUniqueOpts = river.UniqueOpts

type JobInsertOptions struct {
	Channel     string
	MaxAttempts int
	UniqueOpts  JobInsertUniqueOpts
	Tags        []string
}
