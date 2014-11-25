package geometry

type Matrix3D struct {
	X Vector3D
	Y Vector3D
	Z Vector3D
}

func NewMatrix(x Vector3D, y Vector3D, z Vector3D) Matrix3D {
	return Matrix3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func Identity() Matrix3D {
	return Matrix3D{
		X: Vector3D{X: 1.0, Y: 0.0, Z: 0.0},
		Y: Vector3D{X: 0.0, Y: 1.0, Z: 0.0},
		Z: Vector3D{X: 0.0, Y: 0.0, Z: 1.0},
	}
}

func (matrix Matrix3D) ToOrthonormalBasis() Matrix3D {
	var dirX = matrix.X.ToUnit()
	var vecX = dirX.ToVector()

	var dirY = (matrix.Y.Sub(vecX.Scale(Dot(vecX, matrix.Y)))).ToUnit()
	var vecY = dirY.ToVector()

	var dirZ = (matrix.Z.Sub(vecX.Scale(Dot(vecX, matrix.Z))).Sub(vecY.Scale(Dot(vecY, matrix.Z)))).ToUnit()
	var vecZ = dirZ.ToVector()

	return NewMatrix(vecX, vecY, vecZ)
}

func (matrix Matrix3D) Equals(other Matrix3D) bool {
	return matrix.X.Equals(other.X) &&
		matrix.Y.Equals(other.Y) &&
		matrix.Z.Equals(other.Z)
}

func (matrix Matrix3D) EqualsTol(other Matrix3D, tolerance float32) bool {
	return matrix.X.EqualsTol(other.X, tolerance) &&
		matrix.Y.EqualsTol(other.Y, tolerance) &&
		matrix.Z.EqualsTol(other.Z, tolerance)
}
