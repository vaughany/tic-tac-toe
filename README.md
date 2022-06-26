# Tic-Tac-Toe in Go 

(Tic-Tac-Go?)

```
Welcome to:
 T | i | c
---+---+---
 T | a | c
---+---+---
 T | o | e
```

The standard tic-tac-toe game written in Go, and attempting to be a well-written piece of code.

---

## Installtion

1. Have Go installed ([go.dev/doc/install](https://go.dev/doc/install))
2. Clone this repository ("`git clone git@github.com:vaughany/tic-tac-toe.git`")

---

## Use

`cd` to the repository and run with `go run .`

The game is intended to be user-friendly. You use the numbers 1 to 9 to choose a cell in the grid (an example is shown at startup),
or `q` to quit.

```
When choosing a cell, use the following numbers:
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
```

Player 1 plays first, and play alternates between players thereafter.

Each player's wins are totalled, and shown at the end of a game.

---

## Customising

By default, the first player is called `Player 1` and is assigned the mark `o`; the second player is called `Player 2` and is assigned the mark `x`.  You can customise these settings on the command line.

If You wanted Alice and Bob to play with the marks `O` and `X`, you would run the game as follows:

```
$ go run . -p1name Alice -p2name Bob -p1mark O -p2mark X
```

You should be able to use emoji without issue, but as of right now it'll break the board, e.g.:

```
$ go run . -p1name Alice -p2name Bob -p1mark ✅ -p2mark ❌
```

**Note:** If you specify identical names or marks for both players, the default will be used instead. 

---

## To Do

* One player is the computer: need to figure out how to respond to user choices and how to detect winnable situations.
* Persist player data to disk: load it at startup, save it at shutdown.

---

## History

* **v2.1.0**: Added a README.
* **v2.0.0**: Better game. Uses player objects; checks for wins; displays scoreboard.
* **v1.0.0**: Basic game, puts marks in boxes.