package census

import "fmt"

type State int8

func GeoToString(base string, val int) string {
	switch val {
	case 0:
		return ""
	case -1:
		return fmt.Sprintf("%v:*", base)
	default:
		return fmt.Sprintf("%v:%v", base, val)
	}
}

const (
	All State = -1

	Alabama = iota + 1
	_

	Alaska
	AmericanSamoa
	Arizona
	Arkansas
	California
	CanalZone
	Colorado
	Connecticut
	Delaware
	DistrictOfColumbia
	Florida
	Georgia
	Guam
	Hawaii
	Idaho
	Illinois
	Indiana
	Iowa
	Kansas
	Kentucky
	Louisiana
	Maine
	Maryland
	Massachusetts
	Michigan
	Minnesota
	Mississippi
	Missouri
	Montana
	Nebraska
	Nevada
	NewHampshire
	NewJersey
	NewMexico
	NewYork
	NorthCarolina
	NorthDakota
	Ohio
	Oklahoma
	Oregon
	Pennsylvania

	_

	RhodeIsland
	SouthCarolina
	SouthDakota
	Tennessee
	Texas
	Utah
	Vermont
	Virginia
	VirginIslands
	Washington
	WestVirginia
	Wisconsin
	Wyoming

	PuertoRico = 72
)
