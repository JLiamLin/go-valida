package valid

const (
	INVALID_PARAMS				= "非法请求参数"
	NOT_REQUIRED				= "&{field}不能为空"
	NOT_ONE_OF					= "&{field}的值只能为[&{param}]"
)

func CommonErrorMessage() map[string]string {
	em := make(map[string]string)

	// 自定义的错误信息
	em["required"] = NOT_REQUIRED
	em["oneof"] = NOT_ONE_OF

	return em
}
