package peer

import (
	"testing"

	"golang.org/x/mobile/exp/sprite"
)

func TestGetSpriteContainer(t *testing.T) {
	sc := GetSpriteContainer()
	if sc == nil {
		t.Errorf("GetSpriteContainer returned nil. unexpected")
	}
}

type mockGLer struct {
	GLer
}

func (m *mockGLer) NewNode(fn arrangerFunc) *sprite.Node {
	return &sprite.Node{}
}

func (m *mockGLer) SetSubTex(n *sprite.Node, subTex *sprite.SubTex) {
	// nop
}

func (m *mockGLer) AppendChild(n *sprite.Node) {
	// nop
}

func (m *mockGLer) RemoveChild(n *sprite.Node) {
	// nop
}

func TestAddAndRemoveSprite(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gler = &mockGLer{}

	s1 := &Sprite{}
	err := sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	s2 := &Sprite{}
	err = sc.AddSprite(s2, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 2 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	// RemoveSprite marks sprites as "not in use",
	// length of spriteContainer will not be changed
	sc.RemoveSprite(s1)
	if len(sc.spriteNodePairs) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
	sc.RemoveSprite(s2)
	if len(sc.spriteNodePairs) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	// if there're "not in use" sprite in spriteContainer,
	// AddSprite will reuse them. length of spriteContainer will not be changed until the number of
	// sprites don't reach to its capacity.
	err = sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
	err = sc.AddSprite(s2, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	// if there's not "not in use" sprite in spriteContainer,
	// length of spriteContainer will be extended.
	s3 := &Sprite{}
	err = sc.AddSprite(s3, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 3 {
		t.Fatalf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
}

func TestAddSpriteDuplicate(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gler = &mockGLer{}

	s1 := &Sprite{}
	err := sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	// if specified sprite is already added, it will be ignored.
	err = sc.AddSprite(s1, nil, nil)
	if err == nil {
		t.Fatalf("unexpected behaviour. AddSprite should return error for duplicated adding")
	}
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
}

func TestRemoveSpriteDuplicate(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gler = &mockGLer{}

	s1 := &Sprite{}
	err := sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	sc.RemoveSprite(s1)
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	// if specified sprite is already removed, it will be ignored.
	sc.RemoveSprite(s1)
	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
}
