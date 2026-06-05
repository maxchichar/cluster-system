package houses

type House struct {
	Name   string
	Tables []int
}

var Houses = map[string]House{
	"KER": {
		Name:   "Kernel",
		Tables: []int{1, 2, 3, 4, 5, 6},
	},
	"COM": {
		Name:   "Compiler",
		Tables: []int{7, 8, 9, 10, 11, 12, 25},
	},
	"RNT": {
		Name:   "Runtime",
		Tables: []int{13, 14, 15, 16, 17, 18},
	},
	"ALG": {
		Name:   "Algorithm",
		Tables: []int{19, 20, 21, 22, 23, 24},
	},
}

func GetHouseKey(table int) string {
	for k, h := range Houses {
		for _, t := range h.Tables {
			if t == table {
				return k
			}
		}
	}
	return ""
}