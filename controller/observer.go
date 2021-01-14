package controller

import (
	"sync/atomic"

	blaster "github.com/ddrager/go-pi-blaster"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"observer/config"
	"observer/engine"
	"observer/observerconfig"
)

// Observer represents the robot's controller.
type Observer struct {
	name string
	cfg  *observerconfig.ObserverConfig

	// devices
	adaptor *raspi.Adaptor
	camera  *opencv.CameraDriver
	window  *opencv.WindowDriver

	// servos
	blaster blaster.Blaster
	servos  *engine.Servos

	work        func()
	activeFrame *atomic.Value
}

// NewObserver constructs a new Observer controller.
func NewObserver(name string, cfg *config.Config) *Observer {
	o := &Observer{
		name: name,
		cfg:  observerconfig.LoadObserverConfig(cfg),

		adaptor: raspi.NewAdaptor(),
		camera:  opencv.NewCameraDriver(cfg.Camera.Source),
		window:  nil, // optional

		blaster: blaster.Blaster{},
		servos:  nil,

		work:        func() {},
		activeFrame: &atomic.Value{},
	}

	o.initWindow()
	o.initServos()
	o.initCamera()
	o.checkFrequency()

	return o
}
