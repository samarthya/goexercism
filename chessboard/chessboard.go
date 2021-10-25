package chessboard

import "log"

const RANKS = "ABCDEFGH"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (m int) {
	if r, ok := cb[rank]; ok {
		for _, v := range r {
			if v == true {
				m++
			}
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (m int) {
	if (file > 8) || (file <= 0) {
		log.Println(" Negative case")
		return
	}

	for _, v := range RANKS {
		ranks := cb[string(v)]

		if ranks[file-1] {

			m += 1
		}
	}

	log.Println(" Total: ", m)
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (m int) {
	for _, v := range cb {
		m += len(v)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (m int) {
	for _, v := range cb {
		for _, b := range v {
			if b {
				m += 1
			}
		}
	}
	return
}
