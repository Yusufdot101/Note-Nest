package save

type save struct {
	userID, noteID int
}

type repo interface {
	insert(s *save) error
}

type saveService struct {
	repo repo
}

type saveHandler struct {
	svc *saveService
}

func newHandler(svc *saveService) *saveHandler {
	return &saveHandler{
		svc: svc,
	}
}
