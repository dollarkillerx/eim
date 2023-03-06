package resolvers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/eim/internal/conf"
	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/pkg/enum"
	"github.com/dollarkillerx/eim/internal/pkg/errs"
	"github.com/dollarkillerx/eim/internal/pkg/models"
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
	r.cache.Set(code, input.PhoneNumber, time.Second*60)

	log.Printf("SendSMS %s %s %s \n", input.PhoneNumber, code, smsID)
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

	if smsCode != code {
		return &generated.CheckSms{
			Ok: false,
		}, nil
	}

	r.cache.Set(smsID, code, time.Minute*5)
	rj, _ := r.cache.Get(code)
	r.cache.Set(rj.(string), code, time.Minute*5)
	r.cache.Set(code, rj.(string), time.Minute*5)

	return &generated.CheckSms{
		Ok: true,
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*generated.UserInformation, error) {
	fromContext, err := utils.GetUserInformationFromContext(ctx)
	if err != nil {
		return nil, errs.PleaseSignIn
	}

	var us models.User
	err = r.Storage.DB().Model(&models.User{}).Where("id = ?", fromContext.AccountID).First(&us).Error
	if err != nil {
		return nil, errs.SqlSystemError(err)
	}

	return &generated.UserInformation{
		AccountID: us.ID,
		Account:   us.Account,
		FullName:  us.FullName,
		NickName:  us.Nickname,
		Birthday:  us.Birthday,
		Email:     us.Email,
		About:     us.About,
		Avatar:    fmt.Sprintf("%s/%s", conf.CONFIG.AssetUri, us.Avatar),
	}, nil
}

// UserRegistration ...
func (r *mutationResolver) UserRegistration(ctx context.Context, input *generated.UserRegistration) (*generated.AuthPayload, error) {
	// check sms
	rsms, ex := r.cache.Get(input.SmsID)
	if !ex {
		return nil, errs.CaptchaCode
	}
	rphome, ex := r.cache.Get(input.SmsCode)
	if !ex {
		return nil, errs.CaptchaCode
	}
	code := rsms.(string)
	phoneNumber := rphome.(string)

	if input.SmsCode != code {
		return nil, errs.CaptchaCode
	}

	// 查询已有用户是否存在
	var exuser int64
	err := r.Storage.DB().Model(&models.User{}).Where("account = ?", phoneNumber).Count(&exuser).Error
	if err != nil {
		log.Println(err)
		return nil, errs.SqlSystemError(err)
	}

	if exuser != 0 {
		log.Println(err)
		return nil, errs.ExUser
	}

	// 写
	uid := xid.New().String()
	err = r.Storage.DB().Model(&models.User{}).Create(&models.User{
		BasicModel: models.BasicModel{
			ID: uid,
		},
		Account:  phoneNumber,
		FullName: input.FullName,
		Nickname: input.NickName,
		Birthday: input.Birthday,
		Email:    input.Email,
		About:    input.About,
		Avatar:   input.Avatar,
	}).Error
	if err != nil {
		log.Println(err)
		return nil, errs.SqlSystemError(err)
	}

	token, err := utils.JWT.CreateToken(&enum.AuthJWT{
		generated.UserInformation{
			AccountID: uid,
			Account:   phoneNumber,
			FullName:  input.FullName,
			NickName:  input.NickName,
			Birthday:  input.Birthday,
			Email:     input.Email,
			About:     input.About,
			Avatar:    input.Avatar,
		},
	}, 0)
	if err != nil {
		log.Println(err)
		return nil, errs.SystemError(err)
	}

	return &generated.AuthPayload{
		AccessTokenString: token,
		UserID:            uid,
	}, nil
}

// UserLogin ...
func (r *queryResolver) UserLogin(ctx context.Context, smsID string, smsCode string) (*generated.AuthPayload, error) {
	// check sms
	rsms, ex := r.cache.Get(smsID)
	if !ex {
		return nil, errs.CaptchaCode
	}
	rphome, ex := r.cache.Get(smsCode)
	if !ex {
		return nil, errs.CaptchaCode
	}
	code := rsms.(string)
	phoneNumber := rphome.(string)

	if smsCode != code {
		return nil, errs.CaptchaCode
	}

	// 查询已有用户
	var user models.User
	err := r.Storage.DB().Model(&models.User{}).Where("account = ?", phoneNumber).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, errs.SqlSystemError(err)
	}

	token, err := utils.JWT.CreateToken(&enum.AuthJWT{
		generated.UserInformation{
			AccountID: user.ID,
			Account:   phoneNumber,
			FullName:  user.FullName,
			NickName:  user.Nickname,
			Birthday:  user.Birthday,
			Email:     user.Email,
			About:     user.About,
			Avatar:    user.Avatar,
		},
	}, 0)
	if err != nil {
		log.Println(err)
		return nil, errs.SystemError(err)
	}

	return &generated.AuthPayload{
		AccessTokenString: token,
		UserID:            user.ID,
	}, nil
}

// Friendship ...
func (r *queryResolver) Friendship(ctx context.Context) (*generated.Friendships, error) {
	fromContext, err := utils.GetUserInformationFromContext(ctx)
	if err != nil {
		return nil, errs.PleaseSignIn
	}

	var result generated.Friendships

	sql := `SELECT users.*
			FROM friendships
			JOIN users ON users.id = friendships.user1_id  OR users.id = friendships.user2_id 
			WHERE (friendships.user1_id = ? OR friendships.user2_id = ? )
			and users.deleted_at is null 
			and friendships.deleted_at is null group by users.id;
	`

	var users []models.User
	err = r.Storage.DB().Raw(sql, fromContext.AccountID, fromContext.AccountID).Scan(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, v := range users {
		result.Friendships = append(result.Friendships, generated.Friendship{
			AccountID: v.ID,
			Account:   v.Account,
			FullName:  v.FullName,
			NickName:  v.Nickname,
			Birthday:  v.Birthday,
			Email:     v.Email,
			About:     v.About,
			Avatar:    fmt.Sprintf("%s/%s", conf.CONFIG.AssetUri, v.Avatar),
		})
	}

	return &result, nil
}
