package usecases

import "app/src/domain/entities"

type ListContentsRepository func() []entities.Content
type ListContentsDeps struct {
	listContentsRepository ListContentsRepository
}
type ListContentsUseCase func() ListContentsOutput
type ListContentsOutput []entities.Content
