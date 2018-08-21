package context

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestApp_InitializeRoute(t *testing.T) {
	appTest := &App{}

	tests := []struct {
		name string
		a    *App
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.InitializeRoute()
		})
	}
}

func TestApp_InitializeDatabaseTest(t *testing.T) {
	appTest := &App{}
	mock := appTest.InitializeDatabaseTest()
	tests := []struct {
		name string
		a    *App
		want sqlmock.Sqlmock
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
			want: mock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.InitializeDatabaseTest(); got == nil {
				t.Errorf("App.InitializeDatabaseTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_InitializeDatabase(t *testing.T) {
	appTest := &App{}

	tests := []struct {
		name string
		a    *App
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			a:    appTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.InitializeDatabase()
		})
	}
}
