package resolvers

import (
	"context"
	"image/color"
	"time"

	"github.com/afocus/captcha"
	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/storage"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Resolver ...
type Resolver struct {
	Storage storage.Interface
	cache   *cache.Cache
	captcha *captcha.Captcha
}

func NewResolver(db storage.Interface) *Resolver {
	return &Resolver{
		Storage: db,
		cache:   cache.New(15*time.Minute, 30*time.Minute),
		captcha: captchaInit(),
	}
}

func captchaInit() (cca *captcha.Captcha) {
	cca = captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加
	cca.SetFont("./static/comic.ttf")
	// 设置验证码大小
	cca.SetSize(150, 64)
	// 设置干扰强度
	cca.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cca.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cca.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	return
}

// Mutation is the root mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query is the root query resolver
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Healthcheck(ctx context.Context) (string, error) {
	return "ack", nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	return "ack", nil
}

func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	return &timestamppb.Timestamp{
		Seconds: time.Now().Unix(),
	}, nil
}

// field resolver
type fileInfoResolver struct{ *Resolver }
