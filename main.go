package main

import (
	"log"

	"time"

	tele "gopkg.in/telebot.v4"
)

const (
	HelloMessage = `
Мне кажется, я никогда не смогу в достаточной мере выразить то, что я чувствую к тебе,
но вот малая часть того...
`
	Cmsg = `Не останавливайся, у меня есть еще много того, что я хочу тебе сказать
Продолжай, зайка`
	Fmsg = `
	Это лишь малая часть того, что я на самом деле чувствую к тебе, я просто без ума от тебя...
Иди я буду тебя целомкать и обнимать, зайка)
`
	Msg1 = `
Хочешь знать что делает меня счастиливым каждый день?

Ты даже и представить себе не можешь, что я чувствую в такие моменты...

Ответ: https://www.instagram.com/share/reel/BAVVYbFhe4
`
	Msg2 = `
Ты любишь меня спрашивать это и я хочу что бы у тебя был ответ.

Ты все равно будешь спрашивать коненчо, но я буду повторять это столько, сколько потербуется...

Ответ: https://www.instagram.com/share/reel/_gwY99KQi
`

	Msg3 = `
Люблю ли я тебя ?

Ответ: https://www.instagram.com/share/reel/_eoT6S7Pv
`
	Msg4 = `
Не хочу думать о том, что было бы если бы я тебя не встретил.

Но если тебе интересно, то...

Ответ: https://www.instagram.com/share/reel/BAKlMwr4ew
`
	Msg5 = `
Если ты хочешь узнать что я чувствую, когда я рядом с тобой то вот.

Ответ: https://www.instagram.com/share/reel/BAFVRpVfrq
`
	Msg6 = `
А знаешь о чем ты даже не догадываешься м?

Ответ: https://www.instagram.com/share/reel/BADZaxA5pR
`
	Msg7 = `
Я уверен что ты не знаешь, что есть что то, что я люблю больше чем тебя

Больше чем тебя я люблю только...

Ответ: https://www.instagram.com/share/reel/BANGuB8U_L
`
	Msg8 = `
Хочешь знать, что я чувствую каждое утро с тобой?

Почувствуй

Ответ: https://www.instagram.com/share/reel/BAIzAv04uk

`
	Msg9 = `
Идеальные не идеальны

Но занешь что?

Ответ: https://www.instagram.com/share/reel/BADgsoc1cE

`
	Msg10 = `
А заешь, что я знаю на счет 2025 года?

Знаешь?

Ответ: https://www.instagram.com/share/reel/_lfA275hE

`
	Msg11 = `
Знаешь что делает мой день лучше?

Ответ: https://www.instagram.com/share/reel/_kgBsw3lD
`
)

func main() {
	pref := tele.Settings{
		Token:  "8148615418:AAEjFxv0BpEzqLdCTvqGv06ZsKiPCeVLbCc",
		Poller: &tele.LongPoller{Timeout: 9 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var count int

	var selector *tele.ReplyMarkup
	var selector2 *tele.ReplyMarkup

	cbtn := selector.Data("Нажми и продолжим, зайка", "c")
	btn1 := selector.Data("Нажми и узнаешь", "1")
	btn2 := selector.Data("Нажми и узнаешь", "2")
	btn3 := selector.Data("Нажми и узнаешь", "3")
	btn4 := selector.Data("Нажми и узнаешь", "4")
	btn5 := selector.Data("Нажми и узнаешь", "5")
	btn6 := selector.Data("Нажми и узнаешь", "6")
	btn7 := selector.Data("Нажми и узнаешь", "7")
	btn8 := selector.Data("Нажми и узнаешь", "8")
	btn9 := selector.Data("Нажми и узнаешь", "9")
	btn10 := selector.Data("Нажми и узнаешь", "10")
	btn11 := selector.Data("Что новенькое для тебя", "11")

	selector3 := &tele.ReplyMarkup{}
	selector3.Inline(
		selector.Row(btn11),
	)

	b.Handle("/start", func(c tele.Context) error {
		count = 0
		selector = &tele.ReplyMarkup{}
		selector2 = &tele.ReplyMarkup{}
		selector.Inline(
			selector.Row(btn1),
			selector.Row(btn2),
			selector.Row(btn3),
			selector.Row(btn4),
			selector.Row(btn5),
			selector.Row(btn6),
			selector.Row(btn7),
			selector.Row(btn8),
			selector.Row(btn9),
			selector.Row(btn10),
		)
		selector2.Inline(selector2.Row(cbtn))
		return c.Send(HelloMessage, selector)
	})

	b.Handle(&btn1, func(c tele.Context) error {
		b := selector.InlineKeyboard[0][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[0][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg1, selector2)
	})
	b.Handle(&btn2, func(c tele.Context) error {
		b := selector.InlineKeyboard[1][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[1][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg2, selector2)
	})

	b.Handle(&btn3, func(c tele.Context) error {
		b := selector.InlineKeyboard[2][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[2][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg3, selector2)
	})

	b.Handle(&btn4, func(c tele.Context) error {
		b := selector.InlineKeyboard[3][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[3][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg4, selector2)
	})

	b.Handle(&btn5, func(c tele.Context) error {
		b := selector.InlineKeyboard[4][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[4][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg5, selector2)
	})

	b.Handle(&btn6, func(c tele.Context) error {
		b := selector.InlineKeyboard[5][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[5][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg6, selector2)
	})

	b.Handle(&btn7, func(c tele.Context) error {
		b := selector.InlineKeyboard[6][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[6][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg7, selector2)
	})

	b.Handle(&btn8, func(c tele.Context) error {
		b := selector.InlineKeyboard[7][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[7][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg8, selector2)
	})

	b.Handle(&btn9, func(c tele.Context) error {
		b := selector.InlineKeyboard[8][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[8][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg9, selector2)
	})

	b.Handle(&btn10, func(c tele.Context) error {
		b := selector.InlineKeyboard[9][0]
		if b.Text != "Это ты уже знаешь, зайка" {
			count += 1
		}
		b.Text = "Это ты уже знаешь, зайка"
		selector.InlineKeyboard[9][0] = b
		if count == 10 {
			b2 := selector2.InlineKeyboard[0][0]
			b2.Text = "Осталось нажать еще один раз..."
			selector2.InlineKeyboard[0][0] = b2
		}
		return c.Send(Msg10, selector2)
	})

	b.Handle("/new", func(c tele.Context) error {
		return c.Send("Что то, что я еще хочу сказать тебе, Зайка", selector3)
	})

	b.Handle(&btn11, func(c tele.Context) error {
		b := selector3.InlineKeyboard[0][0]
		b.Text = "Уже не новенькое"
		selector3.InlineKeyboard[0][0] = b
		return c.Send(Msg11)
	})

	b.Handle(&cbtn, func(c tele.Context) error {
		if count == 10 {
			return c.Send(Fmsg)
		}
		return c.Send(Cmsg, selector)

	})
	b.Start()
}
