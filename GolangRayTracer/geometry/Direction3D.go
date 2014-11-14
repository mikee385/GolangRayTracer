package geometry

type Direction3D struct {
	direction Vector3D
}

func NewDirection(x float32, y float32, z float32) Direction3D {
	return Direction3D{direction: NewVector(x, y, z)}
}

func NewDirection_FromVector(vector Vector3D) Direction3D {
	var magnitude = vector.Magnitude()
	if magnitude > 0.0 {
		return Direction3D{direction: vector.Scale(1.0 / magnitude)}
	} else {
		return Direction3D{direction: vector}
	}
}
