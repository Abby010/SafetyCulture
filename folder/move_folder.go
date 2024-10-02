package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders

	// Step 1: Find the source and destination folders
	var srcFolder *Folder
	var dstFolder *Folder

	for i, folder := range folders {
		if folder.Name == name {
			srcFolder = &folders[i]
		}
		if folder.Name == dst {
			dstFolder = &folders[i]
		}
	}

	// Step 2: Error handling for invalid or conflicting operations
	if srcFolder == nil {
		return nil, fmt.Errorf("source folder %s does not exist", name) // lowercase "source"
	}
	if dstFolder == nil {
		return nil, fmt.Errorf("destination folder %s does not exist", dst) // lowercase "destination"
	}
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, fmt.Errorf("cannot move a folder to a different organization") // lowercase "cannot"
	}
	if name == dst {
		return nil, fmt.Errorf("cannot move a folder to itself") // lowercase "cannot"
	}
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths+".") {
		return nil, fmt.Errorf("cannot move a folder to a child of itself") // lowercase "cannot"
	}

	// Step 3: Update paths for the source folder and its child folders
	newPrefix := dstFolder.Paths + "." + srcFolder.Name
	oldPrefix := srcFolder.Paths

	for i, folder := range folders {
		// Update paths for all folders in the subtree of srcFolder
		if strings.HasPrefix(folder.Paths, oldPrefix) {
			newPath := strings.Replace(folder.Paths, oldPrefix, newPrefix, 1)
			folders[i].Paths = newPath
		}
	}

	return folders, nil
}
