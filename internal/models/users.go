package models

import (
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

type User struct {
	ID              int       `db:"id" goqu:"skipinsert"`
	UserID          string    `db:"user_id"`
	Username        string    `db:"username"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	PlanExpiredDate time.Time `db:"plan_expired_date"`
	DateJoined      time.Time `db:"date_joined" goqu:"skipupdate"`
}

func (obj User) Recipient() string {
	if obj.UserID != "" {
		return obj.UserID
	} else if obj.Username != "" {
		return obj.Username
	}

	return ""
}

func (obj *User) updateUserData(user *tele.Chat) {
	obj.FirstName = user.FirstName
	obj.LastName = user.LastName
	obj.Username = user.Username

	// TODO: update user data on database
}

func (obj *User) GetChat(b *tele.Bot) (*tele.Chat, error) {
	id, err := strconv.Atoi(obj.UserID)
	if err != nil {
		return nil, err
	}

	user, err := b.ChatByID(int64(id))
	if err != nil {
		return nil, err
	}

	obj.updateUserData(user)

	return user, nil
}
