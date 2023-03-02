package resolvers

import (
	"bytes"
	"context"
	"fmt"
	"image/png"

	"github.com/afocus/captcha"
	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/utils"
	"github.com/patrickmn/go-cache"
)

func (r *queryResolver) Captcha(ctx context.Context) (*generated.Captcha, error) {
	img, str := r.captcha.Create(4, captcha.CLEAR)

	buffer := bytes.NewBuffer([]byte(""))

	err := png.Encode(buffer, img)
	if err != nil {
		return nil, err
	}

	i := buffer.Bytes()
	encode := utils.Base64Encode(i)

	captchaID := utils.RandKey(6)

	r.cache.Set(fmt.Sprintf("%s_captccha", captchaID), str, cache.DefaultExpiration)

	return &generated.Captcha{
		Base64Captcha: encode,
		CaptchaID:     captchaID,
	}, nil
}
