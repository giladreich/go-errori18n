package errori18n

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type LangCode uint32

const (
	LANG_UNKNOWN LangCode = iota
	LANG_EN
	LANG_DE
)

func (l LangCode) String() string {
	switch l {
	case LANG_EN:
		return "en"
	case LANG_DE:
		return "de"
	default:
		return "unknown"
	}
}

var (
	langCode LangCode
	langMap  map[ErrorCode]string

	//go:embed language/en.yaml
	enYaml string

	//go:embed language/de.yaml
	deYaml string

	langFiles = map[LangCode]string{
		LANG_EN: enYaml,
		LANG_DE: deYaml,
	}
)

func init() {
	SetLanguage(LANG_EN)
}

func SetLanguage(code LangCode) {
	if langCode == code {
		return
	}

	var langYaml string
	var found bool
	if langYaml, found = langFiles[code]; !found {
		panic("Undefined language.")
	}

	langMap = nil
	if err := yaml.Unmarshal([]byte(langYaml), &langMap); err != nil {
		panic(err)
	}
	langCode = code
}

func (e ErrorCode) String() string {
	if msg, found := langMap[e]; found {
		return msg
	}
	return fmt.Sprintf("(null message - %s:%#v)", langCode, e)
}
