package main

import "fmt"

/**
 * The while loop represents the game.
 * Each iteration represents a turn of the game
 * where you are given inputs (the heights of the mountains)
 * and where you have to print an output (the index of the mountain to fire on)
 * The inputs you are given are automatically updated according to your last actions.
 **/
func main() {
	for {
		maxHeight := -1
		result := -1

		for i := 0; i < 8; i++ {
			// mountainHeight: represents the height of one mountain.
			var mountainHeight int

			_, _ = fmt.Scan(&mountainHeight)

			if mountainHeight > maxHeight {
				maxHeight = mountainHeight
				result = i
			}
		}

		fmt.Println(result)
	}
}
