package scene

import (
	"fmt"
	"github.com/mikee385/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/table"
	"math"
)

type Camera struct {
	position        geometry.Point3D
	orientation     geometry.Matrix3D
	imageWidth      int
	imageHeight     int
	xMin            float32
	yMax            float32
	dx              float32
	dy              float32
	distanceToPlane float32
}

func NewCamera_FromFOV(imageWidth int, imageHeight int, fieldOfView float32, distanceToPlane float32, position geometry.Point3D, lookAtPoint geometry.Point3D) Camera {
	var yMax = float32(math.Tan(float64(fieldOfView/2.0*geometry.DegreesToRadians))) * distanceToPlane
	var xMin = -yMax * float32(imageWidth) / float32(imageHeight)

	return Camera{
		position:        position,
		orientation:     computeOrientation(position, lookAtPoint),
		imageWidth:      imageWidth,
		imageHeight:     imageHeight,
		xMin:            xMin,
		yMax:            yMax,
		dx:              -2.0 * xMin / float32(imageWidth),
		dy:              2.0 * yMax / float32(imageHeight),
		distanceToPlane: distanceToPlane,
	}
}

func NewCamera_FromDimensions(imageWidth int, imageHeight int, planeWidth float32, planeHeight float32, distanceToPlane float32, position geometry.Point3D, lookAtPoint geometry.Point3D) Camera {
	return Camera{
		position:        position,
		orientation:     computeOrientation(position, lookAtPoint),
		imageWidth:      imageWidth,
		imageHeight:     imageHeight,
		xMin:            -planeWidth / 2.0,
		yMax:            planeHeight / 2.0,
		dx:              planeWidth / float32(imageWidth),
		dy:              planeHeight / float32(imageHeight),
		distanceToPlane: distanceToPlane,
	}
}

func (camera Camera) Position() geometry.Point3D {
	return camera.position
}

func (camera Camera) Orientation() geometry.Matrix3D {
	return camera.orientation
}

func (camera Camera) ImageWidth() int {
	return camera.imageWidth
}

func (camera Camera) ImageHeight() int {
	return camera.imageHeight
}

func (camera Camera) PrimaryRay(row int, column int) geometry.Ray3D {
	var pointInCamera = camera.pixelCenter(row, column)
	var rayDirection = geometry.NewDirection_BetweenPoints(camera.position, camera.convertCameraToWorld(pointInCamera))
	return geometry.NewRay(camera.position, rayDirection)
}

func (camera Camera) SubRays(row int, column int, rays *table.Table) {
	var width = rays.Width()
	if width < 2 {
		panic(fmt.Sprintf("Camera::SubRays: `Width` of `rays` table is too small (%v < %v)", width, 2))
	}
	var height = rays.Height()
	if height < 2 {
		panic(fmt.Sprintf("Camera::SubRays: `Height` of `rays` table is too small (%v < %v)", height, 2))
	}

	var xStep = camera.dx / float32(width-1)
	var yStep = camera.dy / float32(height-1)

	var x0 = camera.xMin + camera.dx*float32(column)
	var y0 = camera.yMax - camera.dy*float32(row)
	var z0 = camera.distanceToPlane

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			var pointInCamera = geometry.NewPoint(x0+float32(column)*xStep, y0-float32(row)*yStep, z0)
			var rayDirection = geometry.NewDirection_BetweenPoints(camera.position, camera.convertCameraToWorld(pointInCamera))
			rays.Set(row, column, geometry.NewRay(camera.position, rayDirection))
		}
	}
}

func (camera Camera) pixelCenter(row int, column int) geometry.Point3D {
	var x = camera.xMin + camera.dx*(float32(column)+0.5)
	var y = camera.yMax - camera.dy*(float32(row)+0.5)
	var z = camera.distanceToPlane

	return geometry.NewPoint(x, y, z)
}

func (camera Camera) convertCameraToWorld(pointInCamera geometry.Point3D) geometry.Point3D {
	return pointInCamera.Rotate(camera.orientation).Translate_Vec(geometry.NewVector_FromPoint(camera.position))
}

func computeOrientation(position geometry.Point3D, lookAtPoint geometry.Point3D) geometry.Matrix3D {
	var z = geometry.NewDirection_BetweenPoints(position, lookAtPoint).ToVector()
	var x = geometry.Cross(geometry.UnitY().ToVector(), z).ToUnit().ToVector()
	var y = geometry.Cross(z, x).ToUnit().ToVector()

	return geometry.NewMatrix(x, y, z)
}
