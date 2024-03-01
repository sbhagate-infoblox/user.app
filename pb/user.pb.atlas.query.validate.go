// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sbhagate-infoblox/user.app/pb/user.proto

package pb // import "github.com/sbhagate-infoblox/user.app/pb"

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"

// Reference imports to suppress errors if they are not otherwise used.

var UserMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{}
var UserMethodsRequireSortingValidation = map[string][]string{}
var UserMethodsRequireFieldSelectionValidation = map[string][]string{}

func UserValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := UserMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func UserValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := UserMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func UserValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := UserMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}
