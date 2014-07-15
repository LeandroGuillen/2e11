2e11
====

A game theory experiment programmed completely in Go.

Components:
* **engine**: contains the algorithms of the game, movements, score calculations, etc. The game is based on [2048](http://gabrielecirulli.github.io/2048/)
* **aiplayer**: it will play the game and try to optimize the score through a strategy.
* **strategy**: contains the different strategies used to solve problem.
  * Random. Pick one direction and go for it.
