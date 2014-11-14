package geometry

type Point3D struct {
	X float32
	Y float32
	Z float32
}

func NewPoint(x float32, y float32, z float32) Point3D {
	return Point3D{X: x, Y: y, Z: z}
}

func NewPoint_FromVector(vector Vector3D) Point3D {
	return Point3D{X: vector.X, Y: vector.Y, Z: vector.Z}
}
