package Shapes

type Shape struct{
    name  string
}

func (s Shape) GetName() string{
    return s.name
}
