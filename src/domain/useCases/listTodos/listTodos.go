package usecases

func SetupListTodosUseCase(listTodosRepository ListTodosRepository) ListTodosUseCase {
	return func() ListTodosOutput {
		return listTodosRepository()
	}
}
