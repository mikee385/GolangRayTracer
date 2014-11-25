package geometry

import "math"

type Point3D struct {
	X float32
	Y float32
	Z float32
}

var origin = Point3D{X: 0.0, Y: 0.0, Z: 0.0}

func NewPoint(x float32, y float32, z float32) Point3D {
	return Point3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewPoint_FromVector(vector Vector3D) Point3D {
	return Point3D{
		X: vector.X,
		Y: vector.Y,
		Z: vector.Z,
	}
}

func Origin() Point3D {
	return origin
}

func (point Point3D) Translate_Dist(direction Direction3D, magnitude float32) Point3D {
	return Point3D{
		X: point.X + direction.X()*magnitude,
		Y: point.Y + direction.Y()*magnitude,
		Z: point.Z + direction.Z()*magnitude,
	}
}

func (point Point3D) Translate_Vec(vector Vector3D) Point3D {
	return Point3D{
		X: point.X + vector.X,
		Y: point.Y + vector.Y,
		Z: point.Z + vector.Z,
	}
}

func (point Point3D) Rotate(matrix Matrix3D) Point3D {
	return Point3D{
		X: point.X*matrix.X.X + point.Y*matrix.Y.X + point.Z*matrix.Z.X,
		Y: point.X*matrix.X.Y + point.Y*matrix.Y.Y + point.Z*matrix.Z.Y,
		Z: point.X*matrix.X.Z + point.Y*matrix.Y.Z + point.Z*matrix.Z.Z,
	}
}

func (point Point3D) Equals(other Point3D) bool {
	return (point.X == other.X) &&
		(point.Y == other.Y) &&
		(point.Z == other.Z)
}

func (point Point3D) EqualsTol(other Point3D, tolerance float32) bool {
	return float32(math.Abs(float64(point.X-other.X))) < tolerance &&
		float32(math.Abs(float64(point.Y-other.Y))) < tolerance &&
		float32(math.Abs(float64(point.Z-other.Z))) < tolerance
}

func Distance(point1 Point3D, point2 Point3D) float32 {
	return NewVector(
		point1.X-point2.X,
		point1.Y-point2.Y,
		point1.Z-point2.Z).Magnitude()
}

func Midpoint(point1 Point3D, point2 Point3D) Point3D {
	return Point3D{
		X: 0.5 * (point1.X + point2.X),
		Y: 0.5 * (point1.Y + point2.Y),
		Z: 0.5 * (point1.Z + point2.Z),
	}
}
