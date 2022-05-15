package char

import "github.com/project-midgard/midgarts/internal/character"

type Point struct {
	X, Y int
}

type Accessory struct {
	Type           character.AttachmentType
	Offset         Point
	PositionFrame  Point
	PositionLayer  Point
	PositionAnchor Point
}

func NewAccessory(elem character.AttachmentType, offset, positionFrame, positionLayer Point) Accessory {
	return Accessory{
		elem,
		offset,
		positionFrame,
		positionLayer,
		Point{},
	}
}

// Anchor takes a set of accessories, calculates the proper anchor
// points for each one of them and returns a new set of accessories
// with the calculated anchor points.
func Anchor(accessories ...Accessory) []Accessory {
	var offset Point

	res := make([]Accessory, 0)
	for _, acc := range accessories {
		var pos Point

		if acc.Type != character.AttachmentBody &&
			acc.Type != character.AttachmentShield {
			pos.X = offset.X - acc.PositionFrame.X
			pos.Y = offset.Y - acc.PositionFrame.Y
		}

		acc.PositionAnchor.X = acc.PositionLayer.X + pos.X
		acc.PositionAnchor.Y = acc.PositionLayer.Y + pos.Y

		offset = Point{acc.PositionFrame.X, acc.PositionFrame.Y}

		res = append(res, acc)
	}

	return res
}
