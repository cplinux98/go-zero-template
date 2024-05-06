package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	// 用户服务
	message[USER_EXISTS_ERROR] = "用户已存在"
	message[USER_ID_NOT_EXISTS_ERROR] = "用户不存在"
	message[USER_MOBILE_ALREADY_EXISTS_ERROR] = "手机号已存在"
	message[USER_BATCH_DELETE_HAS_NOT_EXISTS_ERROR] = "批量删除的id里面存在无效id"
	message[USER_BATCH_DELETE_AFFECTED_ZERO_ERROR] = "批量删除影响的行数与传入总数不符，已回滚"
	message[USER_BATCH_ADD_HAS_MOBILE_ALREADY_EXISTS_ERROR] = "批量新增时有手机号码已经存在，请检查后重试"
	message[USER_BATCH_ADD_HAS_SAME_MOBILE_ERROR] = "批量新增时有相同的手机号码，请检查后重试"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
