package apis

// Step 3 scaffold: concrete CRD schemas are defined in Step 4.
type WorkspacePhase string

type SparkRunPhase string

type OperationPhase string

const (
	WorkspacePhasePending      WorkspacePhase = "Pending"
	WorkspacePhaseProvisioning WorkspacePhase = "Provisioning"
	WorkspacePhaseReady        WorkspacePhase = "Ready"
	WorkspacePhaseFailed       WorkspacePhase = "Failed"
	WorkspacePhaseDeleting     WorkspacePhase = "Deleting"
)

const (
	SparkRunPhaseQueued    SparkRunPhase = "Queued"
	SparkRunPhaseStarting  SparkRunPhase = "Starting"
	SparkRunPhaseRunning   SparkRunPhase = "Running"
	SparkRunPhaseSucceeded SparkRunPhase = "Succeeded"
	SparkRunPhaseFailed    SparkRunPhase = "Failed"
	SparkRunPhaseCanceled  SparkRunPhase = "Canceled"
)

const (
	OperationPhasePending    OperationPhase = "Pending"
	OperationPhaseInProgress OperationPhase = "InProgress"
	OperationPhaseSucceeded  OperationPhase = "Succeeded"
	OperationPhaseFailed     OperationPhase = "Failed"
)
