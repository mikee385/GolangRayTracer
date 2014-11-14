package scene

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
)

type SceneObject interface {
	Intersect(ray geometry.Ray3D) (float32, bool)
	GetNormal(point geometry.Point3D) geometry.Direction3D
	GetMaterial(point geometry.Point3D)
}
