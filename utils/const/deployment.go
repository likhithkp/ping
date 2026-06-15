package _const

type DeploymentType string

const (
	Deployment_Production DeploymentType = "production"
	Deployment_Dev        DeploymentType = "dev"
	Deployment_LocalDev   DeploymentType = "local-dev"
)
