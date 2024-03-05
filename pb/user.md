# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [github.com/sbhagate-infoblox/user.app/pb/user.proto](#github-com_sbhagate-infoblox_user-app_pb_user-proto)
    - [CreateUserRequest](#pb-CreateUserRequest)
    - [CreateUserResponse](#pb-CreateUserResponse)
    - [DeleteUserRequest](#pb-DeleteUserRequest)
    - [DeleteUserResponse](#pb-DeleteUserResponse)
    - [DeleteUserSetRequest](#pb-DeleteUserSetRequest)
    - [DeleteUserSetResponse](#pb-DeleteUserSetResponse)
    - [ReadUserRequest](#pb-ReadUserRequest)
    - [ReadUserResponse](#pb-ReadUserResponse)
    - [UpdateUserRequest](#pb-UpdateUserRequest)
    - [UpdateUserResponse](#pb-UpdateUserResponse)
    - [User](#pb-User)
  
    - [Users](#pb-Users)
  
- [Scalar Value Types](#scalar-value-types)



<a name="github-com_sbhagate-infoblox_user-app_pb_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/sbhagate-infoblox/user.app/pb/user.proto



<a name="pb-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [User](#pb-User) |  |  |






<a name="pb-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [User](#pb-User) |  |  |






<a name="pb-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [atlas.rpc.Identifier](#atlas-rpc-Identifier) |  |  |






<a name="pb-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="pb-DeleteUserSetRequest"></a>

### DeleteUserSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [atlas.rpc.Identifier](#atlas-rpc-Identifier) | repeated |  |






<a name="pb-DeleteUserSetResponse"></a>

### DeleteUserSetResponse







<a name="pb-ReadUserRequest"></a>

### ReadUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [atlas.rpc.Identifier](#atlas-rpc-Identifier) |  |  |






<a name="pb-ReadUserResponse"></a>

### ReadUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [User](#pb-User) |  |  |






<a name="pb-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payload | [User](#pb-User) |  |  |






<a name="pb-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [User](#pb-User) |  |  |






<a name="pb-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [atlas.rpc.Identifier](#atlas-rpc-Identifier) |  | string id = 1 [(gorm.field).tag = {type: &#34;uuid&#34; primary_key: true}]; |
| user_name | [string](#string) |  |  |





 

 

 


<a name="pb-Users"></a>

### Users


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateUserRequest](#pb-CreateUserRequest) | [CreateUserResponse](#pb-CreateUserResponse) | Use this method to create a user. |
| Read | [ReadUserRequest](#pb-ReadUserRequest) | [ReadUserResponse](#pb-ReadUserResponse) | Use this method to retrieve the particular user. |
| Update | [UpdateUserRequest](#pb-UpdateUserRequest) | [UpdateUserResponse](#pb-UpdateUserResponse) | Use this method to update user. |
| Delete | [DeleteUserRequest](#pb-DeleteUserRequest) | [DeleteUserResponse](#pb-DeleteUserResponse) | Use this method to delete a user. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

