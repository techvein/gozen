package json

type ErrorJSON struct {
	// TODO: エラーレスポンス要検討
	Code    int    `json:"-"`
	Message string `json:"message"`
}
