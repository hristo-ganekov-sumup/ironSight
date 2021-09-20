package common

//Structures and utility

type TargetPair struct {
	Target string
	Port string
}

type InEg struct {
	Ingress []TargetPair
	Egress []TargetPair
}


type Sgs map[string]InEg


func GetString(ref *string) string {
	return *ref
}

func NewSgExploded() *Sgs {
	sge := make(Sgs)
	return &sge
}

func NewSgExplodedEntry() *TargetPair {
	sgee := new(TargetPair)
	return sgee
}

func PointersOf(v []string) []*string {
	out := make([]*string, len(v))
	for i := 0; i < len(v); i++ {
		out[i] = &v[i]
	}
	return out
}
