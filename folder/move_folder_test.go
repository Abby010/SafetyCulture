package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	orgID1 := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID2 := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")

	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: orgID1},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
		{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
		{Name: "golf", Paths: "golf", OrgId: orgID1},
	}

	tests := []struct {
		name        string
		source      string
		destination string
		expected    []folder.Folder
		expectErr   bool
	}{
		{
			name:        "Move bravo under delta",
			source:      "bravo",
			destination: "delta",
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
				{Name: "golf", Paths: "golf", OrgId: orgID1},
			},
			expectErr: false,
		},
		{
			name:        "Move bravo under itself",
			source:      "bravo",
			destination: "bravo",
			expected:    nil,
			expectErr:   true,
		},
		{
			name:        "Move bravo under charlie (child)",
			source:      "bravo",
			destination: "charlie",
			expected:    nil,
			expectErr:   true,
		},
		{
			name:        "Move bravo to different org (foxtrot)",
			source:      "bravo",
			destination: "foxtrot",
			expected:    nil,
			expectErr:   true,
		},
		{
			name:        "Move invalid folder",
			source:      "invalid_folder",
			destination: "delta",
			expected:    nil,
			expectErr:   true,
		},
		{
			name:        "Move bravo to invalid folder",
			source:      "bravo",
			destination: "invalid_folder",
			expected:    nil,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(folders)
			got, err := driver.MoveFolder(tt.source, tt.destination)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
