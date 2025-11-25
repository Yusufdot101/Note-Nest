package like

type like struct {
	ID     int
	userID int
	noteID int
}

type Repo interface {
	insert(l *like) error
	delete(userID, noteID int) error
	isLiked(userID, noteID int) (bool, error)
}

type mockRepo struct {
	insertCalled  bool
	deleteCalled  bool
	isLikedCalled bool
}

func (mr *mockRepo) insert(l *like) error {
	mr.insertCalled = true
	return nil
}

func (mr *mockRepo) delete(userID, noteID int) error {
	mr.deleteCalled = true
	return nil
}

func (mr *mockRepo) isLiked(userID, noteID int) (bool, error) {
	mr.isLikedCalled = true
	return true, nil
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
