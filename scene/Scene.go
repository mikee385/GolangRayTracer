package scene

import (
	"github.com/mikee385/GolangRayTracer/color"
	"github.com/mikee385/GolangRayTracer/geometry"
	"math"
)

const Bias = 1.0E-4

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

func (scene *Scene) AddLightSource(light *SceneLight) {
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

func (scene *Scene) AddObject(object SceneObject) {
	var index = len(scene.items)
	scene.items = append(scene.items, internalObject{
		index:   index,
		object:  object,
		isLight: false,
	})
}

func (scene *Scene) Trace(ray geometry.Ray3D, depth uint) TraceResult {
	var nearestDistance float32 = 0.0
	var nearestItem internalObject
	var nearestIntersection = false

	// Find the nearest object that the ray intersects.
	for _, item := range scene.items {
		var currentDistance, hasIntersection = item.object.Intersect(ray)
		if hasIntersection {
			if !nearestIntersection || currentDistance < nearestDistance {
				nearestDistance = currentDistance
				nearestItem = item
				nearestIntersection = true
			}
		}
	}

	// If the ray doesn't hit any objects, return the background color.
	if !nearestIntersection {
		return TraceResult{
			Color:    scene.backgroundColor,
			Distance: 0.0,
		}
	}

	// Get the point where the ray intersects the object.
	var point = ray.Point(nearestDistance)

	// If the ray intersects a light source, simply return the color of the light.
	if nearestItem.isLight {
		return TraceResult{
			Color:    nearestItem.object.Material(point).Color,
			Distance: nearestDistance,
		}
	}

	// Get the surface normal and color at the intersection point.
	var normal = nearestItem.object.Normal(point)
	var surfaceMaterial = nearestItem.object.Material(point)

	var rayVector = ray.Direction.ToVector()
	var normalVector = normal.ToVector()

	// Calculate the color at the intersection point.
	var totalRayColor = color.Black()

	if depth < scene.maxRayDepth {
		// TODO: Add Fresnel effects (?)

		// Calculate the color from the reflected ray.
		var reflection = surfaceMaterial.Reflection
		if reflection > 0.0 {
			var reflectedDirection = rayVector.Sub(normalVector.Scale(2.0 * geometry.Dot(rayVector, normalVector))).ToUnit()
			var nearbyPoint = point.Translate_Dist(reflectedDirection, Bias)
			var reflectedResult = scene.Trace(geometry.NewRay(nearbyPoint, reflectedDirection), depth+1)
			totalRayColor = totalRayColor.Add(reflectedResult.Color.Scale(reflection).Mul(surfaceMaterial.Color))
		}

		// Calculate the color from the refracted ray.
		var refraction = surfaceMaterial.Refraction
		if refraction > 0.0 {
			var n, cosI float32
			if geometry.Dot(rayVector, normalVector) > 0.0 {
				// Internal refraction
				n = surfaceMaterial.RefractiveIndex / scene.refractiveIndex
				cosI = -geometry.Dot(rayVector, normalVector.Neg())
			} else {
				// External refraction
				n = scene.refractiveIndex / surfaceMaterial.RefractiveIndex
				cosI = -geometry.Dot(rayVector, normalVector)
			}

			var cos2T = 1 - n*n*(1-cosI*cosI)
			if cos2T > 0.0 {
				var refractedDirection = rayVector.Scale(n).Add(normalVector.Scale(n*cosI - float32(math.Sqrt(float64(cos2T))))).ToUnit()
				var nearbyPoint = point.Translate_Dist(refractedDirection, Bias)
				var refractedResult = scene.Trace(geometry.NewRay(nearbyPoint, refractedDirection), depth+1)

				// Beer's Law
				var absorbance = surfaceMaterial.Color.Scale(0.15 * -refractedResult.Distance)
				var transparency = color.New(
					float32(math.Exp(float64(absorbance.Red))),
					float32(math.Exp(float64(absorbance.Green))),
					float32(math.Exp(float64(absorbance.Blue))))
				totalRayColor = totalRayColor.Add(refractedResult.Color.Mul(transparency))
			}
		}
	}

	// Calculate the color from each light in the scene.
	for _, lightItem := range scene.lights {
		var light = lightItem.light
		var lightColor = light.Material(point).Color
		var vectorToLight = geometry.NewVector_BetweenPoints(point, light.Center())
		var distanceToLight = vectorToLight.Magnitude()
		var directionToLight = vectorToLight.ToUnit()
		var directionToLightVector = directionToLight.ToVector()

		// Calculate the shading from the light.
		var shade float32 = 1.0
		var nearbyPoint = point.Translate_Dist(directionToLight, Bias)
		var shadowRay = geometry.NewRay(nearbyPoint, directionToLight)
		for _, shadowItem := range scene.items {
			if shadowItem.index != lightItem.index {
				var shadowDistance, hasIntersection = shadowItem.object.Intersect(shadowRay)
				if hasIntersection && shadowDistance < distanceToLight {
					shade = 0.0
					break
				}
			}
		}

		if shade != 0.0 {
			// Calculate the diffusive lighting from the light.
			var diffuse = surfaceMaterial.Diffuse
			if diffuse > 0.0 {
				var percentageOfLight = geometry.Dot(normalVector, directionToLightVector)
				if percentageOfLight > 0.0 {
					totalRayColor = totalRayColor.Add(lightColor.Scale(shade * diffuse * percentageOfLight).Mul(surfaceMaterial.Color))
				}
			}

			// Calculate the specular lighting from the light.
			var specular = surfaceMaterial.Specular
			var shininess = surfaceMaterial.Shininess
			if specular > 0.0 && shininess > 0 {
				var reflectedDirection = directionToLightVector.Sub(normalVector.Scale(2.0 * geometry.Dot(directionToLightVector, normalVector))).ToUnit()
				var percentageOfLight = geometry.Dot(rayVector, reflectedDirection.ToVector())
				if percentageOfLight > 0.0 {
					totalRayColor = totalRayColor.Add(lightColor.Scale(shade * specular * float32(math.Pow(float64(percentageOfLight), float64(shininess)))))
				}
			}
		}
	}

	return TraceResult{
		Color:    totalRayColor,
		Distance: nearestDistance,
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
