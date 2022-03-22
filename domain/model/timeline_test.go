package model

import (
	"reflect"
	"testing"
)

func TestNewTimeline(t *testing.T) {
	type args struct {
		user_id int
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *Timeline
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeline(tt.args.user_id, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimeline() = %v, want %v", got, tt.want)
			}
		})
	}
}
