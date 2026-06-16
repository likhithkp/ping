package aws

import "go.uber.org/fx"

var Module = fx.Module("utils-aws",
	fx.Provide(
		NewAWSConfig,
		NewSesClient,
		NewS3Client,
	),
)
