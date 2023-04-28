package pkg

import (
	"encoding/json"
	//"fmt"
	"os"

	"github.com/Hayao0819/lico/vars"
)

type OSList map[string][]P


type List map[string]OSList


func ReadList()(*List, error){
	file, err := os.ReadFile(vars.GetPkgList())
	if err != nil{
		return nil, err
	}
	
	pkglist := List{}

	if json.Unmarshal(file, &pkglist); err !=nil{
		return nil, err
	}

	//fmt.Println(pkglist)

	return &pkglist, nil
}

func (p *List)OSList()([]string){
	keys := []string{}
	for k := range *p{
		keys = append(keys, k)
	}
	return keys
}


func (p *List)GetOS(name string)(*OSList){
	oslist := (*p)[name]
	return &oslist
}

func (o *OSList)GetPkgs(id string)(*[]P){
	p := (*o)[id]
	return &p
}

func (p *List)GetCurrent()(*[]P, error){
	// あとで実装する
	return nil, nil
}
