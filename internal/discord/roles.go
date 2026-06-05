package discord

func getRoleID(house string) string {
	switch house {
	case "KER":
		return "1512484236160925956"
	case "COM":
		return "1512484234554773616"
	case "RNT":
		return "1512484229928325171"
	case "ALG":
		return "1512484230909661334"
	default:
		return ""
	}
}