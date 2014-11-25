package scene

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/material"
)

type SceneObject interface {
	Intersect(ray geometry.Ray3D) (float32, bool)
	Normal(point geometry.Point3D) geometry.Direction3D
	Material(point geometry.Point3D) material.Material
}
