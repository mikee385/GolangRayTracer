package scene

import (
	"github.com/mikee385/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/geometry"
)

const Bias = 1.0E-9

type Scene struct {
	backgroundColor color.ColorRGB
	refractiveIndex float32
	maxRayDepth     uint

	items  []internalObject
	lights []internalLight
}

func NewScene(backgroundColor color.ColorRGB, refractiveIndex float32, maxRayDepth uint) Scene {
	return Scene{
		backgroundColor: backgroundColor,
		refractiveIndex: refractiveIndex,
		maxRayDepth:     maxRayDepth,
		items:           make([]internalObject, 0, 4),
		lights:          make([]internalLight, 0, 4),
	}
}

func (scene Scene) AddLightSource(light *SceneLight) {
	var index = len(scene.items)
	scene.items = append(scene.items, internalObject{
		index:   index,
		object:  light,
		isLight: true,
	})
	scene.lights = append(scene.lights, internalLight{
		index: index,
		light: light,
	})
}

func (scene Scene) AddObject(object SceneObject) {
	var index = len(scene.items)
	scene.items = append(scene.items, internalObject{
		index:   index,
		object:  object,
		isLight: false,
	})
}

func (scene Scene) Trace(ray geometry.Ray3D, depth int) TraceResult {
	var pixelColor color.ColorRGB

	if ray.Direction.X() > 0.0 {
		if ray.Direction.Y() > 0.0 {
			pixelColor = color.Red()
		} else {
			pixelColor = color.Green()
		}
	} else {
		if ray.Direction.Y() > 0.0 {
			pixelColor = color.Blue()
		} else {
			pixelColor = color.White()
		}
	}

	return TraceResult{
		Color:    pixelColor,
		Distance: 0.0,
	}
}

type internalObject struct {
	index   int
	object  SceneObject
	isLight bool
}

type internalLight struct {
	index int
	light *SceneLight
}

type TraceResult struct {
	Color    color.ColorRGB
	Distance float32
}
