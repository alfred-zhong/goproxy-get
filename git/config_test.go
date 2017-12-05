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
		{"http.proxy", args{"http.proxy", "http://127.0.0.1", true}, false},
		{"qwerty", args{"qwerty", "foo", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Config(tt.args.key, tt.args.value, tt.args.global); (err != nil) != tt.wantErr {
				t.Errorf("Config() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		{"http.proxy", args{"http.proxy", true}, false},
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

func TestListNameOnly(t *testing.T) {
	names, err := ListNameOnly(true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(names)
}

func TestRemoveSectionIfEmpty(t *testing.T) {
	if err := RemoveSectionIfEmpty("https"); err != nil {
		t.Error(err)
	}
}
