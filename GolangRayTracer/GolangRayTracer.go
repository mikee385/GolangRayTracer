package main

import (
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/geometry"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/material"
	"github.com/mikee385/GolangRayTracer/GolangRayTracer/scene"
)

func main() {
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
}
