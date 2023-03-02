package resolvers

import (
	"context"
	"time"

	"github.com/dollarkillerx/eim/internal/conf"
	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/pkg/enum"
	"github.com/dollarkillerx/eim/internal/pkg/errs"
	"github.com/dollarkillerx/eim/internal/utils"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (r *mutationResolver) LoginByPassword(ctx context.Context, input *generated.LoginByPassword) (*generated.AuthPayload, error) {
	captchaOK := utils.CheckImgCaptcha(r.cache, input.CaptchaID, input.CaptchaCode)
	if !captchaOK {
		return nil, errs.CaptchaCode
	}

	account, err := r.Storage.GetUserByAccount(input.Account)
	if err != nil {
		return nil, errs.LoginFailed
	}

	if account.Password != utils.GetPassword(input.Password, conf.CONFIG.Salt) {
		return nil, errs.LoginFailed
	}

	token, err := utils.JWT.CreateToken(enum.AuthJWT{
		generated.UserInformation{
			Account:     account.Account,
			Role:        account.Role,
			AccountID:   account.ID,
			AccountName: account.Name,
		},
	}, time.Now().Add(time.Hour*24*7).Unix())
	if err != nil {
		return nil, errs.SystemError(err)
	}

	return &generated.AuthPayload{
		AccessTokenString: token,
		UserID:            account.ID,
	}, nil
}

func (r *mutationResolver) Registry(ctx context.Context, input *generated.Registry) (*wrapperspb.BoolValue, error) {
	captchaOK := utils.CheckImgCaptcha(r.cache, input.CaptchaID, input.CaptchaCode)
	if !captchaOK {
		return nil, errs.CaptchaCode
	}

	err := r.Storage.AccountRegistry(input.Account, input.Name, utils.GetPassword(input.Password, conf.CONFIG.Salt), generated.RoleGeneralUser)
	if err != nil {
		return nil, errs.SystemError(err)
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (r *queryResolver) User(ctx context.Context) (*generated.UserInformation, error) {
	fromContext, err := utils.GetUserInformationFromContext(ctx)
	if err != nil {
		return nil, errs.PleaseSignIn
	}

	return &generated.UserInformation{
		AccountID:   fromContext.AccountID,
		Role:        fromContext.Role,
		Account:     fromContext.Account,
		AccountName: fromContext.AccountName,
	}, nil
}
