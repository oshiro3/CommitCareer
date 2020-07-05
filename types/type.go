package types

type Branch struct {
	name string
}

type Commit struct {
	Sha   string
	Title string
}

type PushRequest struct {
	base    Branch
	compare Branch
	commits []Commit
}
