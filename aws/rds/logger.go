package rds

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
