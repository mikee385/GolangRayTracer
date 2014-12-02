package material

import (
	"github.com/mikee385/GolangRayTracer/color"
)

type Material struct {
	Color           color.ColorRGB
	Diffuse         float32
	Specular        float32
	Shininess       int
	Reflection      float32
	Refraction      float32
	RefractiveIndex float32
}

func NewMaterial(color color.ColorRGB) Material {
	return Material{
		Color:           color,
		Diffuse:         1.0,
		Specular:        0.0,
		Shininess:       0,
		Reflection:      0.0,
		Refraction:      0.0,
		RefractiveIndex: 0.0,
	}
}

type MaterialBuilder interface {
	Color(color color.ColorRGB) MaterialBuilder
	Diffuse(diffuse float32) MaterialBuilder
	Specular(specular float32) MaterialBuilder
	Shininess(shininess int) MaterialBuilder
	Reflection(reflection float32) MaterialBuilder
	Refraction(refraction float32) MaterialBuilder
	RefractiveIndex(refractiveIndex float32) MaterialBuilder
	ToMaterial() Material
}

type materialBuilder struct {
	color           color.ColorRGB
	diffuse         float32
	specular        float32
	shininess       int
	reflection      float32
	refraction      float32
	refractiveIndex float32
}

func NewBuilder() MaterialBuilder {
	return &materialBuilder{
		color:           color.White(),
		diffuse:         1.0,
		specular:        0.0,
		shininess:       0,
		reflection:      0.0,
		refraction:      0.0,
		refractiveIndex: 0.0,
	}
}

func (builder *materialBuilder) Color(color color.ColorRGB) MaterialBuilder {
	builder.color = color
	return builder
}

func (builder *materialBuilder) Diffuse(diffuse float32) MaterialBuilder {
	builder.diffuse = diffuse
	return builder
}

func (builder *materialBuilder) Specular(specular float32) MaterialBuilder {
	builder.specular = specular
	return builder
}

func (builder *materialBuilder) Shininess(shininess int) MaterialBuilder {
	builder.shininess = shininess
	return builder
}

func (builder *materialBuilder) Reflection(reflection float32) MaterialBuilder {
	builder.reflection = reflection
	return builder
}

func (builder *materialBuilder) Refraction(refraction float32) MaterialBuilder {
	builder.refraction = refraction
	return builder
}

func (builder *materialBuilder) RefractiveIndex(refractiveIndex float32) MaterialBuilder {
	builder.refractiveIndex = refractiveIndex
	return builder
}

func (builder materialBuilder) ToMaterial() Material {
	return Material{
		Color:           builder.color,
		Diffuse:         builder.diffuse,
		Specular:        builder.specular,
		Shininess:       builder.shininess,
		Reflection:      builder.reflection,
		Refraction:      builder.refraction,
		RefractiveIndex: builder.refractiveIndex,
	}
}
