package kmnbot

func filterNotEmpty(sa []string) (res []string) {
	for _, str := range sa {
		if str != "" {
			res = append(res, str)
		}
	}
	return
}
