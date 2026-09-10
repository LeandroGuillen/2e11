2e11
====

A game theory experiment programmed completely in Go.

Components:
* **aiplayer**: it will play the game and try to optimize the score through a strategy.
* **engine**: contains the algorithms of the game, movements, score calculations, etc. The game is based on [2048](http://gabrielecirulli.github.io/2048/)
* **strategy**: contains the different strategies used to solve problem.
  * Random. Pick one direction and go for it.
  * Recursive. Try the four different directions and pick whichever gets the most points. Only one level of recursion. Work in progress.


## Requirements

* [Go](https://go.dev/dl/) 1.27 or newer. The project is a standard Go module, so no `GOPATH` setup is required.

## Building and running

```sh
go build ./...   # compile everything
go test ./...    # run the engine test suite
go run .         # play a game with the configured strategy
```

## References

- [The Mathematics of 2048: Optimal Play with Markov Decision Processes](http://jdlm.info/articles/2018/03/18/markov-decision-process-2048.html).