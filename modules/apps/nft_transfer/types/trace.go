package types

import "strings"

func ParseClass(class string) string {
	classSplit := strings.Split(class, "/")

	if classSplit[0] == class {
		return DenomTrace{
			Path:      "",
			BaseDenom: rawDenom,
		}
	}

	return DenomTrace{
		Path:      strings.Join(classSplit[:len(classSplit)-1], "/"),
		BaseDenom: classSplit[len(classSplit)-1],
	}
}
