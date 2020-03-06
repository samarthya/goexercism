package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden type garden
type Garden struct {
	m map[string][]string
}

// NewGarden New garden
func NewGarden(diagram string, children []string) (g *Garden, err error) {
	g = &Garden{m: make(map[string][]string, 0)}

	ch := make([]string, len(children))
	copy(ch, children)
	sort.Strings(ch)

	if strings.HasPrefix(diagram, "\n") == false {
		return nil, errors.New(" window is missing")
	}

	diagram = strings.TrimLeft(diagram, "\n")
	ws := strings.Split(diagram, "\n")
	for i, c := range ch {
		if _, ok := g.m[c]; ok {
			return nil, errors.New(" duplicate value")
		}
		// Add the kids
		g.m[c] = []string{}

		for _, w := range ws {

			if len(w) != 2*len(children) {
				return nil, errors.New(" invalid length")
			}
			for planti, r := range w {
				if i != 0 && planti < (i*2) {
					continue
				}
				if planti > (i*2) && (planti%2) == 0 {
					break
				}
				switch {
				case r == 'V':
					g.m[c] = append(g.m[c], "violets")
				case r == 'R':
					g.m[c] = append(g.m[c], "radishes")
				case r == 'G':
					g.m[c] = append(g.m[c], "grass")
				case r == 'C':
					g.m[c] = append(g.m[c], "clover")
				default:
					return nil, errors.New(" invalid plant")
				}

				if len(g.m[c]) > 4 {
					return nil, errors.New(" more cups for the same not expected")
				}
			}
		}
	}

	return g, nil
}

//Plants plants
func (g *Garden) Plants(child string) ([]string, bool) {
	if val, ok := g.m[child]; ok {
		return val, true
	}
	return nil, false
}
