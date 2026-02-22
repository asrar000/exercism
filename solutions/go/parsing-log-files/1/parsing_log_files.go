package parsinglogfiles
import "regexp"
func IsValidLine(text string) bool {
	re:=regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re:=regexp.MustCompile(`<[-~*=]*>`)
    return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
    re:=regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    count:=0
    for _,line := range lines{
        if(re.MatchString(line)){
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re:=regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))
    for index, line := range lines {
		match := re.FindStringSubmatch(line)

		if match == nil {
			result[index] = line
			continue
		}

		username := match[1]
		result[index] = "[USR] " + username + " " + line
	}
    return result
}
