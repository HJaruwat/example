package errors

import "cabal-api/common"

// ErrorBadRequest is bad request
var ErrorBadRequest = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "server cannot or will not process this request",
}

// ErrorInternalServer is internal_server_error
var ErrorInternalServer = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "some thing wrong",
}

// ErrorMigrate is internal_server_error
var ErrorMigrate = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "migrate have a problem",
}

// ErrorPutMessage is internal_server_error
var ErrorPutMessage = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "put message to queue error",
}

// ErrorCodeNotFound is internal_server_error
var ErrorCodeNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "code not found",
}

// ErrorServerIDNotFound is internal_server_error
var ErrorServerIDNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "server id not found",
}

// ErrorGetReward is
var ErrorGetReward = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "get reward error",
}

// ErrorInsertMember is
var ErrorInsertMember = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "fail to create partner member",
}

// ErrorMemberIsExist is
var ErrorMemberIsExist = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "target account is exist",
}

// ErrorTargetAccountNotFound is
var ErrorTargetAccountNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "target account not found",
}

// ErrorPartnerNotFound is
var ErrorPartnerNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "partner_id not found",
}

// ErrorUsernameNotFound is
var ErrorUsernameNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "username not found",
}

// ErrorReferenceNotFound is
var ErrorReferenceNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "refcode not found",
}

// ErrorReferenceCodeQueryNotFound is
var ErrorReferenceCodeQueryNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "refcode query not found",
}

// ErrorCreateTransaction is
var ErrorCreateTransaction = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error create transaction",
}

// ErrorUpdateTransaction is
var ErrorUpdateTransaction = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error update transaction",
}

// ErrorConvertRate is
var ErrorConvertRate = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error get convertion rate",
}

// ErrorCondition is
var ErrorCondition = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "target account invalid some condition",
}

// ErrorUsername is error username not found
var ErrorUsername = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "username not found",
}

// ErrorPasswordIsRequired is error username not found
var ErrorPasswordIsRequired = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "password field is required",
}

// ErrorLoginPassPort is error username not found
var ErrorLoginPassPort = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "login passport error",
}

// ErrorUIDNotFound is error username not found
var ErrorUIDNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "uid not found",
}

// ErrorAmount is error username not found
var ErrorAmount = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "amount not found",
}

// ErrorRewardsList is error username not found
var ErrorRewardsList = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "rewards not found",
}

// ErrorGetTopupWithReferenceCode is error refid not found
var ErrorGetTopupWithReferenceCode = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "error get topup with referencecode",
}

// ErrorContentType is error can not support content-type
var ErrorContentType = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "this service can not support this content-type",
}

// ErrorBodyDecode is error can not support content-type
var ErrorBodyDecode = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "fail to decodeing parameters",
}

// ErrorBodyDecode is error can not support content-type
var ErrorBodyEncode = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "fail to encodeing parameters",
}

// ErrorServerID is error server_id not found
var ErrorServerID = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "server_id not found",
}

// ErrorEstID is error est_id not found
var ErrorEstID = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "est_id not found",
}

// ErrorJSONEncode is error est_id not found
var ErrorJSONEncode = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error json encode",
}

// ErrorGameUserNotFound is error est_id not found
var ErrorGameUserNotFound = &common.ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "this user do not have account in game",
}

// ErrorCannotConnectGameAPI is error est_id not found
var ErrorCannotConnectGameAPI = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error can not connection game api",
}

// ErrorGetAuthKeyWithGameAPI is error est_id not found
var ErrorGetAuthKeyWithGameAPI = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error get authkey with gameapi",
}

// ErrorResetSubPasswordWithGameAPI is error est_id not found
var ErrorResetSubPasswordWithGameAPI = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "error reset sub-password with gameapi",
}

// ErrorCharacterInfo is Error Est api
var ErrorCharacterInfo = &common.ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "Can't Get character Info.",
}

// ErrorData is
func ErrorData(err string, desc string, data interface{}) *common.ErrorResponse {
	return &common.ErrorResponse{
		Error:            err,
		ErrorDescription: desc,
		ErrorData:        data,
	}
}
