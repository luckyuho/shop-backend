package e

const (
	SUCCESS = 200
	ERROR   = 500

	INVALID_PARAMS = 400
	ERROR_AUTH     = 401
	INVALID_JWT    = 402

	TOOMANYREQ = 411

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 420
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 421
	ERROR_AUTH_CREATE_TOKEN        = 422
	ERROR_ADDRESS_NOTFOUND         = 430
	ERROR_AMOUNT_WITH_PRICELENGTH  = 431
	ERROR_GET_PARAMS               = 432

	ERROR_DB_SEARCH = 534

	NAME_EXISTED          = 601
	COMMODITY_NOT_EXISTED = 602
	DB_INPUT_FAIL         = 603
	DB_OUTPUT_FAIL        = 604
	KYC_NOT_IN_PROGRESS   = 605
	EMAIL_SEND_ERROR      = 606

	ERROR_NOT_ENOUGH_JPOINT           = 620
	ERROR_NFT_ON_SELLING              = 621
	ERROR_NFT_ALREADY_APPLY_WITHDRAWS = 622

	ERROR_2FA_VERIFY       = 623
	ERROR_2FA_OPEN_ALREADY = 624
	ERROR_2FA_NOT_OPEN_YET = 625

	ERROR_QRCODE_NO_ACTIVITY            = 626
	ERROR_QRCODE_ACTIVITY_NO_SERIAL     = 627
	ERROR_QRCODE_ACTIVITY_NOT_START     = 628
	ERROR_QRCODE_ACTIVITY_NOT_AVAILABLE = 629
	ERROR_QRCODE_SERIAL_USED_ALREADY    = 630
	ERROR_QRCODE_SERIAL_NOT_BINDED      = 631
	ERROR_QRCODE_REPEAT_GROUP           = 632

	EMAIL_SEND_SUCCESS = 1001
	UNAUTHORIZED       = 1002
	JWT_EXPIRED        = 1003
	PARAMETER_ERROR    = 1004

	NO_HAVE_DATA = 701
)
