package pkglist

type OS map[string]struct{
	Id string `json:"id"`
	Pkgs PList `json:"pkgs"`
}




