package sysloc

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/text/language"
)

var defaultLocale = language.English

// GetLocale get current locale that used by system
func GetLocale() (language.Tag, error) {
	switch os := runtime.GOOS; os {
	case "darwin", "linux":
		return getLocaleUnix()
	case "windows":
		return getLocaleWindows()
	default:
		return defaultLocale, fmt.Errorf("unknown os")
	}
}

func getLocaleUnix() (language.Tag, error) {
	locale := os.Getenv("LANG")
	locale = strings.SplitN(locale, ".", 2)[0]

	languageTag, err := language.Parse(locale)
	if err != nil {
		return defaultLocale, fmt.Errorf("failed to parse locale %s: %v", locale, err)
	}

	return languageTag, nil
}

func getLocaleWindows() (language.Tag, error) {
	cmdWmic := exec.Command("wmic", "os", "get", "locale")
	btOutput, err := cmdWmic.CombinedOutput()
	if err != nil {
		return defaultLocale, fmt.Errorf("failed to run wmic: %v", err)
	}

	strLCID := string(btOutput)
	strLCID = strings.ToLower(strLCID)
	strLCID = strings.Replace(strLCID, "locale", "", -1)
	strLCID = strings.TrimSpace(strLCID)

	lcid, err := strconv.ParseInt(strLCID, 16, 64)
	if err != nil {
		return defaultLocale, fmt.Errorf("failed to parse lcid: %v", err)
	}

	locale, exist := mapLCID[lcid]
	if !exist {
		return defaultLocale, fmt.Errorf("no locale exists for lcid %d", lcid)
	}

	languageTag, err := language.Parse(locale)
	if err != nil {
		return defaultLocale, fmt.Errorf("failed to parse locale %s: %v", locale, err)
	}

	return languageTag, nil
}
