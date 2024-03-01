package pb

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	fmt "fmt"
	runtime1 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	runtime "github.com/infobloxopen/protoc-gen-atlas-validate/runtime"
	metadata "google.golang.org/grpc/metadata"
	ioutil "io/ioutil"
	http "net/http"
)

// validate_Object_User function validates a JSON for a given object.
func validate_Object_User(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&User{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_User(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "id":
		case "user_name":
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object User.
func (_ *User) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&User{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_User(ctx, r, path)
}

func validate_required_Object_User(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_CreateUserRequest function validates a JSON for a given object.
func validate_Object_CreateUserRequest(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&CreateUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_CreateUserRequest(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "payload":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := runtime.JoinPath(path, k)
			if err = validate_Object_User(ctx, vv, vvPath); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object CreateUserRequest.
func (_ *CreateUserRequest) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&CreateUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_CreateUserRequest(ctx, r, path)
}

func validate_required_Object_CreateUserRequest(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_CreateUserResponse function validates a JSON for a given object.
func validate_Object_CreateUserResponse(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&CreateUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_CreateUserResponse(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := runtime.JoinPath(path, k)
			if err = validate_Object_User(ctx, vv, vvPath); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object CreateUserResponse.
func (_ *CreateUserResponse) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&CreateUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_CreateUserResponse(ctx, r, path)
}

func validate_required_Object_CreateUserResponse(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_ReadUserRequest function validates a JSON for a given object.
func validate_Object_ReadUserRequest(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&ReadUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_ReadUserRequest(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "id":
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ReadUserRequest.
func (_ *ReadUserRequest) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&ReadUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_ReadUserRequest(ctx, r, path)
}

func validate_required_Object_ReadUserRequest(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_ReadUserResponse function validates a JSON for a given object.
func validate_Object_ReadUserResponse(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&ReadUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_ReadUserResponse(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := runtime.JoinPath(path, k)
			if err = validate_Object_User(ctx, vv, vvPath); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object ReadUserResponse.
func (_ *ReadUserResponse) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&ReadUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_ReadUserResponse(ctx, r, path)
}

func validate_required_Object_ReadUserResponse(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_UpdateUserRequest function validates a JSON for a given object.
func validate_Object_UpdateUserRequest(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&UpdateUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_UpdateUserRequest(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "payload":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := runtime.JoinPath(path, k)
			if err = validate_Object_User(ctx, vv, vvPath); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object UpdateUserRequest.
func (_ *UpdateUserRequest) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&UpdateUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_UpdateUserRequest(ctx, r, path)
}

func validate_required_Object_UpdateUserRequest(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_UpdateUserResponse function validates a JSON for a given object.
func validate_Object_UpdateUserResponse(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&UpdateUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_UpdateUserResponse(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "result":
			if v[k] == nil {
				continue
			}
			vv := v[k]
			vvPath := runtime.JoinPath(path, k)
			if err = validate_Object_User(ctx, vv, vvPath); err != nil {
				return err
			}
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object UpdateUserResponse.
func (_ *UpdateUserResponse) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&UpdateUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_UpdateUserResponse(ctx, r, path)
}

func validate_required_Object_UpdateUserResponse(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_DeleteUserRequest function validates a JSON for a given object.
func validate_Object_DeleteUserRequest(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_DeleteUserRequest(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "id":
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteUserRequest.
func (_ *DeleteUserRequest) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_DeleteUserRequest(ctx, r, path)
}

func validate_required_Object_DeleteUserRequest(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_DeleteUserResponse function validates a JSON for a given object.
func validate_Object_DeleteUserResponse(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_DeleteUserResponse(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteUserResponse.
func (_ *DeleteUserResponse) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_DeleteUserResponse(ctx, r, path)
}

func validate_required_Object_DeleteUserResponse(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_DeleteUserSetRequest function validates a JSON for a given object.
func validate_Object_DeleteUserSetRequest(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserSetRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_DeleteUserSetRequest(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		case "ids":
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteUserSetRequest.
func (_ *DeleteUserSetRequest) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserSetRequest{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_DeleteUserSetRequest(ctx, r, path)
}

func validate_required_Object_DeleteUserSetRequest(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

// validate_Object_DeleteUserSetResponse function validates a JSON for a given object.
func validate_Object_DeleteUserSetResponse(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserSetResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}

	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("invalid value for %q: expected object.", path)
	}

	if err = validate_required_Object_DeleteUserSetResponse(ctx, v, path); err != nil {
		return err
	}

	allowUnknown := runtime.AllowUnknownFromContext(ctx)

	for k, _ := range v {
		switch k {
		default:
			if !allowUnknown {
				return fmt.Errorf("unknown field %q.", runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object DeleteUserSetResponse.
func (_ *DeleteUserSetResponse) AtlasValidateJSON(ctx context.Context, r json.RawMessage, path string) (err error) {
	if hook, ok := interface{}(&DeleteUserSetResponse{}).(interface {
		AtlasJSONValidate(context.Context, json.RawMessage, string) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(ctx, r, path); err != nil {
			return err
		}
	}
	return validate_Object_DeleteUserSetResponse(ctx, r, path)
}

func validate_required_Object_DeleteUserSetResponse(ctx context.Context, v map[string]json.RawMessage, path string) error {
	method := runtime.HTTPMethodFromContext(ctx)
	_ = method
	return nil
}

var validate_Patterns = []struct {
	pattern    runtime1.Pattern
	httpMethod string
	validator  func(context.Context, json.RawMessage) error
	// Included for introspection purpose.
	allowUnknown bool
}{
	// patterns for file github.com/infobloxopen/protoc-gen-gorm/proto/options/gorm.proto

	// patterns for file github.com/sbhagate-infoblox/user.app/pb/user.proto

	// patterns for file google/protobuf/descriptor.proto

}

// AtlasValidateAnnotator parses JSON input and validates unknown fields
// based on 'allow_unknown_fields' options specified in proto file.
func AtlasValidateAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)
	for _, v := range validate_Patterns {
		if r.Method == v.httpMethod && runtime.PatternMatch(v.pattern, r.URL.Path) {
			var b []byte
			var err error
			if b, err = ioutil.ReadAll(r.Body); err != nil {
				md.Set("Atlas-Validation-Error", "invalid value: unable to parse body")
				return md
			}
			r.Body = ioutil.NopCloser(bytes.NewReader(b))
			ctx := context.WithValue(context.WithValue(context.Background(), runtime.HTTPMethodContextKey, r.Method), runtime.AllowUnknownContextKey, v.allowUnknown)
			if err = v.validator(ctx, b); err != nil {
				md.Set("Atlas-Validation-Error", err.Error())
			}
			break
		}
	}
	return md
}
