package client

import (
	"strconv"
)

func StrInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

//type Duration time.Duration
//
//func (i Duration) UnmarshalText(text []byte) error {
//	i = ParseDuration(string(text))
//	return nil
//}
//func (i Duration) MarshalText() (text []byte, err error) {
//	return []byte(fmt.Sprintf("test")), nil
//}
//
//type durationUnit struct {
//	U   byte
//	Val uint64
//}
//
//var durationUnits = []durationUnit{
//	durationUnit{'D', 24 * 3600e3},
//	durationUnit{'H', 3600e3},
//	durationUnit{'M', 60e3},
//	durationUnit{'S', 1e3},
//}
//
//func ParseDuration(d string) (durationVal Duration) {
//	if d[0] != 'P' {
//		return //Must start with P
//	}
//	if len(d) < 4 {
//		return //No value defined
//	}
//	val := 0
//	unit := 0
//	for i := 1; i < len(d); i++ {
//		switch {
//		case d[i] >= '0' && d[i] <= '9':
//			val = val*10 + int(d[i]) - '0'
//		case d[i] == 'T':
//			if val > 0 {
//				return // Malformed time
//			}
//			if unit <= 1 {
//				unit = 1
//			}
//		default:
//			for d[i] != durationUnits[unit].U && unit < len(durationUnits) {
//				unit++
//			}
//			if d[i] == durationUnits[unit].U {
//				durationVal += durationUnits[unit].Val * uint64(val)
//				val = 0
//				unit++
//			} else {
//				fmt.Println("found val", string(d[i]))
//				return
//			}
//		}
//
//	}
//	return
//}
//
//func quote(s string) string {
//	return "\"" + s + "\""
//}

//var errLeadingInt = errors.New("time: bad [0-9]*") // never printed

//// leadingInt consumes the leading [0-9]* from s.
//func leadingInt(s string) (x int64, rem string, err error) {
//	i := 0
//	for ; i < len(s); i++ {
//		c := s[i]
//		if c < '0' || c > '9' {
//			break
//		}
//		if x > (1<<63-1)/10 {
//			// overflow
//			return 0, "", errLeadingInt
//		}
//		x = x*10 + int64(c) - '0'
//		if x < 0 {
//			// overflow
//			return 0, "", errLeadingInt
//		}
//	}
//	return x, s[i:], nil
//}
//
//// leadingFraction consumes the leading [0-9]* from s.
//// It is used only for fractions, so does not return an error on overflow,
//// it just stops accumulating precision.
//func leadingFraction(s string) (x int64, scale float64, rem string) {
//	i := 0
//	scale = 1
//	overflow := false
//	for ; i < len(s); i++ {
//		c := s[i]
//		if c < '0' || c > '9' {
//			break
//		}
//		if overflow {
//			continue
//		}
//		if x > (1<<63-1)/10 {
//			// It's possible for overflow to give a positive number, so take care.
//			overflow = true
//			continue
//		}
//		y := x*10 + int64(c) - '0'
//		if y < 0 {
//			overflow = true
//			continue
//		}
//		x = y
//		scale *= 10
//	}
//	return x, scale, s[i:]
//}
//
//var unitMap = map[string]int64{
//	"DT": int64(24 * time.Hour),
//	"S":  int64(time.Second),
//	"s":  int64(time.Second),
//	"M":  int64(time.Minute),
//	"m":  int64(time.Minute),
//	"H":  int64(time.Hour),
//	"h":  int64(time.Hour),
//}
//
//// ParseDuration parses a Cisco duration string.
//func ParseDuration(s string) Duration {
//	s = strings.TrimPrefix(s, "P")
//	s = strings.TrimPrefix(s, "T")
//	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
//	//orig := s
//	var d int64
//	neg := false
//
//	// Consume [-+]?
//	if s != "" {
//		c := s[0]
//		if c == '-' || c == '+' {
//			neg = c == '-'
//			s = s[1:]
//		}
//	}
//	// Special case: if all that is left is "0", this is zero.
//	if s == "0" {
//		return 0 //, nil
//	}
//	if s == "" {
//		return 0 //, errors.New("time: invalid duration " + quote(orig))
//	}
//	for s != "" {
//		var (
//			v, f  int64       // integers before, after decimal point
//			scale float64 = 1 // value = v + f/scale
//		)
//
//		var err error
//
//		// The next character must be [0-9.]
//		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
//			return 0 //, errors.New("time: invalid duration " + quote(orig))
//		}
//		// Consume [0-9]*
//		pl := len(s)
//		v, s, err = leadingInt(s)
//		if err != nil {
//			return 0 //, errors.New("time: invalid duration " + quote(orig))
//		}
//		pre := pl != len(s) // whether we consumed anything before a period
//
//		// Consume (\.[0-9]*)?
//		post := false
//		if s != "" && s[0] == '.' {
//			s = s[1:]
//			pl := len(s)
//			f, scale, s = leadingFraction(s)
//			post = pl != len(s)
//		}
//		if !pre && !post {
//			// no digits (e.g. ".s" or "-.s")
//			return 0 //, errors.New("time: invalid duration " + quote(orig))
//		}
//
//		// Consume unit.
//		i := 0
//		for ; i < len(s); i++ {
//			c := s[i]
//			if c == '.' || '0' <= c && c <= '9' {
//				break
//			}
//		}
//		if i == 0 {
//			return 0 //, errors.New("time: missing unit in duration " + quote(orig))
//		}
//		u := s[:i]
//		s = s[i:]
//		unit, ok := unitMap[u]
//		if !ok {
//			return 0 //, errors.New("time: unknown unit " + quote(u) + " in duration " + quote(orig))
//		}
//		if v > (1<<63-1)/unit {
//			// overflow
//			return 0 //, errors.New("time: invalid duration " + quote(orig))
//		}
//		v *= unit
//		if f > 0 {
//			// float64 is needed to be nanosecond accurate for fractions of hours.
//			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
//			v += int64(float64(f) * (float64(unit) / scale))
//			if v < 0 {
//				// overflow
//				return 0 //, errors.New("time: invalid duration " + quote(orig))
//			}
//		}
//		d += v
//		if d < 0 {
//			// overflow
//			return 0 //, errors.New("time: invalid duration " + quote(orig))
//		}
//	}
//
//	if neg {
//		d = -d
//	}
//	return Duration(d) //, nil
//}
