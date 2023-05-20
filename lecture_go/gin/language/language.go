//language 개념적인 예제
type Language struct {
	Golang string
}

func (p *Language) temp() {
}
func (p *Language) temp1() {
}

type Developer struct {
	Language
	Chain string
}

func (d *Developer) test() {
	d.Language.Golang
}
func (d *Developer) test1() {
}