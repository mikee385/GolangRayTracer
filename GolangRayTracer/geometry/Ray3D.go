package geometry

type Ray3D struct {
	Origin    Point3D
	Direction Direction3D
}

func NewRay(origin Point3D, direction Direction3D) Ray3D {
	return Ray3D{Origin: origin, Direction: direction}
}
