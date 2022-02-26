package app

import (
	"fmt"
	"log"
	"time"

	_ "github.com/maktoobgar/go_telegram_template/internal/app/load"
	g "github.com/maktoobgar/go_telegram_template/internal/global"
	"github.com/maktoobgar/go_telegram_template/internal/models"
	"github.com/maktoobgar/go_telegram_template/internal/routers"
	tele "gopkg.in/telebot.v3"
)

func HelloDeveloper(b *tele.Bot) {
	user := &models.User{UserID: g.CFG.DeveloperID}
	chat, err := user.GetChat(b)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = b.Send(chat, fmt.Sprintf("hello %s!", user.Username))
	if err != nil {
		log.Fatalln(err)
	}
}

func Run() {
	// Setup Bot
	pref := tele.Settings{
		Token:  g.CFG.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Global Access To Bot
	g.App = b

	// Register Routers
	routers.Register(b)

	// Say Hi To Developer
	HelloDeveloper(b)

	// Start Bot
	b.Start()
}
