## Day 4

#### Part 1
Ouf, part 1 was not fun.
Stuff like making a sorting function or reversing is nice to haves, but a lot of the mangling is not any easier. Perhaps I become better at it with time.

#### Part 2
Hmm, gluing stuff together was impossible, so I had to cheat by adding the line:
```
result, resultInts = getCorrectRooms(input)
```
then after doing it, it was able to generate it by itself.

The rest was fairly easy, it generated a mostly correct shift function by itself. Also realised today that the `sort` function from day 3 doesn't actually sort, but works because the length of the array is 3 every time.

To run:
```
go run .
```
First number is part1, last number is part 2.
