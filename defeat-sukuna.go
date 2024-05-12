package sukuna

var blackFlashesRequired int = 8

func DefeatSukuna(x int) string {
	if x > blackFlashesRequired {
		return "Sukuna Defeated!!!"
	} else {
		return "Sukuna is alive and strong!! Try again!!"
	}
}
