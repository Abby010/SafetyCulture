package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

// GetFoldersByOrgID returns folders filtered by organization ID
func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	res := make([]Folder, 0) // Initialize as empty slice
	for _, folder := range f.folders {
		if folder.OrgId == orgID {
			res = append(res, folder)
		}
	}
	return res
}

// GetAllChildFolders returns all child folders of a given folder for a specific organization
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Step 1: Check if the folder exists at all across all orgs
	var folderExists bool
	for _, folder := range f.folders {
		if folder.Name == name {
			folderExists = true
			break
		}
	}

	// If the folder does not exist at all, return a generic error
	if !folderExists {
		return nil, fmt.Errorf("folder does not exist")
	}

	// Step 2: Check if the folder exists in the specified organization
	foldersInOrg := f.GetFoldersByOrgID(orgID)
	var parentFolder *Folder
	for i := range foldersInOrg {
		if foldersInOrg[i].Name == name {
			parentFolder = &foldersInOrg[i]
			break
		}
	}

	// If the folder does not exist in the specified organization, return an org-specific error
	if parentFolder == nil {
		return nil, fmt.Errorf("folder does not exist in the specified organization")
	}

	// Step 3: Find all child folders within the same organization
	var childFolders []Folder
	prefix := parentFolder.Paths + "."
	for _, folder := range foldersInOrg {
		if strings.HasPrefix(folder.Paths, prefix) {
			childFolders = append(childFolders, folder)
		}
	}

	// Return child folders (or empty slice if none exist)
	return childFolders, nil
}
