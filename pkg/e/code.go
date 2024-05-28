package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user模块
	ErrorNoPassword        = 30001
	ErrorFailEncryption    = 30002
	ErrorExistUserNotFound = 30003
	ErrorNotCompare        = 30004
	ErrorAuthToken         = 30005
	ErrorAuthTokenTimeout  = 30006
	ErrorUploadFail        = 30007
	ErrorSendEmail         = 30008

	//movie 模块
	ErrorStatus           = 40001
	ErrorProductImgUpload = 40002
	ErrorMovieIndex       = 40003
)
