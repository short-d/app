package env

type runtime string

const (
	// Production implies the server is running under production environment.
	Production runtime = "production"
	// PreProd implies the server is running under pre-prod environment.
	PreProd runtime = "pre-prod"
	// Staging implies the server is running under staging environment.
	Staging runtime = "staging"
	// Testing implies the server is running under testing environment.
	Testing runtime = "testing"
	// Development implies the server is running on developer's local machine.
	Development runtime = "development"
)

type Deployment struct {
	env Env
}

func (d Deployment) IsProduction() bool {
	return d.getEnv() == Production
}

func (d Deployment) IsPreProd() bool {
	return d.getEnv() == PreProd
}

func (d Deployment) IsStaging() bool {
	return d.getEnv() == Staging
}

func (d Deployment) IsTesting() bool {
	return d.getEnv() == Testing
}

func (d Deployment) IsLocal() bool {
	return d.getEnv() == Development
}

func (d Deployment) getEnv() runtime {
	return runtime(d.env.GetVar("ENV", string(Development)))
}

func NewDeployment(env Env) Deployment {
	return Deployment{env: env}
}
