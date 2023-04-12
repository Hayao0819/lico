package conf

import (
	"fmt"
	"strings"
)

type EntryOption struct{
	//WithRoot bool
}

func DefaultOption()(*EntryOption){
	return &EntryOption{
		//WithRoot: false,
	}
}

func ParseEntryOption(opt string)(*EntryOption, error){
	o := DefaultOption()
	for _, s := range strings.Split(opt, ","){
		s = strings.TrimSpace(s)
		switch strings.ToLower(s) {
			/*
			case "withroot":
				o.WithRoot=true
			case "withoutroot", "noroot":
				o.WithRoot=false
			*/
			default:
				return nil, fmt.Errorf("unknown option: %s\n", s)
		}
	}
	return o, nil
}

