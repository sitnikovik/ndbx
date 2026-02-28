package autograder

// Autograder performs automated checks for a labwork by running a series of grading jobs.
type Autograder struct {
	// jobs is the list of grading jobs
	// that will be executed as part of the autograder process.
	jobs []Runner
}

// NewAutograder creates a new Autograder instance
// with the provided jobs to run.
func NewAutograder(jobs ...Runner) *Autograder {
	return &Autograder{
		jobs: jobs,
	}
}
