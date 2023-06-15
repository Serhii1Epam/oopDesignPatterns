package builder

type Computer struct {
	Cpu string
	Ram int
	Mb  string
}

type IBuilder interface {
	CPU(val string) IBuilder
	RAM(val int) IBuilder
	MB(val string) IBuilder
	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() IBuilder {
	return computerBuilder{}
}

func (c computerBuilder) CPU(val string) IBuilder {
	c.cpu = val
	return c
}
func (c computerBuilder) RAM(val int) IBuilder {
	c.ram = val
	return c
}
func (c computerBuilder) MB(val string) IBuilder {
	c.mb = val
	return c
}

func (c computerBuilder) Build() Computer {
	return Computer{Cpu: c.cpu, Ram: c.ram, Mb: c.mb}
}

type officeComputerBuilder struct {
	computerBuilder
}

func NewOfficeCumputerBuilder() IBuilder {
	return officeComputerBuilder{}.CPU("Pentium III").RAM(4).MB("ASUS")
}

func (c officeComputerBuilder) Build() Computer {
	return Computer{
		Cpu: c.cpu,
		Ram: c.ram,
		Mb:  c.mb,
	}
}
