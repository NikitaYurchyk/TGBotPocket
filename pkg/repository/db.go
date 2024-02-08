package repository

type Bucket string

const (
	AccessTokens  Bucket = "access_token"
	RequestTokens Bucket = "request_token"
)

type TokenRepository interface {
	Save(chatId int64, token string, bucket Bucket) error
	Get(chatId int64, bucket Bucket) (string, error)
}
