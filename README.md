2e11
====

A game theory experiment programmed completely in Go.

Components:
* **aiplayer**: it will play the game and try to optimize the score through a strategy.
* **engine**: contains the algorithms of the game, movements, score calculations, etc. The game is based on [2048](http://gabrielecirulli.github.io/2048/)
* **strategy**: contains the different strategies used to solve problem.
  * Random. Pick one direction and go for it.
  * Recursive. Try the four different directions and pick whichever gets the most points. Only one level of recursion. Work in progress.
