package usecases

func SetupListContentsUseCase(listContentsRepository ListContentsRepository) ListContentsUseCase {
	return func() ListContentsOutput {
		return listContentsRepository()
	}
}
