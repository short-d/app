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
	env Env
}

func (d Deployment) IsProduction() bool {
	return d.GetRuntime() == Production
}

func (d Deployment) IsPreProd() bool {
	return d.GetRuntime() == PreProd
}

func (d Deployment) IsStaging() bool {
	return d.GetRuntime() == Staging
}

func (d Deployment) IsTesting() bool {
	return d.GetRuntime() == Testing
}

func (d Deployment) IsLocal() bool {
	return d.GetRuntime() == Development
}

func (d Deployment) GetRuntime() Runtime {
	return Runtime(d.env.GetVar("ENV", string(Development)))
}

func NewDeployment(env Env) Deployment {
	return Deployment{env: env}
}
