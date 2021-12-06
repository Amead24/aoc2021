# Day 04

Really lucked out here, my design for part 01 was able to be flipped with two changes to solve part 02.

It feels like i'm really absuing for loops rather than grasping some stronger fundamental of Go.  i'm hoping to get a better understanding of concurrency out of this month, something like  this should be doable?

```
boards := [numDraws, answer] 
for i, input := range inputs {
    board := NewBoard(input)
    // Rather than loop through all the numbers
    // have each board (or function) track those 
    // numbers itself in it's own thread.
    boards = append(checkWinner(board, numbers))
}

map(lowestNumdraws, boards)
```
