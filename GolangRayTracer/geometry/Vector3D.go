package geometry

import "math"

type Vector3D struct {
	X float32
	Y float32
	Z float32
}

func NewVector(x float32, y float32, z float32) Vector3D {
	return Vector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewVector_FromPoint(point Point3D) Vector3D {
	return Vector3D{
		X: point.X,
		Y: point.Y,
		Z: point.Z,
	}
}

func (vector Vector3D) Magnitude() float32 {
	return float32(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)))
}

func (vector Vector3D) Scale(scale float32) Vector3D {
	return Vector3D{
		X: vector.X * scale,
		Y: vector.Y * scale,
		Z: vector.Z * scale,
	}
}
