package vendena

// GeometryTypePoint describes the point type.
const GeometryTypePoint = "Point"

// The Geometry model.
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// NewGeometryPoint creates a geometry object of the type "Point".
func (api *API) NewGeometryPoint(latitude float64, longitute float64) Geometry {
	return Geometry{
		Type:        GeometryTypePoint,
		Coordinates: []float64{latitude, longitute},
	}
}
