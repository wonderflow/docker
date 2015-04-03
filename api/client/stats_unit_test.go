package client

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/docker/docker/api/types"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/pkg/units"
)

func TestDisplay(t *testing.T) {
	c = &containerStats{
		Name:             "app",
		CPUPercentage:    30.0,
		Memory:           100 * 1024 * 1024.0,
		MemoryLimit:      2048 * 1024 * 1024.0,
		MemoryPercentage: 100.0 / 2048.0 * 100.0,
		NetworkRx:        100 * 1024,
		NetworkTx:        800 * 1024,
		mu:               &sync.RWMutex{},
	}
	`TODO:`
	c.Display()
}
