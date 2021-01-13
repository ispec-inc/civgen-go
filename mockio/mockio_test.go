package main

import (
	"reflect"
	"testing"

	"github.com/ispec-inc/civgen-go/mockio/model"
)

func TestGetArgNames(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		method   *model.Method
		expected []string
	}{
		{
			name: "NamedArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "firstArg",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "secondArg",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"firstArg", "secondArg"},
		},
		{
			name: "NotNamedArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"arg0", "arg1"},
		},
		{
			name: "MixedNameArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "firstArg",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "_",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"firstArg", "arg1"},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			g := generator{}

			result := g.getArgNames(testCase.method)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Fatalf("expected %s, got %s", result, testCase.expected)
			}
		})
	}
}

func Test_createPackageMap(t *testing.T) {
	tests := []struct {
		name            string
		importPath      string
		wantPackageName string
		wantOK          bool
	}{
		{"golang package", "context", "context", true},
		{"third party", "golang.org/x/tools/present", "present", true},
		//{"modules", "rsc.io/quote/v3", "quote", true},
		{"fail", "this/should/not/work", "", false},
	}
	var importPaths []string
	for _, t := range tests {
		importPaths = append(importPaths, t.importPath)
	}
	packages := createPackageMap(importPaths)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPackageName, gotOk := packages[tt.importPath]
			if gotPackageName != tt.wantPackageName {
				t.Errorf("createPackageMap() gotPackageName = %v, wantPackageName = %v", gotPackageName, tt.wantPackageName)
			}
			if gotOk != tt.wantOK {
				t.Errorf("createPackageMap() gotOk = %v, wantOK = %v", gotOk, tt.wantOK)
			}
		})
	}
}
