package geometry

type Ray3D struct {
	Origin    Point3D
	Direction Direction3D
}

func NewRay(origin Point3D, direction Direction3D) Ray3D {
	return Ray3D{Origin: origin, Direction: direction}
}

func (ray Ray3D) Point(distance float32) Point3D {
	return ray.Origin.Translate_Dist(ray.Direction, distance)
}

func (ray Ray3D) Equals(other Ray3D) bool {
	return ray.Origin.Equals(other.Origin) &&
		ray.Direction.Equals(other.Direction)
}

func (ray Ray3D) EqualsTol(other Ray3D, tolerance float32) bool {
	return ray.Origin.EqualsTol(other.Origin, tolerance) &&
		ray.Direction.EqualsTol(other.Direction, tolerance)
}
