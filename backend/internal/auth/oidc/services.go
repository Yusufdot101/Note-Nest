package oidc

import (
	"errors"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
)

func (os *oidcService) findOrInsertUser(ui *userInfo) (string, string, error) {
	u, err := os.userSvc.GetUserByProvider(ui.ProviderName, ui.Sub)
	if err != nil && !errors.Is(err, customerrors.ErrNoRecord) {
		return "", "", err
	} else if errors.Is(err, customerrors.ErrNoRecord) {
		err = os.repo.insert(ui)
		if err != nil {
			return "", "", err
		}
	} else {
		ui.UserID = u.ID
	}

	return os.tokenSvc.GetAccessAndRefreshTokens(ui.UserID)
}
