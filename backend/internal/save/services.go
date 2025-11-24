package save

func (svc *saveService) saveNote(userID, noteID int) error {
	s := &save{
		userID: userID,
		noteID: noteID,
	}
	return svc.repo.insert(s)
}
