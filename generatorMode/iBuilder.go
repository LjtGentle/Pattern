package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

const (
	NormalType string = "normal"
	IglooType string = "igloo"
)

func getBuilder(builderType string) iBuilder {
	if builderType == NormalType {
		return  &normalBuilder{}
	}
	if builderType == IglooType {
		return  &iglooBuilder{}
	}
	return nil
}

