
Tic Tac Toe
=

Commands
-


These commands were used to generate the scaffolding for this game:

*NOTE*: The original repository was created from a command which used `checkers` module name as an example, hence this repository and commands mention checkers (instead of tictactoe).

This work doesn't have any test automation, which will be necessary for any production code. Rules of the game were extracted in `rules.go` and tested manually using main method in the same file. 


Storage/Schema

```
ignite scaffold single nextGame creator idValue:uint --module checkers --no-message
ignite scaffold map waitingGame creator idValue:uint --module checkers --no-message
ignite scaffold map storedGame creator idValue:uint game crossPlayer circlePlayer --module checkers --no-message
```

Messages

```
ignite scaffold message createGame --module checkers --response idValue
ignite scaffold message joinGame idValue:uint --module checkers --response success:bool
ignite scaffold message markSpace idValue:uint x:uint y:uint --module checkers --response success:bool
```


How to Interact with the game
=

You can use the following CLI instructions to interact with the game

I had to declare the following alias as the `checkersd` command was not added to my environment already. 

`alias tictactoe='go run cmd/checkersd/main.go'`


Export these pre-created users's addresses in environment, so that they can be used in the game commands:

```
export alice=$(tictactoe keys show alice -a)
export bob=$(tictactoe keys show bob -a)
```

Explore the commands, and state in stores:

```
tictactoe status | jq
tictactoe query checkers --help 
tictactoe query checkers list-stored-game 
tictactoe query checkers list-waiting-game 
```


Play game

```
tictactoe tx checkers create-game --from $alice --gas auto
tictactoe tx checkers join-game 1 --from $bob --gas auto
tictactoe tx checkers mark-space 1 2 2 --from $bob --gas auto
tictactoe tx checkers mark-space 1 2 2 --from $alice --gas auto
```

