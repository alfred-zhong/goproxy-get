package git

import (
	"testing"
)

func TestConfig(t *testing.T) {
	type args struct {
		key    string
		value  string
		global bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"qwerty", args{"qwerty", "foo", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Config(tt.args.key, tt.args.value, tt.args.global); (err != nil) != tt.wantErr {
				t.Fatalf("Config() error = %v, wantErr %v", err, tt.wantErr)
			}
			Unset(tt.args.key, tt.args.global)
		})
	}
}

func TestUnset(t *testing.T) {
	type args struct {
		key    string
		global bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"qwerty", args{"qwerty", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unset(tt.args.key, tt.args.global); (err != nil) != tt.wantErr {
				t.Errorf("Unset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveSection(t *testing.T) {
	type args struct {
		sectionKey string
		global     bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"qwerty", args{"qwery", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveSection(tt.args.sectionKey, tt.args.global); (err != nil) != tt.wantErr {
				t.Errorf("RemoveSection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListNameOnly(t *testing.T) {
	names, err := ListNameOnly(true)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) == 0 {
		t.Errorf("names should not be empty, got length: %d", len(names))
	}
}
