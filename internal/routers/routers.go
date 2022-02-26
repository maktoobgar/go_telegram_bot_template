package routers

import (
	"github.com/maktoobgar/go_telegram_template/internal/handlers"
	tele "gopkg.in/telebot.v3"
)

func Register(b *tele.Bot) {
	b.Handle("/hello", handlers.Hello)
}
