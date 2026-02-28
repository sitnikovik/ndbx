package wait

import (
	"time"
)

const (
	// Name is the name of the wait step.
	Name = "Wait Step"
	// Description is a brief explanation of what the wait step does.
	Description = "A simple step that waits for a specified duration before completing."
)

// WaitingStep is a simple step that wait for a specified duration.
type WaitingStep struct {
	// duration is the amount of time the step will sleep before completing.
	duration time.Duration
}

// NewWaitingStep creates a new WaitingStep that will wait for the specified duration.
func NewWaitingStep(dur time.Duration) *WaitingStep {
	return &WaitingStep{
		duration: dur,
	}
}

// Name returns the name of the step.
func (s *WaitingStep) Name() string {
	return Name
}

// Description returns a brief explanation of what the step does.
func (s *WaitingStep) Description() string {
	return Description
}
