package component

import (
	"gorm.io/gorm"
	"web/components/config"
	"web/components/mycache"
	"web/components/tokenprovider"
	"web/components/uploadprovider"
	"web/pubsub"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
	GetMyCache() mycache.Cache
	GetTokenProvider() tokenprovider.TokenProvider
	GetUploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	appConfig      *config.AppConfig
	database       *gorm.DB
	myCache        mycache.Cache
	tokenProvider  tokenprovider.TokenProvider
	uploadProvider uploadprovider.UploadProvider
	pubSub         pubsub.PubSub
}

func NewAppContext(appConfig *config.AppConfig, database *gorm.DB, myCache mycache.Cache, tokenProvider tokenprovider.TokenProvider, uploadProvider uploadprovider.UploadProvider, pubSub pubsub.PubSub) *appCtx {
	return &appCtx{
		appConfig:      appConfig,
		database:       database,
		myCache:        myCache,
		tokenProvider:  tokenProvider,
		uploadProvider: uploadProvider,
		pubSub:         pubSub,
	}
}

func (ctx *appCtx) GetAppConfig() *config.AppConfig {
	return ctx.appConfig
}

func (ctx *appCtx) GetDatabase() *gorm.DB {
	return ctx.database
}

func (ctx *appCtx) GetMyCache() mycache.Cache {
	return ctx.myCache
}

func (ctx *appCtx) GetTokenProvider() tokenprovider.TokenProvider {
	return ctx.tokenProvider
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.pubSub
}
