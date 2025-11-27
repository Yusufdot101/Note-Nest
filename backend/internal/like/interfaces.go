package like

type resourceType string

const (
	noteResource    resourceType = "note"
	commentResource resourceType = "comment"
)

type like struct {
	ID           int
	userID       int
	resourceType resourceType
	resourceID   int
}

type Repo interface {
	insert(l *like) error
	delete(l *like) error
	isLiked(l *like) (bool, error)
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

func (mr *mockRepo) delete(l *like) error {
	mr.deleteCalled = true
	return nil
}

func (mr *mockRepo) isLiked(l *like) (bool, error) {
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
