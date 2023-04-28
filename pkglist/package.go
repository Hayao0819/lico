package pkglist

type P string

func NewPkg(name string)(P){
	return P(name)
}


