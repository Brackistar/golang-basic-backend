package constants

import "github.com/Brackistar/golang-basic-backend/shared/models"

const (
	CtxKeyPath      models.Key = models.Key("path")
	CtxKeyMethod    models.Key = models.Key("method")
	CtxKeyUser      models.Key = models.Key("user")
	CtxKeyPswd      models.Key = models.Key("pswrd")
	CtxKeyHost      models.Key = models.Key("host")
	CtxKeyDb        models.Key = models.Key("db")
	CtxKeyJwt       models.Key = models.Key("jwt")
	CtxKeyBdy       models.Key = models.Key("body")
	CtxKeyBckt      models.Key = models.Key("bucket")
	CtxKeyDbManager models.Key = models.Key("dbManager")
)

const (
	CtxPassLenght int = 64
)
