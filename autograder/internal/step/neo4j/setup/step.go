package setup

import "context"

const (
	// Name is the name of the step.
	Name = "Neo4j Setup"
	// Description provides a brief description of the step.
	Description = "Sets up the Neo4j environment for tests."
)

// neo4jClient defines the interface for interacting
// with Neo4j to perform the setup job.
type neo4jClient interface {
	// DeleteAll removes all nodes and relationships from the database.
	DeleteAll(ctx context.Context) error
}

// Step represents the Neo4j setup step for tests.
type Step struct {
	// cli is the Neo4j client used to perform the setup job.
	cli neo4jClient
}

// NewStep creates a new Step instance
// with the provided Neo4j client.
func NewStep(cli neo4jClient) *Step {
	return &Step{
		cli: cli,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return Name
}

// Description returns a brief explanation of what the step does.
func (s *Step) Description() string {
	return Description
}