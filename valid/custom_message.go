package valid

const (
	NotRequired				= ""
)

func CommonErrorMessage() map[string]string {
	em := make(map[string]string)

	// 自定义的错误信息
	em["required"] = NotRequired

	return em
}
