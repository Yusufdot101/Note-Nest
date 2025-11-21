package like

func (svc *likeService) addLike(userID, noteID int) error {
	l := &like{
		userID: userID,
		noteID: noteID,
	}
	return svc.repo.insert(l)
}

func (svc *likeService) removeLike(userID, noteID int) error {
	return svc.repo.delete(userID, noteID)
}

func (svc *likeService) noteIsLiked(userID, noteID int) (bool, error) {
	return svc.repo.isLiked(userID, noteID)
}
