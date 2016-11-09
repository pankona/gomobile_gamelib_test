package scene

import (
	"image"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Stage1 represents scene of Stage1.
type Stage1 struct {
	models        models
	views         views
	ball          Ball
	obstacle      [2]Obstacle
	background    [2]Background
	isTouching    bool
	remainingLife int
	life          [3]Life
}

// Life represents view part of remaining life
type Life struct {
	simra.Sprite
}

func (life *Life) getPosition() (x float32, y float32) {
	x, y = life.X, life.Y
	return
}

func (life *Life) setPosition(x float32, y float32) {
	life.X, life.Y = x, y
}

func (life *Life) setSpeed(s float64) {
}

func (life *Life) getSpeed() float64 {
	return 0
}

func (life *Life) setDirection(radian float64) {
}
func (life *Life) getDirection() float64 {
	return 0
}

func (life *Life) move() {
}

const (
	remainingLifeAtStart = 3
)

// Initialize initializes Stage1 scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (scene *Stage1) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(scene)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(Stage1)

	scene.resetPosition()
	scene.setupSprites()
	scene.registerViews()
	scene.registerModels()
	scene.remainingLife = remainingLifeAtStart

	simra.GetInstance().AddCollisionListener(&scene.ball, &scene.obstacle[0], &scene.models)
	simra.GetInstance().AddCollisionListener(&scene.ball, &scene.obstacle[1], &scene.models)

	simra.LogDebug("[OUT]")
}

// OnTouchBegin is called when Stage1 scene is Touched.
func (scene *Stage1) OnTouchBegin(x, y float32) {
	scene.isTouching = true
}

// OnTouchMove is called when Stage1 scene is Touched and moved.
func (scene *Stage1) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Stage1 scene is Touched and it is released.
func (scene *Stage1) OnTouchEnd(x, y float32) {
	scene.isTouching = false
}

func (scene *Stage1) resetPosition() {
	// set size of background
	scene.background[0].W = config.ScreenWidth + 1
	scene.background[0].H = config.ScreenHeight

	// put center of screen
	scene.background[0].X = config.ScreenWidth / 2
	scene.background[0].Y = config.ScreenHeight / 2

	// set size of background
	scene.background[1].W = config.ScreenWidth + 1
	scene.background[1].H = config.ScreenHeight

	// put out of screen
	scene.background[1].X = config.ScreenWidth/2 + (config.ScreenWidth)
	scene.background[1].Y = config.ScreenHeight / 2

	// set size of ball
	scene.ball.W = float32(48)
	scene.ball.H = float32(48)

	// put center of screen at start
	scene.ball.X = config.ScreenWidth / 2
	scene.ball.Y = config.ScreenHeight / 2

	// set size of obstacle
	scene.obstacle[0].W = 50
	scene.obstacle[0].H = 100
	scene.obstacle[1].W = 50
	scene.obstacle[1].H = 100

	// put center/upper side of screen
	scene.obstacle[0].X = config.ScreenWidth + config.ScreenWidth/2
	scene.obstacle[0].Y = config.ScreenHeight / 3 * 2

	// put center/lower side of screen
	scene.obstacle[1].X = config.ScreenWidth + config.ScreenWidth/2
	scene.obstacle[1].Y = config.ScreenHeight / 3 * 1

	scene.life[0].X = 48
	scene.life[0].Y = 30
	scene.life[0].W = float32(48)
	scene.life[0].H = float32(48)
	scene.life[1].X = 48 * 2
	scene.life[1].Y = 30
	scene.life[1].W = float32(48)
	scene.life[1].H = float32(48)
	scene.life[2].X = 48 * 3
	scene.life[2].Y = 30
	scene.life[2].W = float32(48)
	scene.life[2].H = float32(48)
}

func (scene *Stage1) setupSprites() {

	simra.GetInstance().AddSprite("bg.png",
		image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight),
		&scene.background[0].Sprite)

	simra.GetInstance().AddSprite("bg.png",
		image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight),
		&scene.background[1].Sprite)

	simra.GetInstance().AddSprite("ball.png",
		image.Rect(0, 0, int(scene.ball.W), int(scene.ball.H)),
		&scene.ball.Sprite)

	simra.GetInstance().AddSprite("obstacle.png",
		image.Rect(0, 0, 100, 100),
		&scene.obstacle[0].Sprite)

	simra.GetInstance().AddSprite("obstacle.png",
		image.Rect(0, 0, 100, 100),
		&scene.obstacle[1].Sprite)

	simra.GetInstance().AddSprite("heart.png",
		image.Rect(0, 0, 384, 384),
		&scene.life[0].Sprite)

	simra.GetInstance().AddSprite("heart.png",
		image.Rect(0, 0, 384, 384),
		&scene.life[1].Sprite)

	simra.GetInstance().AddSprite("heart.png",
		image.Rect(0, 0, 384, 384),
		&scene.life[2].Sprite)
}

func (scene *Stage1) registerViews() {
	scene.views.registerBall(&scene.ball)
	scene.views.addEventListener(scene)
}

func (scene *Stage1) onFinishDead() {
	scene.resetPosition()
	scene.views.restart()
	scene.models.restart()
}

func (scene *Stage1) registerModels() {
	scene.models.registerBall(&scene.ball)
	scene.models.registerObstacle(&scene.obstacle[0], 0)
	scene.models.registerObstacle(&scene.obstacle[1], 1)
	scene.models.registerBackground(&scene.background[0], 0)
	scene.models.registerBackground(&scene.background[1], 1)
	scene.models.addEventListener(&scene.views)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (scene *Stage1) Drive() {
	scene.models.Progress(scene.isTouching)
	scene.views.Progress(scene.isTouching)
}
