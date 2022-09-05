package wrapper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/pkg/plugin"
)

func TestWrapperBuilderInitialState(t *testing.T) {
	w := NewPluginWrapper(plugin.NewDummyPluginRunner())

	assert.False(t, w.IsPrepared())
	assert.False(t, w.IsProcessed())
	assert.False(t, w.HasError())
}

func TestWrapperBuilderPreparedState(t *testing.T) {
	os.Args = []string{"/wrapper", "ls"}

	w := NewPluginWrapper(plugin.NewDummyPluginRunner()).Prepare()

	assert.True(t, w.IsPrepared())
	assert.False(t, w.IsProcessed())
	assert.False(t, w.HasError())
}

func TestWrapperBuilderAlreadyPreparedState(t *testing.T) {
	os.Args = []string{"/wrapper", "ls"}

	w := NewPluginWrapper(plugin.NewDummyPluginRunner()).Prepare().Prepare()

	assert.True(t, w.IsPrepared())
	assert.False(t, w.IsProcessed())
	assert.False(t, w.HasError())
}

func TestWrapperBuilderProcessedState(t *testing.T) {
	os.Args = []string{"/wrapper", "ls"}

	w := NewPluginWrapper(plugin.NewDummyPluginRunner()).Run()

	assert.False(t, w.IsPrepared())
	assert.False(t, w.IsProcessed())
	assert.False(t, w.HasError())
}

func TestWrapperBuilder(t *testing.T) {
	os.Args = []string{"/wrapper", "ls"}

	w := NewPluginWrapper(plugin.NewDummyPluginRunner()).Prepare().Run()

	assert.True(t, w.IsPrepared())
	assert.True(t, w.IsProcessed())
	assert.False(t, w.HasError())
}
