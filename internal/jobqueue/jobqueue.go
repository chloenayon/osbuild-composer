// Package jobqueue provides a generic interface to a simple job queue.
//
// Jobs are pushed to the queue with Enqueue(). Workers call Dequeue() to
// receive a job and FinishJob() to report one as finished.
//
// Each job has a type and arguments corresponding to this type. These are
// opaque to the job queue, but it mandates that the arguments must be
// serializable to JSON. Similarly, a job's result has opaque result arguments
// that are determined by its type.
//
// A job can have dependencies. It is not run until all its dependencies have
// finished.
package jobqueue

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// JobQueue is an interface to a simple job queue. It is safe for concurrent use.
type JobQueue interface {
	// Enqueues a job.
	//
	// `args` must be JSON-serializable and fit the given `jobType`, i.e., a worker
	// that is running that job must know the format of `args`.
	//
	// All dependencies must already exist, but the job isn't run until all of them
	// have finished.
	//
	// Returns the id of the new job, or an error.
	Enqueue(jobType string, args interface{}, dependencies []uuid.UUID) (uuid.UUID, error)

	// Dequeues a job, blocking until one is available.
	//
	// Waits until a job with a type of any of `jobTypes` is available, or `ctx` is
	// canceled.
	//
	// All jobs in `jobTypes` must take the same type of `args`, corresponding to
	// the one that was passed to Enqueue().
	//
	// Returns the job's id or an error.
	Dequeue(ctx context.Context, jobTypes []string, args interface{}) (uuid.UUID, error)

	// Mark the job with `id` as finished. `result` must fit the associated
	// job type and must be serializable to JSON.
	FinishJob(id uuid.UUID, result interface{}) error

	// Cancel a job. Does nothing if the job has already finished.
	CancelJob(id uuid.UUID) error

	// Returns the current status of the job, in the form of three times:
	// queued, started, and finished. `started` and `finished` might be the
	// zero time (check with t.IsZero()), when the job is not running or
	// finished, respectively.
	//
	// If the job is finished, its result will be returned in `result`.
	JobStatus(id uuid.UUID, result interface{}) (queued, started, finished time.Time, canceled bool, err error)
}

var (
	ErrNotExist   = errors.New("job does not exist")
	ErrNotRunning = errors.New("job is not running")
	ErrCanceled   = errors.New("job ws canceled")
)
