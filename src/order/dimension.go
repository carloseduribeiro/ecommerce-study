package order

type Dimension struct {
	width  float64
	height float64
	length float64
}

func NewDimension(width, height, length float64) Dimension {
	return Dimension{
		width:  width,
		height: height,
		length: length,
	}
}

func (d Dimension) Volume() float64 {
	return d.width / 100 * d.height / 100 * d.length / 100
}
