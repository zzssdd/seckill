package cache

import (
	"context"
)

const EmailPreffix = "email:"

type User struct {
}

func EmailTag(email string) string {
	return EmailPreffix + email
}

func (u *User) StoreUserInfo(ctx context.Context, emial, password string, id int64) error {
	data := map[string]interface{}{
		"password": password,
		"id":       id,
	}
	return rds.HMSet(ctx, EmailTag(emial), data).Err()
}

func (u *User) ExistUserInfo(ctx context.Context, email string) bool {
	return rds.Exists(ctx, EmailTag(email)).Val() == 1
}

func (u *User) CheckLogin(ctx context.Context, email string, password string) bool {
	result, err := rds.HGetAll(ctx, EmailTag(email)).Result()
	if err != nil {
		return false
	}
	return result["password"] == password
}
