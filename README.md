## Instructions for running

### Docker 
- Install Docker from https://www.docker.com/products/docker-desktop
- Run `docker build .`

### Using your local setup: 
Install go, once done run `go test` in your game folder

### Game of Life
This Kata is about calculating the next generation of Conway's game of life, given any starting position. See http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life for background.
You start with a two dimensional grid of cells, where each cell is either alive or dead. In this version of the problem, the grid is finite, and no life can exist off the edges. When calcuating the next generation of the grid, follow these rules:
1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2. Any live cell with more than three live neighbours dies, as if by overcrowding.
3. Any live cell with two or three live neighbours lives on to the next generation.
4. Any dead cell with exactly three live neighbours becomes a live cell.
   You should write a program that can accept an arbitrary grid of cells, and will output a similar grid showing the next generation.

```
Example input:Grid size: 4,8
   Live cells: [2,5], [3,4], [3,5]
   Expected output:
   Compute and display the next generation on any key press, and exit the program when the subsequent generation doesn't change/evolve or all cells are dead. 
   Generation 1:
   ........
   ....*...
   ...**...
   ........
   Generation 2:
   ........
   ...**...
   ...**...
   ........
   And so on and so forth.
```
