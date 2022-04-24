package main

import (
	"bad-chess-bots/bot"
	"fmt"
	"github.com/notnil/chess"
	"log"
	"syscall/js"
)

func main() {
	js.Global().Set("oppositeColor", js.FuncOf(jsCompatiableFunc(bot.OppositeColor)))
	js.Global().Set("sameColor", js.FuncOf(jsCompatiableFunc(bot.SameColor)))
	js.Global().Set("random", js.FuncOf(jsCompatiableFunc(bot.Random)))
	<-make(chan bool)
}

// jsCompatiableFunc decorates a bot move function such that it is usable through Javascript.
func jsCompatiableFunc(moveFunc func(*chess.Position) *chess.Move) func(this js.Value, args []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		fen := args[0].String()
		gameFunc, err := chess.FEN(args[0].String())
		if err != nil {
			log.Fatal("parsing fen", fen, err.Error())
		}
		game := chess.NewGame(gameFunc)

		m := moveFunc(game.Position())
		moveStr := fmt.Sprintf("%s,%s", m.S1().String(), m.S2().String())
		return js.ValueOf(moveStr)
	}
}
