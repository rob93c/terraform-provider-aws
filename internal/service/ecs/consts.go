package ecs

import (
	"time"
)

const (
	propagationTimeout = 2 * time.Minute
)

const (
	clusterStatusActive         = "ACTIVE"
	clusterStatusDeprovisioning = "DEPROVISIONING"
	clusterStatusInactive       = "INACTIVE"
	clusterStatusProvisioning   = "PROVISIONING"
)
