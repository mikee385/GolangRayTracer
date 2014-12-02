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

func NewVector_BetweenPoints(from Point3D, to Point3D) Vector3D {
	return Vector3D{
		X: to.X - from.X,
		Y: to.Y - from.Y,
		Z: to.Z - from.Z,
	}
}

func NewVector_FromPoint(point Point3D) Vector3D {
	return Vector3D{
		X: point.X,
		Y: point.Y,
		Z: point.Z,
	}
}

func Zero() Vector3D {
	return Vector3D{X: 0.0, Y: 0.0, Z: 0.0}
}

func (vector Vector3D) Magnitude() float32 {
	return float32(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)))
}

func (vector Vector3D) ToUnit() Direction3D {
	return NewDirection_FromVector(vector)
}

func (vector Vector3D) ToOrthonormalBasis() Matrix3D {
	return NewDirection_FromVector(vector).ToOrthonormalBasis()
}

func (vector Vector3D) Scale(scale float32) Vector3D {
	return Vector3D{
		X: vector.X * scale,
		Y: vector.Y * scale,
		Z: vector.Z * scale,
	}
}

func (vector Vector3D) Projection_Dir(direction Direction3D) Vector3D {
	return direction.ToVector().Scale(Dot(vector, direction.ToVector()))
}

func (vector Vector3D) Projection_Vec(direction Vector3D) Vector3D {
	var denominator = Dot(direction, direction)
	if denominator > 0.0 {
		return direction.Scale(Dot(vector, direction) / denominator)
	} else {
		return Zero()
	}
}

func (vector Vector3D) Rotate(matrix Matrix3D) Vector3D {
	return Vector3D{
		X: vector.X*matrix.X.X + vector.Y*matrix.Y.X + vector.Z*matrix.Z.X,
		Y: vector.X*matrix.X.Y + vector.Y*matrix.Y.Y + vector.Z*matrix.Z.Y,
		Z: vector.X*matrix.X.Z + vector.Y*matrix.Y.Z + vector.Z*matrix.Z.Z,
	}
}

func (vector Vector3D) Equals(other Vector3D) bool {
	return (vector.X == other.X) &&
		(vector.Y == other.Y) &&
		(vector.Z == other.Z)
}

func (vector Vector3D) EqualsTol(other Vector3D, tolerance float32) bool {
	return float32(math.Abs(float64(vector.X-other.X))) < tolerance &&
		float32(math.Abs(float64(vector.Y-other.Y))) < tolerance &&
		float32(math.Abs(float64(vector.Z-other.Z))) < tolerance
}

func (vector Vector3D) Add(other Vector3D) Vector3D {
	return NewVector(
		vector.X+other.X,
		vector.Y+other.Y,
		vector.Z+other.Z,
	)
}

func (vector Vector3D) Sub(other Vector3D) Vector3D {
	return NewVector(
		vector.X-other.X,
		vector.Y-other.Y,
		vector.Z-other.Z,
	)
}

func (vector Vector3D) Mul(scale float32) Vector3D {
	return vector.Scale(scale)
}

func (vector Vector3D) Div(scale float32) Vector3D {
	var inv_scale = 1.0 / scale
	return vector.Scale(inv_scale)
}

func (vector Vector3D) Neg() Vector3D {
	return NewVector(
		-vector.X,
		-vector.Y,
		-vector.Z,
	)
}

func Dot(vector1 Vector3D, vector2 Vector3D) float32 {
	return vector1.X*vector2.X +
		vector1.Y*vector2.Y +
		vector1.Z*vector2.Z
}

func Cross(vector1 Vector3D, vector2 Vector3D) Vector3D {
	return NewVector(
		vector1.Y*vector2.Z-vector1.Z*vector2.Y,
		vector1.Z*vector2.X-vector1.X*vector2.Z,
		vector1.X*vector2.Y-vector1.Y*vector2.X,
	)
}
