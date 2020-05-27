package env

type Runtime string

const (
	// Production implies the server is running under production environment.
	Production Runtime = "production"
	// PreProd implies the server is running under pre-prod environment.
	PreProd Runtime = "pre-prod"
	// Staging implies the server is running under staging environment.
	Staging Runtime = "staging"
	// Testing implies the server is running under testing environment.
	Testing Runtime = "testing"
	// Development implies the server is running on developer's local machine.
	Development Runtime = "development"
)

type Deployment struct {
	runtime Runtime
}

func (d Deployment) IsProduction() bool {
	return d.runtime == Production
}

func (d Deployment) IsPreProd() bool {
	return d.runtime == PreProd
}

func (d Deployment) IsStaging() bool {
	return d.runtime == Staging
}

func (d Deployment) IsTesting() bool {
	return d.runtime == Testing
}

func (d Deployment) IsDevelopment() bool {
	return d.runtime == Development
}

func NewDeployment(runtime Runtime) Deployment {
	return Deployment{runtime: runtime}
}
