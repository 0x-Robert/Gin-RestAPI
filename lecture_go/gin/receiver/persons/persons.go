package persons



type Persons struct {
	Name string
	Age  int
	Pnum string
}

func (p Persons) Alice() { //밸류 리시버
	p.Name = "alice"
	p.Age = 7
	p.Pnum = "01012345678"
}

func (p *Persons) Bob() { //포인터 리시버
	p.Name = "bob"
	p.Age = 5
	p.Pnum = "01098765432"
}





/*
### Value(값 리시버)와 Pointer(포인터 리시버)의 차이

> **Value 리시버**는 해당 구조체 객체의 복사본을 만들어 함수에 전달합니다. 즉, 메모리의 다른 위치에 존재하기에, 함수 안에서 수행된 변경 사항에 대해서 원래 객체는 변경되지 않습니다.

**Pointer 리시버**는 해당 구조체 객체의 주소를 함수에 전달합니다. 즉, 동일한 메모리 위치에 대해 참조하기에, 함수 안에서 수행된 변경 사항에 대해서 원래 객체가 변경됩니다.
>

밸류리시버는 객체 복사본을 사용하는 것이기 때문에 데이터의 복사본이나 조작데이터로 새로운 조작을 하고 원본데이터는 놔둘때 사용하고
포인터 리시버는 원본 객체를 변경할 때 사용하는 것인가??
*/