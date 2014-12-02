package color

type ColorRGB struct {
	Red   float32
	Green float32
	Blue  float32
}

func New(red float32, green float32, blue float32) ColorRGB {
	return ColorRGB{Red: red, Green: green, Blue: blue}
}

func (self ColorRGB) Add(other ColorRGB) ColorRGB {
	return ColorRGB{
		Red:   self.Red + other.Red,
		Green: self.Green + other.Green,
		Blue:  self.Blue + other.Blue,
	}
}

func (self ColorRGB) Scale(scale float32) ColorRGB {
	return ColorRGB{
		Red:   self.Red * scale,
		Green: self.Green * scale,
		Blue:  self.Blue * scale,
	}
}

func (self ColorRGB) Mul(other ColorRGB) ColorRGB {
	return ColorRGB{
		Red:   self.Red * other.Red,
		Green: self.Green * other.Green,
		Blue:  self.Blue * other.Blue,
	}
}

func White() ColorRGB {
	return ColorRGB{Red: 1.0, Green: 1.0, Blue: 1.0}
}

func Black() ColorRGB {
	return ColorRGB{Red: 0.0, Green: 0.0, Blue: 0.0}
}

func Red() ColorRGB {
	return ColorRGB{Red: 1.0, Green: 0.0, Blue: 0.0}
}

func Green() ColorRGB {
	return ColorRGB{Red: 0.0, Green: 1.0, Blue: 0.0}
}

func Blue() ColorRGB {
	return ColorRGB{Red: 0.0, Green: 0.0, Blue: 1.0}
}
