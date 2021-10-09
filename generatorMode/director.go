package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{builder: b}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setNumFloor()
	d.builder.setDoorType()
	d.builder.setWindowType()
	return d.builder.getHouse()
}