package conf

import (
	"fmt"
	"strings"
)

// 現在準備中

type EntryOption struct {
	//WithRoot bool
}

func DefaultOption() *EntryOption {
	return &EntryOption{
		//WithRoot: false,
	}
}

func ParseEntryOption(opt string) (*EntryOption, error) {
	o := DefaultOption()
	for _, s := range strings.Split(opt, ",") {
		s = strings.TrimSpace(s)
		switch strings.ToLower(s) {
		case "":
			continue
		default:
			return nil, fmt.Errorf("unknown option: %s", s)
		}
	}
	return o, nil
}
