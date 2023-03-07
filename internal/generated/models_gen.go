// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
)

type AuthPayload struct {
	AccessTokenString string `json:"accessTokenString"`
	UserID            string `json:"userID"`
}

type Captcha struct {
	Base64Captcha string `json:"base64Captcha"`
	CaptchaID     string `json:"captchaId"`
}

type CheckSms struct {
	Ok bool `json:"ok"`
}

type Friendships struct {
	Friendships []UserInformation `json:"friendships"`
}

type PhoneInput struct {
	PhoneNumber string `json:"phoneNumber"`
}

type Result struct {
	Ok bool `json:"ok"`
}

type Sms struct {
	SmsID string `json:"smsId"`
}

type SearchUser struct {
	Users []UserInformation `json:"users"`
}

type UserInformation struct {
	AccountID string `json:"accountId"`
	Account   string `json:"account"`
	FullName  string `json:"fullName"`
	NickName  string `json:"nickName"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
	About     string `json:"about"`
	Avatar    string `json:"avatar"`
}

type UserRegistration struct {
	FullName string `json:"fullName"`
	NickName string `json:"nickName"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	About    string `json:"about"`
	Avatar   string `json:"avatar"`
	SmsID    string `json:"smsId"`
	SmsCode  string `json:"smsCode"`
}

type Role string

const (
	RoleGeneralUser Role = "GeneralUser"
	RoleAdmin       Role = "Admin"
)

var AllRole = []Role{
	RoleGeneralUser,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleGeneralUser, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
