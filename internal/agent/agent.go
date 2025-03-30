package agent

import (
	"context"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
)

// Agent может быть любой сущностью, реализующей необходимые методы
type Agent interface {
	getAge(ctx context.Context, name string) (int, error)
	getGender(ctx context.Context, name string) (string, error)
	getNationalize(ctx context.Context, name string) (string, error)

	EnrichInformation(ctx context.Context, user e.User) (*e.User, error)
}
