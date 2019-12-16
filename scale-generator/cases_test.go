package scale

// Source: exercism/problem-specifications
// Commit: 65cc51b Update scale-generator version to 2.0.0
// Problem Specifications Version: 2.0.0

type scaleTest struct {
	description string
	sharp 		bool
	tonic       string
	interval    string
	expected    []string
}

var scaleTestCases = []scaleTest{
	{
		description: "Chromatic scale with sharps",
		sharp: 		true,
		tonic:       "C",
		interval:    "",
		expected:    []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	},
	{
		description: "Chromatic scale with flats",
		sharp: 		 false,
		tonic:       "F",
		interval:    "",
		expected:    []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"},
	},
	{
		description: "Simple major scale",
		sharp: 		 true,
		tonic:       "C",
		interval:    "MMmMMMm",
		expected:    []string{"C", "D", "E", "F", "G", "A", "B"},
	},
	{
		description: "Major scale with sharps",
		sharp: 		 true,
		tonic:       "G",
		interval:    "MMmMMMm",
		expected:    []string{"G", "A", "B", "C", "D", "E", "F#"},
	},
	{
		description: "Major scale with flats",
		sharp: 		 false,
		tonic:       "F",
		interval:    "MMmMMMm",
		expected:    []string{"F", "G", "A", "Bb", "C", "D", "E"},
	},
	{
		description: "Minor scale with sharps",
		sharp: 		 true,
		tonic:       "f#",
		interval:    "MmMMmMM",
		expected:    []string{"F#", "G#", "A", "B", "C#", "D", "E"},
	},
	{
		description: "Minor scale with flats",
		sharp: 		 false,
		tonic:       "bb",
		interval:    "MmMMmMM",
		expected:    []string{"Bb", "C", "Db", "Eb", "F", "Gb", "Ab"},
	},
	{
		description: "Dorian mode",
		sharp: 		 true,
		tonic:       "d",
		interval:    "MmMMMmM",
		expected:    []string{"D", "E", "F", "G", "A", "B", "C"},
	},
	{
		description: "Mixolydian mode",
		sharp: 		 false,
		tonic:       "Eb",
		interval:    "MMmMMmM",
		expected:    []string{"Eb", "F", "G", "Ab", "Bb", "C", "Db"},
	},
	{
		description: "Lydian mode",
		sharp: 		 true,
		tonic:       "a",
		interval:    "MMMmMMm",
		expected:    []string{"A", "B", "C#", "D#", "E", "F#", "G#"},
	},
	{
		description: "Phrygian mode",
		sharp: 		 true,
		tonic:       "e",
		interval:    "mMMMmMM",
		expected:    []string{"E", "F", "G", "A", "B", "C", "D"},
	},
	{
		description: "Locrian mode",
		sharp: 		 false,
		tonic:       "g",
		interval:    "mMMmMMM",
		expected:    []string{"G", "Ab", "Bb", "C", "Db", "Eb", "F"},
	},
	{
		description: "Harmonic minor",
		sharp: 		 false,
		tonic:       "d",
		interval:    "MmMMmAm",
		expected:    []string{"D", "E", "F", "G", "A", "Bb", "Db"},
	},
	{
		description: "Octatonic",
		sharp: 		 true,
		tonic:       "C",
		interval:    "MmMmMmMm",
		expected:    []string{"C", "D", "D#", "F", "F#", "G#", "A", "B"},
	},
	{
		description: "Hexatonic",
		sharp: 		 false,
		tonic:       "Db",
		interval:    "MMMMMM",
		expected:    []string{"Db", "Eb", "F", "G", "A", "B"},
	},
	{
		description: "Pentatonic",
		sharp: 		 true,
		tonic:       "A",
		interval:    "MMAMA",
		expected:    []string{"A", "B", "C#", "E", "F#"},
	},
	{
		description: "Enigmatic",
		sharp: 		 true,
		tonic:       "G",
		interval:    "mAMMMmm",
		expected:    []string{"G", "G#", "B", "C#", "D#", "F", "F#"},
	},
}
