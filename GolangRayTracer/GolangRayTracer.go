package main

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/material"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/scene"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/table"
)

var EXAMPLE_TO_RUN = 1

func main() {

	if EXAMPLE_TO_RUN == 1 {
		//----------------------------------------------------------------------
		// Scratchapixel Tutorial
		//----------------------------------------------------------------------

		var backgroundColor = color.New(2.0, 2.0, 2.0)
		var imageScene = scene.NewScene(backgroundColor, 1.0, 5)

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
		var camera = scene.NewCamera_FromFOV(imageWidth, imageHeight, fieldOfView, 1.0, geometry.Origin(), geometry.NewPoint(0.0, 0.0, 1.0))
		var imageTable = render(&imageScene, &camera)

	} else if EXAMPLE_TO_RUN == 2 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 1 & version 2
		//----------------------------------------------------------------------

		var imageScene = scene.NewScene(color.Black(), 1.0, 5)

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
		var camera = scene.NewCamera_FromDimensions(imageWidth, imageHeight, 8.0, 6.0, 5.0, geometry.NewPoint(0.0, 0.0, -5.0), geometry.NewPoint(0.0, 0.0, 1.0))
		var imageTable = render(&imageScene, &camera)

	} else if EXAMPLE_TO_RUN == 3 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 3
		//----------------------------------------------------------------------

		var imageScene = scene.NewScene(color.Black(), 1.0, 5)

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
		var camera = scene.NewCamera_FromDimensions(imageWidth, imageHeight, 8.0, 6.0, 5.0, geometry.NewPoint(0.0, 0.0, -5.0), geometry.NewPoint(0.0, 0.0, 1.0))
		var imageTable = render(&imageScene, &camera)
	}
}

func render(imageScene *scene.Scene, camera *scene.Camera) *table.Table {
	return nil
}
