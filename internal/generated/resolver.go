package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Resolver struct{}

// // foo
func (r *mutationResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UploadFile(ctx context.Context, file graphql.Upload) (string, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) SendSms(ctx context.Context, input *PhoneInput) (*Sms, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) User(ctx context.Context) (*UserInformation, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) CheckSms(ctx context.Context, smsID string, smsCode string) (*CheckSms, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Captcha(ctx context.Context) (*Captcha, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
