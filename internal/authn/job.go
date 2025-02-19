package authn

import (
	"butta/pkg/logger"
	"butta/pkg/queue"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SendPasswordResetLinkJobArgs struct {
	RecipientName string `json:"recipient-name"`
	Email         string `json:"email"`
	Link          string `json:"reset-link"`
}

func (w SendPasswordResetLinkJobArgs) Kind() string {
	return "email-password-reset-link"
}

type SendPasswordResetLinkWorker struct {
	PsqlPool *pgxpool.Pool
	// An embedded WorkerDefaults sets up default methods to fulfill the rest of
	// the Worker interface:
	queue.WorkerDefaults[SendPasswordResetLinkJobArgs]
}

func (w *SendPasswordResetLinkWorker) Work(ctx context.Context, job *queue.Job[SendPasswordResetLinkJobArgs]) error {

	for {
		select {
		case <-ctx.Done():
			logger.Info("job canceled")
			return ctx.Err()
		default:
			//todo send email to the user here
			logger.Fatal("Sending emails not implement", "email", job.Args.Email, "link", job.Args.Link)
			return nil
		}
	}
}
