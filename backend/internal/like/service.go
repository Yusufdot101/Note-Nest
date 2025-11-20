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
