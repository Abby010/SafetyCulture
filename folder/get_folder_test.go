package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestFolder(name, paths string, orgID uuid.UUID) folder.Folder {
	return folder.Folder{
		Name:  name,
		Paths: paths,
		OrgId: orgID,
	}
}

func Test_GetFoldersByOrgID(t *testing.T) {
	orgID1 := uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))
	orgID2 := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a"))

	folders := []folder.Folder{
		createTestFolder("creative-scalphunter", "creative-scalphunter", orgID1),
		createTestFolder("clear-arclight", "creative-scalphunter.clear-arclight", orgID1),
		createTestFolder("nearby-secret", "nearby-secret", orgID2),
	}

	tests := []struct {
		name  string
		orgID uuid.UUID
		want  []folder.Folder
	}{
		{
			name:  "Get folders for orgID1",
			orgID: orgID1,
			want:  []folder.Folder{folders[0], folders[1]},
		},
		{
			name:  "Get folders for orgID2",
			orgID: orgID2,
			want:  []folder.Folder{folders[2]},
		},
		{
			name:  "No folders for unknown orgID",
			orgID: uuid.Must(uuid.NewV4()),
			want:  []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(folders)
			got := driver.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetAllChildFolders(t *testing.T) {
	orgID1 := uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))
	orgID2 := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a"))

	folders := []folder.Folder{
		createTestFolder("creative-scalphunter", "creative-scalphunter", orgID1),
		createTestFolder("clear-arclight", "creative-scalphunter.clear-arclight", orgID1),
		createTestFolder("topical-micromax", "creative-scalphunter.clear-arclight.topical-micromax", orgID1),
		createTestFolder("nearby-secret", "nearby-secret", orgID2),
	}

	tests := []struct {
		name        string
		orgID       uuid.UUID
		folderName  string
		want        []folder.Folder
		wantErr     bool
		expectedErr string
	}{
		{
			name:       "Get child folders for creative-scalphunter in orgID1",
			orgID:      orgID1,
			folderName: "creative-scalphunter",
			want:       []folder.Folder{folders[1], folders[2]},
			wantErr:    false,
		},
		{
			name:       "Get child folders for clear-arclight in orgID1",
			orgID:      orgID1,
			folderName: "clear-arclight",
			want:       []folder.Folder{folders[2]},
			wantErr:    false,
		},
		{
			name:       "No child folders for topical-micromax in orgID1",
			orgID:      orgID1,
			folderName: "topical-micromax",
			want:       []folder.Folder{},
			wantErr:    false,
		},
		{
			name:        "Invalid folder name in orgID1",
			orgID:       orgID1,
			folderName:  "invalid-folder",
			want:        nil,
			wantErr:     true,
			expectedErr: "folder does not exist",
		},
		{
			name:        "Folder exists but not in specified org",
			orgID:       orgID1,
			folderName:  "nearby-secret",
			want:        nil,
			wantErr:     true,
			expectedErr: "folder does not exist in the specified organization",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(folders)
			got, err := driver.GetAllChildFolders(tt.orgID, tt.folderName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllChildFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr {
				assert.EqualError(t, err, tt.expectedErr)
				return
			}
			// If no error, compare the results
			if err == nil {
				assert.ElementsMatch(t, tt.want, got)
			}
		})
	}
}
