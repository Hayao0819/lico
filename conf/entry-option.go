package conf

import (
	"fmt"
	"os"
	"strings"
)

// 現在準備中

type EntryOption struct {
	Template bool
}

func DefaultOption() *EntryOption {
	return &EntryOption{
		Template: false,
	}
}

func ParseEntryOption(opt string) (*EntryOption, error) {
	o := DefaultOption()
	for _, s := range strings.Split(opt, ",") {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "#"){
			continue
		}
		switch strings.ToLower(s) {
		case "template":
			fmt.Fprintln(os.Stderr, "Currently, template mode is not supported")
			o.Template=true
		case "no-template" , "notemplate":
			o.Template=false
		case "":
			continue
		default:
			return nil, fmt.Errorf("unknown option: %s", s)
		}
	}
	return o, nil
}
