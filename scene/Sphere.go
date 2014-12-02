package scene

import (
	"github.com/mikee385/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/material"
	"math"
)

type Sphere struct {
	center   geometry.Point3D
	radius   float32
	radius2  float32
	material material.Material
}

func NewSphere(center geometry.Point3D, radius float32, material material.Material) Sphere {
	return Sphere{
		center:   center,
		radius:   radius,
		radius2:  radius * radius,
		material: material,
	}
}

func (sphere Sphere) Center() geometry.Point3D {
	return sphere.center
}

func (sphere Sphere) Radius() float32 {
	return sphere.radius
}

func (sphere Sphere) Intersect(ray geometry.Ray3D) (float32, bool) {
	var sphereToRay = geometry.NewVector_BetweenPoints(ray.Origin, sphere.center)
	var b = geometry.Dot(sphereToRay, ray.Direction.ToVector())
	if b < 0.0 {
		return 0.0, false
	}

	var d2 = geometry.Dot(sphereToRay, sphereToRay) - b*b
	if d2 > sphere.radius2 {
		return 0.0, false
	}

	var c = float32(math.Sqrt(float64(sphere.radius2 - d2)))
	var t = b - c
	if t < 0 {
		t = b + c
	}

	return t, true
}

func (sphere Sphere) Normal(point geometry.Point3D) geometry.Direction3D {
	return geometry.NewDirection_BetweenPoints(sphere.center, point)
}

func (sphere Sphere) Material(point geometry.Point3D) material.Material {
	return sphere.material
}
