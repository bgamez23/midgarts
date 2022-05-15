package system_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/project-midgard/midgarts/internal/character"
	"github.com/project-midgard/midgarts/internal/character/jobspriteid"
	"github.com/project-midgard/midgarts/internal/component"
	"github.com/project-midgard/midgarts/internal/entity"
	"github.com/project-midgard/midgarts/internal/fileformat/grf"
	"github.com/project-midgard/midgarts/internal/graphic"
	"github.com/project-midgard/midgarts/internal/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

type nilTextureProvider struct {
}

func (n nilTextureProvider) NewTextureFromRGBA(rgba *graphic.UniqueRGBA) (tex *graphic.Texture, err error) {
	return nil, nil
}

func TestRenderAttachment(t *testing.T) {
	GRF, err := grf.Load("./../../assets/grf/data.grf")
	assert.NoError(t, err)
	assert.NotNil(t, GRF)

	sys := system.NewCharacterRenderSystem(GRF, &nilTextureProvider{})

	char := entity.NewCharacter(character.Female, jobspriteid.Blacksmith, 23)
	cmp, err := component.NewCharacterAttachmentComponent(GRF, component.CharacterAttachmentComponentConfig{
		Gender:           char.Gender,
		JobSpriteID:      char.JobSpriteID,
		HeadIndex:        char.HeadIndex,
		EnableShield:     char.HasShield,
		ShieldSpriteName: char.ShieldSpriteName,
	})

	assert.NoError(t, err)

	char.SetCharacterAttachmentComponent(cmp)

	offset := [2]float32{0, 0}

	sys.RenderAttachment(0, char, character.AttachmentShadow, &offset)
	spew.Dump(offset)

	sys.RenderAttachment(0, char, character.AttachmentBody, &offset)
	spew.Dump(offset)

	sys.RenderAttachment(0, char, character.AttachmentHead, &offset)
	spew.Dump(offset)
}
