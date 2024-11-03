package engine

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Engine struct {
	mx sync.Mutex
	wg *sync.WaitGroup

	Network      *Network
	LoadedScenes map[string]*Scene
	tickRate     time.Duration

	Debug bool
}

func NewEngine(tickRate time.Duration) *Engine {
	e := new(Engine)

	e.tickRate = tickRate
	e.wg = new(sync.WaitGroup)
	e.LoadedScenes = make(map[string]*Scene)
	e.Debug = false

	return e
}

func (e *Engine) Run(ctx context.Context) {
	ticker := time.NewTicker(e.tickRate)
	defer ticker.Stop()

	dt := e.tickRate.Seconds()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e.Update(dt)
		}
	}
}

func (e *Engine) Update(dt float64) {
	e.mx.Lock()
	defer e.mx.Unlock()

	if e.Debug {
		log.Println("=========ENGINE UPDATE START==========")
		defer log.Println("=========ENGINE UPDATE FINISH=========")
	}

	e.Network.Update()

	sceneLen := len(e.LoadedScenes)
	if sceneLen == 0 {
		if e.Debug {
			log.Println("NO ACTIVE SCENES")
		}

		return
	}

	e.wg.Add(sceneLen)
	for i := range e.LoadedScenes {
		go updateSceneAsync(e.LoadedScenes[i], dt, e.wg)
	}

	e.wg.Wait()
}

func (e *Engine) LoadScene(scene Scene) {
	e.mx.Lock()
	defer e.mx.Unlock()

	if e.Debug {
		log.Println("Loading scene:", scene.Name)
		defer log.Println("Scene loaded:", scene.Name)
	}

	scene.Engine = e

	scene.Load()

	//check if scene already exists
	suffix := 1

	for {
		prefixedName := fmt.Sprintf("%s_%d", scene.Name, suffix)

		if _, ok := e.LoadedScenes[prefixedName]; ok {
			suffix++
			continue
		}

		scene.Name = prefixedName
		break
	}

	e.LoadedScenes[scene.Name] = &scene
}

func (e *Engine) UnloadScene(name string) {
	e.mx.Lock()
	defer e.mx.Unlock()

	if e.Debug {
		log.Println("Unloading scene: ", name)
		defer log.Println("Scene unloaded: ", name)
	}

	// check if scene exists
	if _, ok := e.LoadedScenes[name]; !ok {
		return
	}

	e.LoadedScenes[name].Unload()

	delete(e.LoadedScenes, name)
}

func (e *Engine) UnloadAllScenes() {
	for i := range e.LoadedScenes {
		e.UnloadScene(i)
	}
}

func (e *Engine) SetDebug(mode bool) *Engine {
	e.Debug = mode

	return e
}

func updateSceneAsync(scene *Scene, dt float64, wg *sync.WaitGroup) {
	defer wg.Done()
	scene.Update(dt)
}
