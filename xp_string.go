package xparam

import (
	"fmt"
	"strings"
)

//------------------------------------------------------------
// Xparam access for string
//------------------------------------------------------------

// Gets parameter as string.
func (xp XP) As_String(key string) (str string) {

	if val, ok := xp[key]; ok && val != nil {
		str = strings.TrimSpace(fmt.Sprint(val))
	}
	return
}

// Gets parameter as string pointer allowing for the nil option.
func (xp XP) As_StringNil(key string) (str *string) {

	if val, ok := xp[key]; ok && val != nil {
		s := strings.TrimSpace(fmt.Sprint(val))
		str = &s
	}
	return
}

// Sets parameter as string.
func (xp XP) To_String(to *map[string]string, key string) {
	if val, ok := xp[key]; ok && val != nil {
		str := strings.TrimSpace(fmt.Sprint(val))
		if str != "" {
			(*to)[key] = str
		}
	}
}
