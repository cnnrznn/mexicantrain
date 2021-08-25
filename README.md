# mexicantrain

This tool is a solver for the domino game "mexican train."
The aim of the game is to run out of dominos after drawing a starting set.
One of the most important stages of the game is the initial train build, where the player
is allowed to lay down as many dominos as they wish to build a train.
Given that trains are built by matching the ends of dominos in point value, this tool is a thin wrapper around DFS.

This tool provides the user with two possible trains, both the longest train (most dominos played)
as well as the highest value train (greatest sum of dots).
Scoring works by totalling for each player the amount of dots they have remaining on their dominos, so depending on the
player's strategy they may decide to build a longer train
or build a higher point-value train.

## Usage

`cd cmd && go build && ./cmd`
