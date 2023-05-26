package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "參數錯誤",
	INVALID_JWT:    "Token未帶入",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token驗證失敗",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超時",
	ERROR_AUTH_CREATE_TOKEN:        "Token生成失敗",
	ERROR_AUTH:                     "Token錯誤",
	ERROR_ADDRESS_NOTFOUND:         "該地址不存在",
	ERROR_GET_PARAMS:               "取得參數失敗",

	NAME_EXISTED:        "名稱已存在",
	DB_INPUT_FAIL:       "資料寫入失敗",
	DB_OUTPUT_FAIL:      "資料讀取失敗",
	NO_HAVE_DATA:        "查無資料",
	KYC_NOT_IN_PROGRESS: "此會員尚未提交KYC資料",
	EMAIL_SEND_ERROR:    "email 發送失敗",

	ERROR_NOT_ENOUGH_JPOINT:           "JPoint不足，請先儲值．",
	ERROR_NFT_ON_SELLING:              "商品於NFTgogo掛單中，請先取消掛單。",
	ERROR_NFT_ALREADY_APPLY_WITHDRAWS: "商品已申請出金，正在審核中．",

	ERROR_2FA_VERIFY:       "驗證碼輸入錯誤",
	ERROR_2FA_OPEN_ALREADY: "已開通過2FA",
	ERROR_2FA_NOT_OPEN_YET: "未開通2FA",

	ERROR_QRCODE_NO_ACTIVITY:            "沒有該活動",
	ERROR_QRCODE_ACTIVITY_NO_SERIAL:     "該活動沒有此序號",
	ERROR_QRCODE_ACTIVITY_NOT_START:     "活動尚未開始",
	ERROR_QRCODE_ACTIVITY_NOT_AVAILABLE: "活動已結束",
	ERROR_QRCODE_SERIAL_USED_ALREADY:    "序號已使用",
	ERROR_QRCODE_SERIAL_NOT_BINDED:      "序號未註冊",
	ERROR_QRCODE_REPEAT_GROUP:           "有大於一個相同的序號在同組活動群組，請通知 Jcard 調整該序號",

	EMAIL_SEND_SUCCESS: "email 發送成功",
	UNAUTHORIZED:       "驗證失敗",
	JWT_EXPIRED:        "token 已過期",

	TOOMANYREQ: "請求次數過多",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
