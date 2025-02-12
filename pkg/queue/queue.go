package queue

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"time"
)

type JobArgs = river.JobArgs

type Dispatchable struct {
	client *river.Client[pgx.Tx]
	err    error
}

func (d Dispatchable) Queue(ctx context.Context, args JobArgs, options *JobInsertOptions) (JobInsertResult, error) {
	if d.err != nil {
		return nil, d.err
	}

	insertOptions := &river.InsertOpts{
		Queue:       options.Channel,
		MaxAttempts: options.MaxAttempts,
		Tags:        options.Tags,
		UniqueOpts:  river.UniqueOpts{},
	}

	return d.client.Insert(ctx, args, insertOptions)
}

func (d Dispatchable) Schedule(ctx context.Context, args JobArgs, scheduledAt time.Time, options *JobInsertOptions) (JobInsertResult, error) {
	if d.err != nil {
		return nil, d.err
	}

	if scheduledAt.IsZero() {
		return nil, errors.New("job scheduled time not allowed")
	}

	var insertOptions *river.InsertOpts
	if options != nil {
		insertOptions = &river.InsertOpts{
			Queue:       string(options.Channel),
			MaxAttempts: options.MaxAttempts,
			Tags:        options.Tags,
			UniqueOpts:  river.UniqueOpts{},
			ScheduledAt: scheduledAt,
		}
	}

	return d.client.Insert(ctx, args, insertOptions)
}

func With(dbPool *pgxpool.Pool) *Dispatchable {

	client, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{})

	return &Dispatchable{
		client: client,
		err:    err,
	}
}
