package utilities

import (
  "strings"
  "io/ioutil"
)

func GetFormattedHTMLFromFile(replacementMap map[string]string, name string) (bool, string) {
  dat, err := ioutil.ReadFile("emails/" + name)
	if err != nil {
    return false, ""
	}

  temp := insertQueryValues(string(dat), replacementMap)
  return true, formatHTML(temp)
}

func formatHTML(email string) string {
	oneLine := strings.Replace(email, "\n", "", -1)
	noTabs := strings.Replace(oneLine, "\t", "", -1)
	return strings.Replace(noTabs, "\"", "'", -1)
}

func insertQueryValues(html string, m map[string]string) string {
	for key, value := range m {
		html = strings.Replace(html, "["+key+"]", value, -1)
	}

	return html
}
