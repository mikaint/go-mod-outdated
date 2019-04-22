package mod_test

import (
	"github.com/psampaz/go-mod-outdated/internal/mod"
	"reflect"
	"testing"
	"time"
)

var mods = []mod.Module{
	mod.Module{
		Path:     "github.com/pk1/pk1",
		Main:     true,
		Indirect: false,
	},
	mod.Module{
		Path:     "github.com/pk2/pk2",
		Main:     false,
		Version:  "v1.0.0",
		Indirect: false,
	},
	mod.Module{
		Path:     "github.com/pk3/pk3",
		Main:     false,
		Version:  "v1.0.0",
		Indirect: true,
	},
	mod.Module{
		Path:     "github.com/pk4/pk4",
		Main:     false,
		Version:  "v1.0.0",
		Indirect: false,
		Update: &mod.Module{
			Version: "v1.1.0",
		},
	},
	mod.Module{
		Path:     "github.com/pk4/pk4",
		Main:     false,
		Version:  "v1.0.0",
		Indirect: true,
		Update: &mod.Module{
			Version: "v1.1.0",
		},
	},
}

func modTime(s string) *time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return &t
}

func TestFilterModules(t *testing.T) {
	want := []mod.Module{
		mod.Module{
			Path:     "github.com/pk2/pk2",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
		},
		mod.Module{
			Path:     "github.com/pk3/pk3",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: true,
		},
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: true,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
	}

	if got := mod.FilterModules(mods, false, false); !reflect.DeepEqual(got, want) {
		t.Errorf("FilterModules() = %v, want %v", got, want)
	}
}

func TestFilterModulesHasUpdate(t *testing.T) {
	want := []mod.Module{
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: true,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
	}

	if got := mod.FilterModules(mods, true, false); !reflect.DeepEqual(got, want) {
		t.Errorf("FilterModules() = %v, want %v", got, want)
	}
}

func TestFilterModulesIsDirect(t *testing.T) {
	want := []mod.Module{
		mod.Module{
			Path:     "github.com/pk2/pk2",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
		},
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
	}

	if got := mod.FilterModules(mods, false, true); !reflect.DeepEqual(got, want) {
		t.Errorf("FilterModules() = %v, want %v", got, want)
	}
}

func TestFilterModulesHasUpdateIsDirect(t *testing.T) {
	want := []mod.Module{
		mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		},
	}

	if got := mod.FilterModules(mods, true, true); !reflect.DeepEqual(got, want) {
		t.Errorf("FilterModules() = %v, want %v", got, want)
	}
}

func TestModule_InvalidTimestamp(t *testing.T) {
	var tests = []struct {
		module           mod.Module
		invalidTimestamp bool
		description      string
	}{
		{mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Time:     modTime("2017-01-01T00:00:00Z"),
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
				Time:    modTime("2018-01-01T00:00:00Z"),
			},
		}, false, "Current version older that the latest version"},
		{mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Time:     modTime("2018-01-01T00:00:00Z"),
			Indirect: false,
			Update: &mod.Module{
				Version: "v1.1.0",
				Time:    modTime("2017-01-01T00:00:00Z"),
			},
		}, true, "Current version newer that latest version"},
		{mod.Module{
			Path:     "github.com/pk4/pk4",
			Main:     false,
			Version:  "v1.0.0",
			Time:     modTime("2018-01-01T00:00:00Z"),
			Indirect: false,
		}, false, "No update"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			i := tt.module.InvalidTimestamp()
			if i != tt.invalidTimestamp {
				t.Errorf("got %v, want %v", i, tt.invalidTimestamp)
			}
		})
	}
}

func TestModule_CurrentVersion(t *testing.T) {
	var tests = []struct {
		module      mod.Module
		version     string
		description string
	}{
		{mod.Module{
			Version:  "v1.0.0",
			Indirect: false,
		}, "v1.0.0", "Current version without replace"},
		{mod.Module{
			Replace: &mod.Module{
				Version: "v0.0.1",
			},
			Version: "v1.0.0",
		}, "v0.0.1", "Current version with replace"}}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			i := tt.module.CurrentVersion()
			if i != tt.version {
				t.Errorf("got %v, want %v", i, tt.version)
			}
		})
	}
}

func TestModule_NewVersion(t *testing.T) {
	var tests = []struct {
		module      mod.Module
		version     string
		description string
	}{
		{mod.Module{
			Version: "v1.0.0",
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		}, "v1.1.0", "New version without replace"},
		{mod.Module{
			Replace: &mod.Module{
				Version: "v0.0.1",
				Update: &mod.Module{
					Version: "v0.0.2",
				},
			},
			Version: "v1.0.0",
			Update: &mod.Module{
				Version: "v1.1.0",
			},
		}, "v0.0.2", "New version with replace"}}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			i := tt.module.NewVersion()
			if i != tt.version {
				t.Errorf("got %v, want %v", i, tt.version)
			}
		})
	}
}

func TestModule_HasNewVersion(t *testing.T) {
	var tests = []struct {
		module      mod.Module
		hasUpdate   bool
		description string
	}{
		{mod.Module{
			Update: &mod.Module{},
		}, true, "New version without replace"},
		{mod.Module{
			Update: nil,
		}, false, "No new version without replace"},
		{mod.Module{
			Replace: &mod.Module{
				Update: &mod.Module{},
			},
			Update: nil,
		}, true, "New version with replace"},
		{mod.Module{
			Replace: &mod.Module{
				Update: nil,
			},
			Update: nil,
		}, false, "No new version with replace"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			i := tt.module.HasUpdate()
			if i != tt.hasUpdate {
				t.Errorf("got %v, want %v", i, tt.hasUpdate)
			}
		})
	}
}
