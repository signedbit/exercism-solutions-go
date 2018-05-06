package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name                string
	wins, draws, losses uint8
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) MatchesPlayed() uint16 {
	return uint16(t.Wins()) + uint16(t.Draws()) + uint16(t.Losses())
}

func (t *Team) Wins() uint8 {
	return t.wins
}

func (t *Team) Draws() uint8 {
	return t.draws
}

func (t *Team) Losses() uint8 {
	return t.losses
}

func (t *Team) Points() uint16 {
	return uint16(t.Wins())*3 + uint16(t.Draws())
}

type teamMap map[string]*Team
type teamSlice []*Team

func (tm teamMap) getTeam(name string) *Team {
	if team := tm[name]; team == nil {
		tm[name] = &Team{name: name}
	}
	return tm[name]
}

func (tm teamMap) toTeamSlice() (slice teamSlice) {
	slice = make(teamSlice, 0, len(tm))
	for _, team := range tm {
		slice = append(slice, team)
	}
	return
}

func (s teamSlice) Len() int {
	return len(s)
}

func (s teamSlice) Less(i, j int) bool {
	points1, points2 := s[i].Points(), s[j].Points()
	if points1 == points2 {
		return s[j].Name() < s[i].Name()
	}
	return points1 < points2
}

func (s teamSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Tally(in io.Reader, out io.Writer) error {
	teams := make(teamMap)
	if err := parseScores(teams, in); err != nil {
		return err
	}
	writeTable(teams, out)
	return nil
}

func parseScores(teams teamMap, in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := parseLine(teams, line); err != nil {
			return err
		}
	}
	return nil
}

func parseLine(tm teamMap, line string) error {
	parts := strings.Split(line, ";")
	if len(parts) != 3 {
		return fmt.Errorf("Invalid input %s", line)
	}

	team1, team2 := tm.getTeam(parts[0]), tm.getTeam(parts[1])
	switch outcome := parts[2]; outcome {
	case "win":
		team1.wins++
		team2.losses++
	case "loss":
		team1.losses++
		team2.wins++
	case "draw":
		team1.draws++
		team2.draws++
	default:
		return fmt.Errorf("invalid outcome %s", outcome)
	}
	return nil
}

func writeTable(tm teamMap, out io.Writer) {
	teams := tm.toTeamSlice()
	sort.Sort(sort.Reverse(teams))

	w := bufio.NewWriter(out)
	defer w.Flush()

	headerFmt := "%-30s | %2s | %2s | %2s | %2s | %2s\n"
	rowFmt := "%-30s | %2d | %2d | %2d | %2d | %2d\n"

	fmt.Fprintf(w, headerFmt, "Team", "MP", "W", "D", "L", "P")
	for _, t := range teams {
		fmt.Fprintf(w, rowFmt, t.Name(), t.MatchesPlayed(), t.Wins(), t.Draws(), t.Losses(), t.Points())
	}
}
