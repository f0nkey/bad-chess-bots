# bad-chess-bots

A bunch of bad chess bots that you can actually beat for once!

* Written in Go
* Compiled to WebAssembly
* Inspired by [this awesome video](https://www.youtube.com/watch?v=DpXy041BIlA)
* Playable in the browser

[Play Now](https://f0nkey.github.io/bad-chess-bots/)

## Bots

| Bot            | Strategy                                                             |
|----------------|:---------------------------------------------------------------------|
| Random         | Plays random moves                                                   | 
| Same Color     | Plays moves to the same color squares as its pieces                  | 
| Opposite Color | Plays moves to the opposite color square                             | 
| Huddle         | Attempts to put its pieces as close as possible to its own king      |
| Swarm          | Attempts to put its pieces as close as possible to the opposing king |

## Building
This project requires [TinyGo](https://github.com/tinygo-org/tinygo) to be installed.

`tinygo build -o ./docs/bots.wasm -target wasm ./`

## Testing
Opening the `index.html` on its own will not work since the proper `Content-Type` for `bots.wasm` will not be served. Use this command from the root to run a server properly serving `bots.wasm`

`go run ./test-server/main.go`
