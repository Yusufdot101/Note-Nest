package like

func (svc *likeService) addLike(userID int, resourceType resourceType, resourceID int) error {
	l := &like{
		userID:       userID,
		resourceType: resourceType,
		resourceID:   resourceID,
	}
	return svc.repo.insert(l)
}

func (svc *likeService) removeLike(userID int, resourceType resourceType, resourceID int) error {
	l := &like{
		userID:       userID,
		resourceType: resourceType,
		resourceID:   resourceID,
	}

	return svc.repo.delete(l)
}

func (svc *likeService) resourceIsLiked(userID int, resourceType resourceType, resourceID int) (bool, error) {
	l := &like{
		userID:       userID,
		resourceType: resourceType,
		resourceID:   resourceID,
	}
	return svc.repo.isLiked(l)
}
