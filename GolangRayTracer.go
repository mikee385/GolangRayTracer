package main

import (
	"fmt"
	"github.com/mikee385/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/image"
	"github.com/mikee385/GolangRayTracer/material"
	"github.com/mikee385/GolangRayTracer/scene"
	"github.com/mikee385/GolangRayTracer/table"
	"math"
)

var EXAMPLE_TO_RUN = 1

func main() {

	var camera scene.Camera
	var imageScene scene.Scene

	if EXAMPLE_TO_RUN == 1 {
		//----------------------------------------------------------------------
		// Scratchapixel Tutorial
		//----------------------------------------------------------------------

		var backgroundColor = color.New(2.0, 2.0, 2.0)
		imageScene = scene.NewScene(backgroundColor, 1.0, 5)

		var groundSphere = scene.NewSphere(geometry.NewPoint(0.0, -10004.0, 20.0), 10000.0, material.NewBuilder().
			Color(color.New(0.2, 0.2, 0.2)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(groundSphere)

		var sphere1 = scene.NewSphere(geometry.NewPoint(0.0, 0.0, 20.0), 4.0, material.NewBuilder().
			Color(color.New(1.0, 0.32, 0.36)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.5).
			RefractiveIndex(1.1).
			ToMaterial())
		imageScene.AddObject(sphere1)

		var sphere2 = scene.NewSphere(geometry.NewPoint(5.0, -1.0, 15.0), 2.0, material.NewBuilder().
			Color(color.New(0.9, 0.76, 0.46)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(sphere2)

		var sphere3 = scene.NewSphere(geometry.NewPoint(5.0, 0.0, 25.0), 3.0, material.NewBuilder().
			Color(color.New(0.65, 0.77, 0.97)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(sphere3)

		var sphere4 = scene.NewSphere(geometry.NewPoint(-5.5, 0.0, 15.0), 3.0, material.NewBuilder().
			Color(color.New(0.90, 0.90, 0.90)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(sphere4)

		var lightSource = scene.NewLight(geometry.NewPoint(0.0, 20.0, 30.0), 3.0, color.New(3.0, 3.0, 3.0))
		imageScene.AddLightSource(&lightSource)

		var imageWidth = 640
		var imageHeight = 480
		var fieldOfView float32 = 30.0
		camera = scene.NewCamera_FromFOV(imageWidth, imageHeight, fieldOfView, 1.0, geometry.Origin(), geometry.NewPoint(0.0, 0.0, 1.0))

	} else if EXAMPLE_TO_RUN == 2 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 1 & version 2
		//----------------------------------------------------------------------

		imageScene = scene.NewScene(color.Black(), 1.0, 5)

		var groundPlane = scene.NewPlane_FromDVector(4.4, geometry.NewVector(0.0, 1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.3, 0.3)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			ToMaterial())
		imageScene.AddObject(groundPlane)

		var bigSphere = scene.NewSphere(geometry.NewPoint(1.0, -0.8, 3.0), 2.5, material.NewBuilder().
			Color(color.New(0.7, 0.7, 0.7)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.6).
			ToMaterial())
		imageScene.AddObject(bigSphere)

		var smallSphere = scene.NewSphere(geometry.NewPoint(-5.5, -0.5, 7.0), 2.0, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.1).
			Specular(0.9).
			Shininess(20).
			Reflection(1.0).
			ToMaterial())
		imageScene.AddObject(smallSphere)

		var lightSource1 = scene.NewLight(geometry.NewPoint(0.0, 5.0, 5.0), 0.1, color.New(0.6, 0.6, 0.6))
		imageScene.AddLightSource(&lightSource1)

		var lightSource2 = scene.NewLight(geometry.NewPoint(2.0, 5.0, 1.0), 0.1, color.New(0.7, 0.7, 0.9))
		imageScene.AddLightSource(&lightSource2)

		var imageWidth = 800
		var imageHeight = 600
		camera = scene.NewCamera_FromDimensions(imageWidth, imageHeight, 8.0, 6.0, 5.0, geometry.NewPoint(0.0, 0.0, -5.0), geometry.NewPoint(0.0, 0.0, 1.0))

	} else if EXAMPLE_TO_RUN == 3 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 3
		//----------------------------------------------------------------------

		imageScene = scene.NewScene(color.Black(), 1.0, 5)

		var groundPlane = scene.NewPlane_FromDVector(4.4, geometry.NewVector(0.0, 1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.3, 0.3)).
			Diffuse(1.0).
			Specular(0.8).
			Shininess(20).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(groundPlane)

		var bigSphere = scene.NewSphere(geometry.NewPoint(2.0, 0.8, 3.0), 2.5, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.2).
			Refraction(0.8).
			RefractiveIndex(1.3).
			ToMaterial())
		imageScene.AddObject(bigSphere)

		var smallSphere = scene.NewSphere(geometry.NewPoint(-5.5, -0.5, 7.0), 2.0, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.1).
			Specular(0.8).
			Shininess(20).
			Reflection(0.5).
			Refraction(0.0).
			RefractiveIndex(1.3).
			ToMaterial())
		imageScene.AddObject(smallSphere)

		var lightSource1 = scene.NewLight(geometry.NewPoint(0.0, 5.0, 5.0), 0.1, color.New(0.4, 0.4, 0.4))
		imageScene.AddLightSource(&lightSource1)

		var lightSource2 = scene.NewLight(geometry.NewPoint(-3.0, 5.0, 1.0), 0.1, color.New(0.6, 0.6, 0.8))
		imageScene.AddLightSource(&lightSource2)

		var extraSphere = scene.NewSphere(geometry.NewPoint(-1.5, -3.8, 1.0), 1.5, material.NewBuilder().
			Color(color.New(1.0, 0.4, 0.4)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.0).
			Refraction(0.8).
			RefractiveIndex(1.5).
			ToMaterial())
		imageScene.AddObject(extraSphere)

		var backPlane = scene.NewPlane_FromDVector(12.0, geometry.NewVector(0.4, 0.0, -1.0), material.NewBuilder().
			Color(color.New(0.5, 0.3, 0.5)).
			Diffuse(0.6).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(backPlane)

		var ceilingPlane = scene.NewPlane_FromDVector(7.4, geometry.NewVector(0.0, -1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.7, 0.7)).
			Diffuse(0.5).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())
		imageScene.AddObject(ceilingPlane)

		var gridSpheres = make([]scene.Sphere, 0, 8*7)
		for x := 0; x < 8; x++ {
			for y := 0; y < 7; y++ {
				gridSpheres = append(gridSpheres, scene.NewSphere(geometry.NewPoint(-4.5+float32(x)*1.5, -4.3+float32(y)*1.5, 10.0), 0.3, material.NewBuilder().
					Color(color.New(0.3, 1.0, 0.4)).
					Diffuse(0.6).
					Specular(0.6).
					Shininess(20).
					Reflection(0.0).
					Refraction(0.0).
					RefractiveIndex(0.0).
					ToMaterial()))
			}
		}
		for _, gridSphere := range gridSpheres {
			imageScene.AddObject(gridSphere)
		}

		var imageWidth = 800
		var imageHeight = 600
		camera = scene.NewCamera_FromDimensions(imageWidth, imageHeight, 8.0, 6.0, 5.0, geometry.NewPoint(0.0, 0.0, -5.0), geometry.NewPoint(0.0, 0.0, 1.0))
	}

	var pixelTable = render(&imageScene, &camera)
	var image = image.NewPPMImage(fmt.Sprintf("example%v.ppm", EXAMPLE_TO_RUN))
	var err = image.Save(pixelTable)
	if err != nil {
		panic(err)
	}
}

func render(imageScene *scene.Scene, camera *scene.Camera) *table.Table {
	var width = camera.ImageWidth()
	var height = camera.ImageHeight()
	var pixelTable = table.New(width, height)

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			var ray = camera.PrimaryRay(row, column)
			var result = imageScene.Trace(ray, 0)

			var resultColor = color.New(
				float32(math.Min(float64(result.Color.Red), 1.0)),
				float32(math.Min(float64(result.Color.Green), 1.0)),
				float32(math.Min(float64(result.Color.Blue), 1.0)))

			pixelTable.Set(row, column, resultColor)
		}
	}

	var isEdge = table.New(width, height)
	isEdge.Fill(false)
	for row := 1; row < height-1; row++ {
		for column := 1; column < width-1; column++ {
			var p1 = pixelTable.Get(row-1, column-1).(color.ColorRGB)
			var p2 = pixelTable.Get(row-1, column).(color.ColorRGB)
			var p3 = pixelTable.Get(row-1, column+1).(color.ColorRGB)
			var p4 = pixelTable.Get(row, column-1).(color.ColorRGB)
			var p6 = pixelTable.Get(row, column+1).(color.ColorRGB)
			var p7 = pixelTable.Get(row+1, column-1).(color.ColorRGB)
			var p8 = pixelTable.Get(row+1, column).(color.ColorRGB)
			var p9 = pixelTable.Get(row+1, column+1).(color.ColorRGB)

			var r = calculateGradient(p1.Red, p2.Red, p3.Red, p4.Red, p6.Red, p7.Red, p8.Red, p9.Red)
			var g = calculateGradient(p1.Green, p2.Green, p3.Green, p4.Green, p6.Green, p7.Green, p8.Green, p9.Green)
			var b = calculateGradient(p1.Blue, p2.Blue, p3.Blue, p4.Blue, p6.Blue, p7.Blue, p8.Blue, p9.Blue)

			if r+b+g > 0.5 {
				isEdge.Set(row, column, true)
			} else {
				isEdge.Set(row, column, false)
			}
		}
	}

	var subWidth = 3
	var subHeight = 3
	var subSize = subWidth * subHeight
	var invSubSize = 1.0 / float32(subSize)
	var subRays = table.New(subWidth, subHeight)
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			if isEdge.Get(row, column).(bool) {
				var pixelColor = color.Black()

				camera.SubRays(row, column, &subRays)
				for subRow := 0; subRow < subHeight; subRow++ {
					for subColumn := 0; subColumn < subWidth; subColumn++ {
						var result = imageScene.Trace(subRays.Get(subRow, subColumn).(geometry.Ray3D), 0)

						pixelColor = pixelColor.Add(result.Color.Scale(invSubSize))
					}
				}

				pixelTable.Set(row, column, pixelColor)
			}
		}
	}

	return &pixelTable
}

func calculateGradient(p1 float32, p2 float32, p3 float32, p4 float32, p6 float32, p7 float32, p8 float32, p9 float32) float32 {
	var gx = (p3 + 2*p6 + p9) - (p1 + 2*p4 + p7)
	var gy = (p1 + 2*p2 + p3) - (p7 + 2*p8 + p9)
	return float32(math.Sqrt(float64(gx*gx + gy*gy)))
}
