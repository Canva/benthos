package canva_aws

// Specific package for canva purposes because we do not need additional dependencies with elasticsearch and kafka
import (
	// Bring in the internal plugin definitions.
	_ "github.com/benthosdev/benthos/v4/internal/impl/aws"
)
