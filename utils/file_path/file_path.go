package file_paths

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUserProfileFolder(userObjID string) string {
	return fmt.Sprintf("user/%s_%s.png", userObjID, uuid.New().String())
}
