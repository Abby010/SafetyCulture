package folder

import "github.com/gofrs/uuid"

// IDriver defines the interface with the methods required for managing folders
type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error)
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

// driver struct stores the folder data.
type driver struct {
	// Slice to store all folders
	folders []Folder
}

// NewDriver initializes a new driver with a list of folders
func NewDriver(folders []Folder) IDriver {
	return &driver{
		folders: folders,
	}
}
