package services

type IDecorator interface {
	Decorate(function IDecoratedFunction) IDecoratedFunction
}

type IDecoratedFunction interface {
	Run()
}
