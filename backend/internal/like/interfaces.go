package like

type like struct {
	ID     int
	userID int
	noteID int
}

type Repo interface {
	insert(l *like) error
	delete(userID, noteID int) error
}

type likeService struct {
	repo Repo
}

type likeHandler struct {
	svc *likeService
}

func newHandler(svc *likeService) *likeHandler {
	return &likeHandler{
		svc: svc,
	}
}
