package pkglist

import (
	"os"

	"github.com/Hayao0819/lico/vars"
)

type OSInfo map[string]struct{
	Id string `json:"id"`
	Pkgs []P `json:"pkgs"`
}

type PkgList []OSInfo

func ReadPkgList()(*PkgList, error){
	file, err := os.Open(vars.PkgListFile)
	if err != nil{
		return nil, err
	}
	defer file.Close()


	// いい感じに構造体にデコードする処理をあとで書く

	return nil, nil
}

