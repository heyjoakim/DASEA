package depTypes

type DepType string

const (
	DevDependency DepType = "dev"
	Dependency            = "runtime"
)
