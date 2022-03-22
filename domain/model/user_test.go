package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "success",
			args: args{username: "test"},
			want: &User{
				UserName: "test",
			},
			wantErr: false,
		},
		{
			name:    "failure because nil UserName",
			args:    args{username: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username)
			fmt.Print(err != nil, tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Set(t *testing.T) {
	type fields struct {
		ID       int
		UserName string
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{username: "test edited"},
			fields: fields{
				ID:       10,
				UserName: "test",
			},
			wantErr: false,
		},
		{
			name: "failure because nil UserName",
			args: args{username: ""},
			fields: fields{
				ID:       10,
				UserName: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				UserName: tt.fields.UserName,
			}
			if err := u.Set(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("User.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
