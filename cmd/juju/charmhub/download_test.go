// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charmhub

import (
	"net/url"
	"sort"

	"github.com/golang/mock/gomock"
	"github.com/juju/cmd/cmdtesting"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/charmhub/transport"
	"github.com/juju/juju/cmd/juju/charmhub/mocks"
	"github.com/juju/juju/cmd/modelcmd"
	corecharm "github.com/juju/juju/core/charm"
	"github.com/juju/juju/jujuclient"
	"github.com/juju/juju/jujuclient/jujuclienttesting"
	"github.com/juju/juju/testing"
)

type downloadSuite struct {
	testing.FakeJujuXDGDataHomeSuite
	store *jujuclient.MemStore

	charmHubClient *mocks.MockCharmHubClient
	modelConfigAPI *mocks.MockModelConfigGetter
}

var _ = gc.Suite(&downloadSuite{})

func (s *downloadSuite) SetUpTest(c *gc.C) {
	s.FakeJujuXDGDataHomeSuite.SetUpTest(c)
	s.store = jujuclienttesting.MinimalStore()
}

func (s *downloadSuite) TestInitNoArgs(c *gc.C) {
	command := &downloadCommand{}
	err := command.Init([]string{})
	c.Assert(err, gc.ErrorMatches, "expected a charm or bundle name")
}

func (s *downloadSuite) TestInitSuccess(c *gc.C) {
	command := &downloadCommand{}
	err := command.Init([]string{"test"})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *downloadSuite) TestRun(c *gc.C) {
	defer s.setUpMocks(c).Finish()

	url := "http://example.org/"

	s.expectModelGet(url)
	s.expectInfo(url)
	s.expectDownload(c, url)

	command := &downloadCommand{
		modelConfigAPI: s.modelConfigAPI,
		charmHubClient: s.charmHubClient,
	}
	command.SetClientStore(s.store)
	cmd := modelcmd.Wrap(command, modelcmd.WrapSkipModelInit)
	err := cmdtesting.InitCommand(cmd, []string{"test"})
	c.Assert(err, jc.ErrorIsNil)

	ctx := commandContextForTest(c)
	err = cmd.Run(ctx)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *downloadSuite) TestRunWithStdout(c *gc.C) {
	defer s.setUpMocks(c).Finish()

	url := "http://example.org/"

	s.expectModelGet(url)
	s.expectInfo(url)
	s.expectDownload(c, url)

	command := &downloadCommand{
		modelConfigAPI: s.modelConfigAPI,
		charmHubClient: s.charmHubClient,
	}
	command.SetClientStore(s.store)
	cmd := modelcmd.Wrap(command, modelcmd.WrapSkipModelInit)
	err := cmdtesting.InitCommand(cmd, []string{"test", "-"})
	c.Assert(err, jc.ErrorIsNil)

	ctx := commandContextForTest(c)
	err = cmd.Run(ctx)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *downloadSuite) TestRunWithCustomCharmHubURL(c *gc.C) {
	defer s.setUpMocks(c).Finish()

	url := "http://example.org/"

	s.expectInfo(url)
	s.expectDownload(c, url)

	command := &downloadCommand{
		modelConfigAPI: s.modelConfigAPI,
		charmHubClient: s.charmHubClient,
	}
	err := cmdtesting.InitCommand(command, []string{"--charm-hub-url=" + url, "test"})
	c.Assert(err, jc.ErrorIsNil)

	ctx := commandContextForTest(c)
	err = command.Run(ctx)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *downloadSuite) TestRunWithCustomInvalidCharmHubURL(c *gc.C) {
	defer s.setUpMocks(c).Finish()

	url := "meshuggah"

	command := &downloadCommand{
		modelConfigAPI: s.modelConfigAPI,
		charmHubClient: s.charmHubClient,
	}
	err := cmdtesting.InitCommand(command, []string{"--charm-hub-url=" + url, "test"})
	c.Assert(err, gc.ErrorMatches, `unexpected charm-hub-url: parse "meshuggah": invalid URI for request`)
}

func (s *downloadSuite) TestRunWithInvalidStdout(c *gc.C) {
	defer s.setUpMocks(c).Finish()

	command := &downloadCommand{
		modelConfigAPI: s.modelConfigAPI,
		charmHubClient: s.charmHubClient,
	}
	err := cmdtesting.InitCommand(command, []string{"test", "_"})
	c.Assert(err, gc.ErrorMatches, `expected a charm or bundle name, followed by hyphen to pipe to stdout`)
}

func (s *downloadSuite) setUpMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.charmHubClient = mocks.NewMockCharmHubClient(ctrl)
	s.modelConfigAPI = mocks.NewMockModelConfigGetter(ctrl)
	return ctrl
}

func (s *downloadSuite) expectModelGet(charmHubURL string) {
	s.modelConfigAPI.EXPECT().ModelGet().Return(map[string]interface{}{
		"type":          "my-type",
		"name":          "my-name",
		"uuid":          "deadbeef-0bad-400d-8000-4b1d0d06f00d",
		"charm-hub-url": charmHubURL,
	}, nil)
}

func (s *downloadSuite) expectInfo(charmHubURL string) {
	s.charmHubClient.EXPECT().Info(gomock.Any(), "test").Return(transport.InfoResponse{
		Type: "charm",
		Name: "test",
		ChannelMap: []transport.ChannelMap{{
			Channel: transport.Channel{
				Name:  "a",
				Track: "latest",
				Risk:  "stable",
				Platform: transport.Platform{
					Series: "xenial",
				},
			},
			Revision: transport.Revision{
				Revision: 1,
				Download: transport.Download{
					URL: charmHubURL,
				},
			},
		}},
	}, nil)
}

func (s *downloadSuite) expectDownload(c *gc.C, charmHubURL string) {
	resourceURL, err := url.Parse(charmHubURL)
	c.Assert(err, jc.ErrorIsNil)
	s.charmHubClient.EXPECT().Download(gomock.Any(), resourceURL, "test.charm").Return(nil)
}

func (s *downloadSuite) TestLocateRevisionByChannel(c *gc.C) {
	command := &downloadCommand{
		orderedSeries: []string{"focal", "bionic"},
	}
	revision, found := command.locateRevisionByChannel([]transport.ChannelMap{{
		Channel: transport.Channel{
			Name:  "a",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "xenial",
			},
		},
		Revision: transport.Revision{
			Revision: 1,
		},
	}, {
		Channel: transport.Channel{
			Name:  "b",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "bionic",
			},
		},
		Revision: transport.Revision{
			Revision: 2,
		},
	}, {
		Channel: transport.Channel{
			Name:  "c",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "focal",
			},
		},
		Revision: transport.Revision{
			Revision: 3,
		},
	}, {
		Channel: transport.Channel{
			Name:  "d",
			Track: "2.0",
			Risk:  "edge",
			Platform: transport.Platform{
				Series: "focal",
			},
		},
		Revision: transport.Revision{
			Revision: 4,
		},
	}}, corecharm.MustParseChannel("latest/stable"))

	c.Assert(found, jc.IsTrue)
	c.Assert(revision.Revision, gc.Equals, 3)
}

func (s *downloadSuite) TestLocateRevisionByChannelAndSeries(c *gc.C) {
	command := &downloadCommand{
		orderedSeries: []string{"focal", "bionic"},
	}
	revision, found := command.locateRevisionByChannelAndSeries([]transport.ChannelMap{{
		Channel: transport.Channel{
			Name:  "a",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "xenial",
			},
		},
		Revision: transport.Revision{
			Revision: 1,
		},
	}, {
		Channel: transport.Channel{
			Name:  "b",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "bionic",
			},
		},
		Revision: transport.Revision{
			Revision: 2,
		},
	}, {
		Channel: transport.Channel{
			Name:  "c",
			Track: "latest",
			Risk:  "stable",
			Platform: transport.Platform{
				Series: "focal",
			},
		},
		Revision: transport.Revision{
			Revision: 3,
		},
	}, {
		Channel: transport.Channel{
			Name:  "d",
			Track: "2.0",
			Risk:  "edge",
			Platform: transport.Platform{
				Series: "focal",
			},
		},
		Revision: transport.Revision{
			Revision: 4,
		},
	}}, corecharm.MustParseChannel("latest/stable"), "focal")

	c.Assert(found, jc.IsTrue)
	c.Assert(revision.Revision, gc.Equals, 3)
}

func (s *downloadSuite) TestLocateRevisionByChannelMap(c *gc.C) {
	tests := []struct {
		InputTrack string
		InputRisk  string
		Channel    string
	}{{
		InputTrack: "",
		InputRisk:  "stable",
		Channel:    "latest/stable",
	}, {
		InputTrack: "latest",
		InputRisk:  "stable",
		Channel:    "latest/stable",
	}, {
		InputTrack: "2.0",
		InputRisk:  "",
		Channel:    "2.0/stable",
	}, {
		InputTrack: "",
		InputRisk:  "edge",
		Channel:    "edge",
	}}
	for i, test := range tests {
		c.Logf("test %d", i)

		revision, found := locateRevisionByChannelMap(transport.ChannelMap{
			Channel: transport.Channel{
				Track: test.InputTrack,
				Risk:  test.InputRisk,
			},
			Revision: transport.Revision{
				Revision: 1,
			},
		}, corecharm.MustParseChannel(test.Channel))

		c.Assert(found, jc.IsTrue)
		c.Assert(revision.Revision, gc.Equals, 1)
	}
}

func (s *downloadSuite) TestChannelMapSort(c *gc.C) {
	series := channelMapBySeries{
		channelMap: []transport.ChannelMap{{
			Channel: transport.Channel{
				Name: "a",
				Platform: transport.Platform{
					Series: "xenial",
				},
			},
		}, {
			Channel: transport.Channel{
				Name: "b",
				Platform: transport.Platform{
					Series: "bionic",
				},
			},
		}, {
			Channel: transport.Channel{
				Name: "c",
				Platform: transport.Platform{
					Series: "focal",
				},
			},
		}},
		series: []string{"focal", "bionic"},
	}

	sort.Sort(series)

	names := make([]string, 3)
	for k, v := range series.channelMap {
		names[k] = v.Channel.Name
	}

	c.Assert(names, gc.DeepEquals, []string{"c", "b", "a"})
}
