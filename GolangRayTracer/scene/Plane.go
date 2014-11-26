package scene

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/material"
	"math"
)

type Plane struct {
	origin   geometry.Point3D
	normal   geometry.Direction3D
	material material.Material
}

func NewPlane_FromOriginNormal(origin geometry.Point3D, normal geometry.Direction3D, material material.Material) Plane {
	return Plane{
		origin:   origin,
		normal:   normal,
		material: material,
	}
}

func NewPlane_FromDVector(d float32, vector geometry.Vector3D, material material.Material) Plane {
	return Plane{
		origin:   geometry.NewPoint_FromVector(vector.Scale(-d / geometry.Dot(vector, vector))),
		normal:   geometry.NewDirection_FromVector(vector),
		material: material,
	}
}

func (plane Plane) Origin() geometry.Point3D {
	return plane.origin
}

func (plane Plane) D() float32 {
	return -geometry.Dot(geometry.NewVector_FromPoint(plane.origin), plane.normal.ToVector())
}

func (plane Plane) Intersect(ray geometry.Ray3D) (float32, bool) {
	var denominator = geometry.Dot(ray.Direction.ToVector(), plane.normal.ToVector())
	if float32(math.Abs(float64(denominator))) < geometry.Epsilon {
		return 0.0, false
	}

	var t = geometry.Dot(geometry.NewVector_BetweenPoints(ray.Origin, plane.origin), plane.normal.ToVector()) / denominator
	if t < 0.0 {
		return 0.0, false
	}

	return t, true
}

func (plane Plane) Normal(point geometry.Point3D) geometry.Direction3D {
	return plane.normal
}

func (plane Plane) Material(point geometry.Point3D) material.Material {
	return plane.material
}
