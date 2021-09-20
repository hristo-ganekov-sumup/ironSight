module github.com/hristo-ganekov-sumup/ironSight

go 1.16

replace github.com/hristo-ganekov-sumup/ironSight/internal/tfstate => ./internal/tfstate

replace github.com/hristo-ganekov-sumup/ironSight/internal/sg => ./internal/sg

require (
	github.com/aws/aws-sdk-go v1.40.40
	github.com/google/go-cmp v0.5.6
)
