package servicemodels

// ApplicationInfo, will hold information about each and every microservice.
type ApplicationInfo struct {
	Name        string
	Version     string
	Service     string
	Environment string
	CommitDate  string
	BuildDate   string
	BuildHash   string
}
