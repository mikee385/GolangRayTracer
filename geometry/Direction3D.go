package geometry

import "math"

type Direction3D struct {
	direction Vector3D
}

func NewDirection(x float32, y float32, z float32) Direction3D {
	return NewDirection_FromVector(NewVector(x, y, z))
}

func NewDirection_BetweenPoints(from Point3D, to Point3D) Direction3D {
	return NewDirection_FromVector(NewVector_BetweenPoints(from, to))
}

func NewDirection_FromPoint(point Point3D) Direction3D {
	return NewDirection_FromVector(NewVector_FromPoint(point))
}

func NewDirection_FromVector(vector Vector3D) Direction3D {
	var magnitude = vector.Magnitude()
	if magnitude > 0.0 {
		return Direction3D{direction: vector.Scale(1.0 / magnitude)}
	} else {
		return Direction3D{direction: vector}
	}
}

func newDirection_FromNormalizedVector(normalizedVector Vector3D) Direction3D {
	return Direction3D{direction: normalizedVector}
}

func UnitX() Direction3D {
	return Direction3D{direction: Vector3D{X: 1.0, Y: 0.0, Z: 0.0}}
}

func UnitY() Direction3D {
	return Direction3D{direction: Vector3D{X: 0.0, Y: 1.0, Z: 0.0}}
}

func UnitZ() Direction3D {
	return Direction3D{direction: Vector3D{X: 0.0, Y: 0.0, Z: 1.0}}
}

func (direction Direction3D) ToVector() Vector3D {
	return direction.direction
}

func (direction Direction3D) X() float32 {
	return direction.direction.X
}

func (direction Direction3D) Y() float32 {
	return direction.direction.Y
}

func (direction Direction3D) Z() float32 {
	return direction.direction.Z
}

func (direction Direction3D) ToOrthonormalBasis() Matrix3D {
	if math.Abs(float64(direction.direction.X)) >= math.Abs(float64(direction.direction.Y)) && math.Abs(float64(direction.direction.X)) >= math.Abs(float64(direction.direction.Z)) {
		var dirX = direction.direction

		var invXYMagnitude = float32(1.0 / math.Sqrt(float64(dirX.X*dirX.X+dirX.Y*dirX.Y)))
		var dirY = NewVector(
			-dirX.Y*invXYMagnitude,
			dirX.X*invXYMagnitude,
			0.0)

		var dirZ = NewVector(
			-dirX.Z*dirY.Y,
			dirX.Z*dirY.X,
			dirX.X*dirY.Y-dirX.Y*dirY.X)

		return NewMatrix(dirX, dirY, dirZ)
	} else if math.Abs(float64(direction.direction.Y)) >= math.Abs(float64(direction.direction.Z)) {
		var dirY = direction.direction

		var invYZMagnitude = float32(1.0 / math.Sqrt(float64(dirY.Y*dirY.Y+dirY.Z*dirY.Z)))
		var dirZ = NewVector(
			0.0,
			-dirY.Z*invYZMagnitude,
			dirY.Y*invYZMagnitude)

		var dirX = NewVector(
			dirY.Y*dirZ.Z-dirY.Z*dirZ.Y,
			-dirY.X*dirZ.Z,
			dirY.X*dirZ.Y)

		return NewMatrix(dirX, dirY, dirZ)
	} else {
		var dirZ = direction.direction

		var invZXMagnitude = float32(1.0 / math.Sqrt(float64(dirZ.Z*dirZ.Z+dirZ.X*dirZ.X)))
		var dirX = NewVector(
			dirZ.Z*invZXMagnitude,
			0.0,
			-dirZ.X*invZXMagnitude)

		var dirY = NewVector(
			dirZ.Y*dirX.Z,
			dirZ.Z*dirX.X-dirZ.X*dirX.Z,
			-dirZ.Y*dirX.X)

		return NewMatrix(dirX, dirY, dirZ)
	}
}

func (direction Direction3D) Equals(other Direction3D) bool {
	return direction.direction.Equals(other.direction)
}

func (direction Direction3D) EqualsTol(other Direction3D, tolerance float32) bool {
	return direction.direction.EqualsTol(other.direction, tolerance)
}

func (direction Direction3D) Neg() Direction3D {
	return newDirection_FromNormalizedVector(direction.direction.Neg())
}
