package resolvers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/pkg/errs"
	"github.com/dollarkillerx/eim/internal/utils"
	"github.com/rs/xid"
)

// SendSms ...
func (r *mutationResolver) SendSms(ctx context.Context, input *generated.PhoneInput) (*generated.Sms, error) {
	_, ex := r.cache.Get(input.PhoneNumber)
	if ex {
		return nil, errs.SendSmsPleaseWait
	}

	smsID := xid.New().String()
	code := utils.GenerateRandNum(4)

	r.cache.Set(smsID, code, time.Second*60)
	r.cache.Set(input.PhoneNumber, code, time.Second*60)

	log.Printf("SendSMS %s %s \n", input.PhoneNumber, code)
	return &generated.Sms{
		SmsID: smsID,
	}, nil
}

// CheckSms ...
func (r *queryResolver) CheckSms(ctx context.Context, smsID string, smsCode string) (*generated.CheckSms, error) {
	rc, ex := r.cache.Get(smsID)
	if !ex {
		return &generated.CheckSms{
			Ok: false,
		}, nil
	}
	code := rc.(string)

	fmt.Printf("%s %s %s %v\n", smsID, smsCode, code, smsCode == code)

	if smsCode != code {
		return &generated.CheckSms{
			Ok: false,
		}, nil
	}

	return &generated.CheckSms{
		Ok: true,
	}, nil
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
