package map2struct

type M2SType int

const (
	M2SUndefined M2SType = iota
	M2SString
	M2SSInt
	M2SSFloat
	M2SList
	M2SBool
)

type M2STag struct {
	FName string
	Name  string
	Type  M2SType
}

func New(fname, name string, typ M2SType) M2STag {
	return M2STag{
		FName: fname,
		Name:  name,
		Type:  typ,
	}
}
