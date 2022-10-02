// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charmhub

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/juju/charm/v8"
	"github.com/juju/collections/set"
	"github.com/juju/errors"

	"github.com/juju/juju/charmhub/transport"
	"github.com/juju/juju/core/arch"
	corecharm "github.com/juju/juju/core/charm"
	coreseries "github.com/juju/juju/core/series"
)

func convertCharmInfoResult(info transport.InfoResponse, arch, series string) (InfoResponse, error) {
	ir := InfoResponse{
		Type:        string(info.Type),
		ID:          info.ID,
		Name:        info.Name,
		Description: info.Entity.Description,
		Publisher:   publisher(info.Entity),
		Summary:     info.Entity.Summary,
		Tags:        categories(info.Entity.Categories),
		StoreURL:    info.Entity.StoreURL,
	}
	var isKubernetes bool
	switch transport.Type(ir.Type) {
	case transport.BundleType:
		ir.Bundle = convertBundle(info.Entity.Charms)
		// TODO (stickupkid): Get the Bundle.Release and set it to the
		// InfoResponse at a high level.
	case transport.CharmType:
		ir.Charm, isKubernetes = convertCharm(info)
	}

	seriesSet := set.NewStrings()
	for _, base := range info.DefaultRelease.Revision.Bases {
		s, _ := coreseries.VersionSeries(base.Channel)
		if s != "" {
			seriesSet.Add(s)
		}
	}
	ir.Series = seriesSet.SortedValues()

	var err error
	ir.Tracks, ir.Channels, err = filterChannels(info.ChannelMap, isKubernetes, arch, series)
	if err != nil {
		return ir, errors.Trace(err)
	}
	return ir, nil
}

func convertCharmFindResults(responses []transport.FindResponse) []FindResponse {
	results := make([]FindResponse, len(responses))
	for i, resp := range responses {
		results[i] = convertCharmFindResult(resp)
	}
	return results
}

func convertCharmFindResult(resp transport.FindResponse) FindResponse {
	result := FindResponse{
		Type:      string(resp.Type),
		ID:        resp.ID,
		Name:      resp.Name,
		Publisher: publisher(resp.Entity),
		Summary:   resp.Entity.Summary,
		Version:   resp.DefaultRelease.Revision.Version,
		StoreURL:  resp.Entity.StoreURL,
	}
	supported := transformFindArchitectureSeries(resp.DefaultRelease)
	result.Arches, result.OS, result.Series = supported.Architectures, supported.OS, supported.Series
	return result
}

func convertBundle(charms []transport.Charm) *Bundle {
	bundle := &Bundle{
		Charms: make([]BundleCharm, len(charms)),
	}
	for i, v := range charms {
		bundle.Charms[i] = BundleCharm{
			Name: v.Name,
		}
	}
	return bundle
}

func convertCharm(info transport.InfoResponse) (ch *Charm, isKubernetes bool) {
	ch = &Charm{
		UsedBy: info.Entity.UsedBy,
	}
	if meta := unmarshalCharmMetadata(info.DefaultRelease.Revision.MetadataYAML); meta != nil {
		ch.Subordinate = meta.Subordinate
		ch.Relations = transformRelations(meta.Requires, meta.Provides)
		cm := charmMeta{meta: meta}
		isKubernetes = corecharm.IsKubernetes(cm)
	}
	if cfg := unmarshalCharmConfig(info.DefaultRelease.Revision.ConfigYAML); cfg != nil {
		ch.Config = &charm.Config{
			Options: toCharmOptionMap(cfg),
		}
	}
	return ch, isKubernetes
}

func includeChannel(p []corecharm.Platform, architecture, series string) bool {
	allArch := architecture == ArchAll
	allSeries := series == SeriesAll

	// If we're searching for everything then we can skip the filtering logic
	// and return immediately.
	if allArch && allSeries {
		return true
	}

	archSet := channelArches(p)
	seriesSet := channelSeries(p)

	if (allArch || archSet.Contains(architecture)) &&
		(allSeries || seriesSet.Contains(series) || seriesSet.Contains(SeriesAll)) {
		return true
	}
	return false
}

func channelSeries(platforms []corecharm.Platform) set.Strings {
	series := set.NewStrings()
	for _, v := range platforms {
		s, err := coreseries.VersionSeries(v.Channel)
		if err == nil {
			series.Add(s)
		}
	}
	return series
}

func channelBases(platforms []corecharm.Platform) set.Strings {
	bases := set.NewStrings()
	for _, v := range platforms {
		bases.Add(fmt.Sprintf("%s/%s", v.OS, v.Channel))
	}
	return bases
}

func channelArches(platforms []corecharm.Platform) set.Strings {
	arches := set.NewStrings()
	for _, v := range platforms {
		arches.Add(v.Architecture)
	}
	// If the platform contains all the arches, just return them exploded.
	// This makes the filtering logic simpler for plucking an architecture out
	// of the channels, we should aim to do the same for series.
	if arches.Contains(ArchAll) {
		return set.NewStrings(arch.AllArches().StringList()...)
	}
	return arches
}

func publisher(ch transport.Entity) string {
	publisher, _ := ch.Publisher["display-name"]
	return publisher
}

func categories(cats []transport.Category) []string {
	if len(cats) == 0 {
		return nil
	}
	result := make([]string, len(cats))
	for i, v := range cats {
		result[i] = v.Name
	}
	return result
}

// supported defines a tuple of extracted items from a platform.
type supported struct {
	Architectures []string
	OS            []string
	Series        []string
}

// transformFindArchitectureSeries returns a supported type which contains
// architectures and series for a given channel map.
func transformFindArchitectureSeries(channel transport.FindChannelMap) supported {
	if len(channel.Revision.Bases) < 1 {
		return supported{}
	}

	var (
		arches = set.NewStrings()
		os     = set.NewStrings()
		series = set.NewStrings()
	)
	for _, p := range channel.Revision.Bases {
		arches.Add(p.Architecture)
		os.Add(p.Name)
		// TODO hml - for this to be correct, must determine IsKubernetes from metadata.
		s, _ := coreseries.VersionSeries(p.Channel)
		series.Add(s)
	}
	return supported{
		Architectures: arches.SortedValues(),
		OS:            os.SortedValues(),
		Series:        series.SortedValues(),
	}
}

func toCharmOptionMap(config *charm.Config) map[string]charm.Option {
	if config == nil {
		return nil
	}
	result := make(map[string]charm.Option)
	for key, value := range config.Options {
		result[key] = toParamsCharmOption(value)
	}
	return result
}

func toParamsCharmOption(opt charm.Option) charm.Option {
	return charm.Option{
		Type:        opt.Type,
		Description: opt.Description,
		Default:     opt.Default,
	}
}

func unmarshalCharmMetadata(metadataYAML string) *charm.Meta {
	if metadataYAML == "" {
		return nil
	}
	m := metadataYAML
	meta, err := charm.ReadMeta(bytes.NewBufferString(m))
	if err != nil {
		// Do not fail on unmarshalling metadata, log instead.
		// This should not happen, however at implementation
		// we were dealing with handwritten data for test, not
		// the real deal.  Usually charms are validated before
		// being uploaded to the store.
		logger.Warningf(errors.Annotate(err, "cannot unmarshal charm metadata").Error())
		return nil
	}
	return meta
}

func unmarshalCharmConfig(configYAML string) *charm.Config {
	if configYAML == "" {
		return nil
	}
	cfgYaml := configYAML
	cfg, err := charm.ReadConfig(bytes.NewBufferString(cfgYaml))
	if err != nil {
		// Do not fail on unmarshalling metadata, log instead.
		// This should not happen, however at implementation
		// we were dealing with handwritten data for test, not
		// the real deal.  Usually charms are validated before
		// being uploaded to the store.
		logger.Warningf(errors.Annotate(err, "cannot unmarshal charm config").Error())
		return nil
	}
	return cfg
}

func transformRelations(requires, provides map[string]charm.Relation) map[string]map[string]string {
	if len(requires) == 0 && len(provides) == 0 {
		logger.Debugf("no relation data found in charm meta data")
		return nil
	}
	relations := make(map[string]map[string]string)
	if provides, ok := formatRelationPart(provides); ok {
		relations["provides"] = provides
	}
	if requires, ok := formatRelationPart(requires); ok {
		relations["requires"] = requires
	}
	return relations
}

func formatRelationPart(r map[string]charm.Relation) (map[string]string, bool) {
	if len(r) <= 0 {
		return nil, false
	}
	relations := make(map[string]string, len(r))
	for k, v := range r {
		relations[k] = v.Interface
	}
	return relations, true
}

type charmMeta struct {
	meta     *charm.Meta
	manifest *charm.Manifest
}

func (c charmMeta) Meta() *charm.Meta {
	return c.meta
}

func (c charmMeta) Manifest() *charm.Manifest {
	return c.manifest
}

// filterChannels returns channel map data in a format that facilitates
// determining track order and open vs closed channels for displaying channel
// data. The result is filtered on series and arch.
func filterChannels(channelMap []transport.InfoChannelMap, isKub bool, arch, series string) ([]string, map[string]Channel, error) {
	var trackList []string

	seen := set.NewStrings("")
	channels := make(map[string]Channel, len(channelMap))

	for _, cm := range channelMap {
		ch := cm.Channel
		// Per the charmhub/snap channel spec.
		if ch.Track == "" {
			ch.Track = "latest"
		}
		if !seen.Contains(ch.Track) {
			seen.Add(ch.Track)
			trackList = append(trackList, ch.Track)
		}

		platforms := convertBasesToPlatforms(cm.Revision.Bases)
		if !includeChannel(platforms, arch, series) {
			continue
		}

		currentCh := Channel{
			Revision:   cm.Revision.Revision,
			ReleasedAt: ch.ReleasedAt,
			Risk:       ch.Risk,
			Track:      ch.Track,
			Size:       cm.Revision.Download.Size,
			Version:    cm.Revision.Version,
			Arches:     channelArches(platforms).SortedValues(),
			Bases:      channelBases(platforms).SortedValues(),
		}
		// TODO(juju3) - remove series
		if isKub {
			currentCh.Series = []string{"kubernetes"}
		} else {
			currentCh.Series = channelSeries(platforms).SortedValues()
		}

		chName := ch.Track + "/" + ch.Risk
		if existingCh, ok := channels[chName]; ok {

			currentChReleasedAt, err := time.Parse(time.RFC3339, currentCh.ReleasedAt)
			if err != nil {
				return nil, nil, errors.Annotatef(err, "invalid time format, expected RFC3339, got %s", currentCh.ReleasedAt)
			}
			existingChReleasedAt, err := time.Parse(time.RFC3339, existingCh.ReleasedAt)
			if err != nil {
				return nil, nil, errors.Annotatef(err, "invalid time format, expected RFC3339, got %s", existingCh.ReleasedAt)
			}
			if currentChReleasedAt.After(existingChReleasedAt) {
				channels[chName] = currentCh
			}
		} else {
			channels[chName] = currentCh
		}

	}
	return trackList, channels, nil
}

func convertBasesToPlatforms(in []transport.Base) []corecharm.Platform {
	out := make([]corecharm.Platform, len(in))
	for i, v := range in {
		out[i] = corecharm.Platform{
			Architecture: v.Architecture,
			OS:           strings.ToLower(v.Name),
			Channel:      v.Channel,
		}
	}
	return out
}
