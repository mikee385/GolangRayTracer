package scene

import (
	"github.com/mikee385/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/material"
)

type SceneLight struct {
	sphere Sphere
}

func NewLight(center geometry.Point3D, radius float32, color color.ColorRGB) SceneLight {
	return SceneLight{sphere: NewSphere(center, radius, material.NewMaterial(color))}
}

func (light SceneLight) Center() geometry.Point3D {
	return light.sphere.Center()
}

func (light SceneLight) Radius() float32 {
	return light.sphere.Radius()
}

func (light SceneLight) Intersect(ray geometry.Ray3D) (float32, bool) {
	return light.sphere.Intersect(ray)
}

func (light SceneLight) Normal(point geometry.Point3D) geometry.Direction3D {
	return light.sphere.Normal(point)
}

func (light SceneLight) Material(point geometry.Point3D) material.Material {
	return light.sphere.Material(point)
}
