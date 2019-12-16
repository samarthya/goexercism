package scale

import ( 
	"fmt" 
	"strings"
)

const chromaticNotesSharp string = "A,A#,B,C,C#,D,D#,E,F,F#,G,G#"
const chromaticNotesFlat string = "A,Bb,B,C,Db,D,Eb,E,F,Gb,G,Ab"

/** 
 * Find the right tonic to create the musical scale.
 */
func getSkipLength(interval string) int {
	interval = strings.TrimSpace(interval)

	if interval == "" {
		return 1
	}

	switch interval {
		// Whole step.
		case "M":
			return 2
		// Half step; Adjacent Node.
		case "m":
			return 1
		// Augmented first
		case "A":
			return 3
		default:
			return 1
	}
}

/**
 * Given a start tonic, browse through the interval to create the right musical scale.
 */
func findTonicReturnMusic(tonic string, interval string, sharp bool) [] string {
	// For the Musical scale as per the interval.
	var newPitch [] string	

	// Will be used to read the Sharp or Flat notation.
	var chromaticNodes [] string

	var totalLen int = 12

	var stepInterval int = 0

	// Locate the length to skip in the order.
	skipLength := getSkipLength(interval)

	// Trim any spaces if we have any.
	tonic = strings.TrimSpace(tonic)

	// Find the index to Start
	startIndex, chromaticNodes := findIndexInNotes(tonic, sharp)	
	// newPitch = append(newPitch, tonic)


	if interval == "" {
		for {
			if(startIndex < totalLen) {
				if noteAlreadyPresent(newPitch, chromaticNodes[startIndex]) {
					break
				}
				newPitch = append(newPitch, chromaticNodes[startIndex])
			} else {
				// if(len(newPitch) < totalLen)
				// beak
				startIndex = 0
				continue
			}

			startIndex += skipLength
		}
	} else {

		var intervals = [] rune(interval)

		for stepInterval < len(intervals) {
			if(startIndex < totalLen) {
				if noteAlreadyPresent(newPitch, chromaticNodes[startIndex]) {
					break
				}
				newPitch = append(newPitch, chromaticNodes[startIndex])
			} else {
				// if(len(newPitch) < totalLen)
				// beak
				startIndex -= 12
				continue
			}

			skipLength = getSkipLength(string(intervals[stepInterval]))
			stepInterval++
			startIndex += skipLength
		}
	}
	

	return newPitch
}

func noteAlreadyPresent(music [] string, tonic string) bool {
	for _, v := range music {
		if strings.Compare(v, tonic) == 0 {
			return true
		}
	}

	return false
}

func returnSharpArray() ([] string) {
	return strings.Split(chromaticNotesSharp, ",")
}

func returnFlatArray() ([] string) {
	return strings.Split(chromaticNotesFlat, ",")
}

func findIndexInNotes(tonic string, sharp bool) (int, [] string) {
	tones := returnSharpArray()

	if sharp {
		for i, val := range tones {
			if strings.Compare(val, tonic) == 0 || strings.Compare(strings.ToUpper(val), strings.ToUpper(tonic)) == 0{
				return i, tones
			}
		}
	} else {
		tones = returnFlatArray()

		for i, val := range tones {
			if strings.Compare(val, tonic) == 0 || strings.Compare(strings.ToUpper(val), strings.ToUpper(tonic)) == 0{
				return i, tones
			}
		}
	}
	

	
	return -1, nil
}

// Scale Generates some music. Given a tonic and a starting note. 
// Generate the musical scale starting with the tonic and following 
// the specified interval pattern. m (Minor step) or M (Whole step)
func Scale(tonic string, interval string, sharp bool) [] string {

	fmt.Println(tonic, " Tonic ", interval, " Interval ", sharp, " Sharp or Flat ")
	newNotes := findTonicReturnMusic(tonic, interval, sharp)

	return newNotes
}