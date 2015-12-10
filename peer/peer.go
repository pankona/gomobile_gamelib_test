package peer

import (
	"image"
	"log"
	"time"

	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

var self *Peer

var startTime = time.Now()

type Peer struct {
	glctx  gl.Context
	images *glutil.Images
	fps    *debug.FPS
	eng    sprite.Engine
	scene  *sprite.Node
	sz     size.Event
}

func GetInstance() *Peer {
	if self == nil {
		self = &Peer{}
	}
	return self
}

func (self *Peer) Initialize(in_glctx gl.Context) {
	self.glctx = in_glctx

	// transparency of png
	self.glctx.Enable(gl.BLEND)
	self.glctx.BlendEquation(gl.FUNC_ADD)
	self.glctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	onStart(self.glctx)
}

func onStart(glctx gl.Context) {
	self.images = glutil.NewImages(glctx)
	self.fps = debug.NewFPS(self.images)
	loadScene()
}

func loadScene() {
	if self.eng != nil {
		self.eng.Release()
	}
	self.eng = glsprite.Engine(self.images)
	self.scene = &sprite.Node{}
	self.eng.Register(self.scene)
	self.eng.SetTransform(self.scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	self.eng.Register(n)
	self.scene.AppendChild(n)
	return n
}

func loadTextures(assetName string, rect image.Rectangle) sprite.SubTex {

	a, err := asset.Open(assetName)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := self.eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	return sprite.SubTex{t, rect}
}

func (self *Peer) SetScreenSize(in_sz size.Event) {
	self.sz = in_sz
}

func (self *Peer) Stop() {
	self.eng.Release()
	self.fps.Release()
	self.images.Release()
	self.glctx = nil
}

func (self *Peer) Update() {
	if self.glctx == nil {
		return
	}
	self.glctx.ClearColor(1, 1, 1, 1) // white background
	self.glctx.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startTime) * 60 / time.Second)
	self.eng.Render(self.scene, now, self.sz)
	self.fps.Draw(self.sz)
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }