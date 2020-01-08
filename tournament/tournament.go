package tournament

import (
	"fmt"
	"io/ioutil"
	"io"
	"strings"
	"errors"
	"sort"
)

/**
 * Represents the points Table.
 */
type Table struct {
	mp int // Matches played
	w int // Matches won
	d int // Matches Drawn; You can ignore this and make it a calculated field as well.
	l int // Matches lost
	p int // Matches played.
}


// Teams that store the table of tally.
type Teams map [string]*Table

const Format string = "%-31s| %s |  %s |  %s |  %s |  %s\n"
const FormatInt string = "%-31s|  %d |  %d |  %d |  %d |  %d\n"

func (teams Teams) WriteTally(writer io.Writer) {

	fmt.Fprintf(writer, Format, "Team", "MP", "W", "D", "L", "P")

	// Allows to store the keys which will be sorted.
	var sortedKeys [] string
	for k, _:= range teams {
		sortedKeys = append(sortedKeys, k)
	}

	// Sort alphabetically intitally
	sort.Slice(sortedKeys, func(i, j int) bool {
		return strings.Compare(sortedKeys[i], sortedKeys[j]) <= 0
	})

	// Sort by point earned.
	sort.Slice(sortedKeys, func(i,j int) bool {
		return teams[sortedKeys[i]].p > teams[sortedKeys[j]].p 
	})

	for _,k := range sortedKeys {
		v := teams[k]
		fmt.Fprintf(writer,FormatInt, k, v.mp, v.w, v.d, v.l, v.p)
	}
}

// Utility function to Write the structure.
// func writeStructure(teams Teams, writer io.Writer) {
// 	// var finalResult [] string = make([]string, 0)

// 	fmt.Fprintf(writer,"%-31s| %s |  %s |  %s |  %s |  %s\n","Team", "MP", "W", "D", "L", "P")
// 	// fmt.Println( "T: ", teams, n, len(teams))

// 	var sortedKeys [] string
// 	for k, _:= range teams {
// 		sortedKeys = append(sortedKeys, k)
// 	}

// 	sort.Slice(sortedKeys, func(i,j int) bool {
// 		if teams[sortedKeys[i]].p > teams[sortedKeys[j]].p {
// 			return true
// 		} else if teams[sortedKeys[i]].p == teams[sortedKeys[j]].p {
// 			if strings.Compare(sortedKeys[i], sortedKeys[j]) <= 0 {
// 				return true
// 			}
// 		}
// 		return false
// 	})

// 	for _,k := range sortedKeys {
// 		v := teams[k]
// 		fmt.Fprintf(writer,"%-31s|  %d |  %d |  %d |  %d |  %d\n", k, v.mp, v.w, v.d, v.l, v.p)
// 	}

// }


/**
 * Package bufio implements buffered I/O. It wraps an io.Reader 
 * or io.Writer object, creating another object (Reader or Writer) 
 * that also implements the interface but provides buffering and 
 * some help for textual I/O. 
 */
func Tally(reader io.Reader, writer io.Writer) error {
	var teams Teams = make(map[string]*Table)
	data, err:= ioutil.ReadAll(reader)
	// fmt.Printf("Input - %s", data)
	if err == nil {
		// Read the individual line with match and result
		for _, val := range strings.Split(fmt.Sprintf("%s",data), "\n") {
			// Allegoric Alaskans;Blithering Badgers;win
			tm := strings.Split(val, ";") 
			if len(tm) <  3 {
				// fmt.Println(" Skipping..")
			}else { //Else not complete data.
				switch tm[2] {
					case "win":
						if teams[string(tm[0])] == nil {
							teams[string(tm[0])] = &Table {
								0,0,0,0,0,
							}
						}

						teams[string(tm[0])].w += 1
						teams[string(tm[0])].mp += 1
						teams[string(tm[0])].p += 3

						if teams[string(tm[1])] == nil {
							teams[string(tm[1])] = &Table {
								0,0,0,0,0,
							}
						}

						teams[string(tm[1])].l += 1
						teams[string(tm[1])].mp += 1
						// teamA := teams[tm[0]]
						// teamB := teams[tm[1]]
						// AdjustWin(&teamA)
						// AdjustLoss(&teamB)
						
					case "loss":

						if teams[string(tm[0])] == nil{
							teams[string(tm[0])] = &Table {
								0,0,0,0,0,
							}
						}

						teams[string(tm[0])].l += 1
						teams[string(tm[0])].mp += 1

						if teams[string(tm[1])] == nil {
							teams[string(tm[1])] = &Table {
								0,0,0,0,0,
							}
						}
						
						teams[string(tm[1])].w += 1
						teams[string(tm[1])].mp += 1
						teams[string(tm[1])].p += 3
						
					case "draw":
						
						if teams[string(tm[0])] == nil{
							teams[string(tm[0])] = &Table {
								0,0,0,0,0,
							}
						}

						teams[string(tm[0])].d += 1
						teams[string(tm[0])].mp += 1
						teams[string(tm[0])].p += 1

						if teams[string(tm[1])] == nil {
							teams[string(tm[1])] = &Table {
								0,0,0,0,0,
							}
						}

						teams[string(tm[1])].d += 1
						teams[string(tm[1])].mp += 1
						teams[string(tm[1])].p += 1

					default:
						return errors.New("invalid value.")
				}
			}
		}

		if len(teams) == 0 {
			return errors.New(" no element")
		}
		teams.WriteTally(writer)
	}
	return nil
}
	

