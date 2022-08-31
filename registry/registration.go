package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}

type patchEntry struct {
	Name ServiceName
	URL  string
}
