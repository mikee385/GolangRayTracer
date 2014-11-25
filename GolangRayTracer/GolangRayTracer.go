package main

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/material"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/scene"
)

var EXAMPLE_TO_RUN = 1

func main() {

	if EXAMPLE_TO_RUN == 1 {
		//----------------------------------------------------------------------
		// Scratchapixel Tutorial
		//----------------------------------------------------------------------

		var width = 640
		var height = 480

		var groundSphere = scene.NewSphere(geometry.NewPoint(0.0, -10004.0, 20.0), 10000.0, material.NewBuilder().
			Color(color.New(0.2, 0.2, 0.2)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		var sphere1 = scene.NewSphere(geometry.NewPoint(0.0, 0.0, 20.0), 4.0, material.NewBuilder().
			Color(color.New(1.0, 0.32, 0.36)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.5).
			RefractiveIndex(1.1).
			ToMaterial())

		var sphere2 = scene.NewSphere(geometry.NewPoint(5.0, -1.0, 15.0), 2.0, material.NewBuilder().
			Color(color.New(0.9, 0.76, 0.46)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		var sphere3 = scene.NewSphere(geometry.NewPoint(5.0, 0.0, 25.0), 3.0, material.NewBuilder().
			Color(color.New(0.65, 0.77, 0.97)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		var sphere4 = scene.NewSphere(geometry.NewPoint(-5.5, 0.0, 15.0), 3.0, material.NewBuilder().
			Color(color.New(0.90, 0.90, 0.90)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(1.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

	} else if EXAMPLE_TO_RUN == 2 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 1 & version 2
		//----------------------------------------------------------------------

		var width = 800
		var height = 600

		var groundPlane = scene.NewPlane_FromDVector(4.4, geometry.NewVector(0.0, 1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.3, 0.3)).
			Diffuse(1.0).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			ToMaterial())

		var bigSphere = scene.NewSphere(geometry.NewPoint(1.0, -0.8, 3.0), 2.5, material.NewBuilder().
			Color(color.New(0.7, 0.7, 0.7)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.6).
			ToMaterial())

		var smallSphere = scene.NewSphere(geometry.NewPoint(-5.5, -0.5, 7.0), 2.0, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.1).
			Specular(0.9).
			Shininess(20).
			Reflection(1.0).
			ToMaterial())

	} else if EXAMPLE_TO_RUN == 3 {
		//----------------------------------------------------------------------
		// flipcode Tutorial, version 3
		//----------------------------------------------------------------------

		var width = 800
		var height = 600

		var groundPlane = scene.NewPlane_FromDVector(4.4, geometry.NewVector(0.0, 1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.3, 0.3)).
			Diffuse(1.0).
			Specular(0.8).
			Shininess(20).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		var bigSphere = scene.NewSphere(geometry.NewPoint(2.0, 0.8, 3.0), 2.5, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.2).
			Refraction(0.8).
			RefractiveIndex(1.3).
			ToMaterial())

		var smallSphere = scene.NewSphere(geometry.NewPoint(-5.5, -0.5, 7.0), 2.0, material.NewBuilder().
			Color(color.New(0.7, 0.7, 1.0)).
			Diffuse(0.1).
			Specular(0.8).
			Shininess(20).
			Reflection(0.5).
			Refraction(0.0).
			RefractiveIndex(1.3).
			ToMaterial())

		var extraSphere = scene.NewSphere(geometry.NewPoint(-1.5, -3.8, 1.0), 1.5, material.NewBuilder().
			Color(color.New(1.0, 0.4, 0.4)).
			Diffuse(0.2).
			Specular(0.8).
			Shininess(20).
			Reflection(0.0).
			Refraction(0.8).
			RefractiveIndex(1.5).
			ToMaterial())

		var backPlane = scene.NewPlane_FromDVector(12.0, geometry.NewVector(0.4, 0.0, -1.0), material.NewBuilder().
			Color(color.New(0.5, 0.3, 0.5)).
			Diffuse(0.6).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		var ceilingPlane = scene.NewPlane_FromDVector(7.4, geometry.NewVector(0.0, -1.0, 0.0), material.NewBuilder().
			Color(color.New(0.4, 0.7, 0.7)).
			Diffuse(0.5).
			Specular(0.0).
			Shininess(0).
			Reflection(0.0).
			Refraction(0.0).
			RefractiveIndex(0.0).
			ToMaterial())

		for x := 0; x < 8; x++ {
			for y := 0; y < 7; y++ {
				var gridSphere = scene.NewSphere(geometry.NewPoint(-4.5+float32(x)*1.5, -4.3+float32(y)*1.5, 10.0), 0.3, material.NewBuilder().
					Color(color.New(0.3, 1.0, 0.4)).
					Diffuse(0.6).
					Specular(0.6).
					Shininess(20).
					Reflection(0.0).
					Refraction(0.0).
					RefractiveIndex(0.0).
					ToMaterial())
			}
		}
	}
}
