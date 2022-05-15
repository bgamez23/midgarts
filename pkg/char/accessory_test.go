package char

import (
	"github.com/project-midgard/midgarts/internal/character"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnchorSprites(t *testing.T) {
	// short aliases
	type p = Point
	type a = Accessory

	shadow := NewAccessory(character.AttachmentShadow, p{}, p{}, p{})
	body := NewAccessory(character.AttachmentBody, p{}, p{-1, -62}, p{0, -28})
	head := NewAccessory(character.AttachmentHead, p{-1, -62}, p{0, -56}, p{0, -67})

	res := Anchor([]a{shadow, body, head}...)

	shadowPos := res[0].PositionAnchor
	assert.Equal(t, 0, shadowPos.X, "X should be 0")
	assert.Equal(t, 0, shadowPos.Y, "Y should be 0")

	bodyPos := res[1].PositionAnchor
	assert.Equal(t, 0, bodyPos.X, "bodyPos.X should be 0")
	assert.Equal(t, -28, bodyPos.Y, "bodyPos.Y should be -28")

	headPos := res[2].PositionAnchor
	assert.Equal(t, -1, headPos.X, "headPos.X should be -1")
	assert.Equal(t, -73, headPos.Y, "headPos.Y should be -73")
}
