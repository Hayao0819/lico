package vars

import "testing"

func TestErrNoSuchEntry(t *testing.T){
	testcase := []struct{
		fileErr fileErr
		wantErr string
		wantPath string
	}{
		{
			fileErr: *ErrNoSuchEntry("test"),
			wantErr: "no entry which has such path",
			wantPath: "test",
		},
	}

	for _, tc := range testcase{
		if tc.fileErr.Err.Error() != tc.wantErr{
			t.Errorf("ErrNoSuchEntry() error: %v, want: %v", tc.fileErr.Err.Error(), tc.wantErr)
		}
		if tc.fileErr.Path != tc.wantPath{
			t.Errorf("ErrNoSuchEntry() path: %v, want: %v", tc.fileErr.Path, tc.wantPath)
		}
	}
}


