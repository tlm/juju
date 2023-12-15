// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gomock "go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/network"
)

type subnetSuite struct {
	testing.IsolationSuite

	st *MockState
}

var _ = gc.Suite(&subnetSuite{})

func (s *subnetSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.st = NewMockState(ctrl)

	return ctrl
}

func (s *subnetSuite) TestFailAddSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	subnetInfo := network.SubnetInfo{
		CIDR:              "192.168.0.0/20",
		ProviderId:        "provider-id-0",
		ProviderNetworkId: "provider-network-id-0",
		AvailabilityZones: []string{"az0"},
	}

	// Verify that the passed UUID is also returned.
	s.st.EXPECT().AddSubnet(gomock.Any(), gomock.Any(), subnetInfo).
		Return(errors.New("boom"))

	_, err := NewSubnetService(s.st).AddSubnet(context.Background(), subnetInfo)
	c.Assert(err, gc.ErrorMatches, "boom")
}

func (s *subnetSuite) TestAddSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	subnetInfo := network.SubnetInfo{
		CIDR:              "192.168.0.0/20",
		ProviderId:        "provider-id-0",
		ProviderNetworkId: "provider-network-id-0",
		AvailabilityZones: []string{"az0"},
	}

	var expectedUUID string
	// Verify that the passed UUID is also returned.
	s.st.EXPECT().AddSubnet(gomock.Any(), gomock.Any(), subnetInfo).
		Do(
			func(
				ctx context.Context,
				uuid string,
				subnet network.SubnetInfo,
			) error {
				expectedUUID = uuid
				return nil
			})

	returnedUUID, err := NewSubnetService(s.st).AddSubnet(context.Background(), subnetInfo)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(returnedUUID.String(), gc.Equals, expectedUUID)
}

func (s *subnetSuite) TestRetrieveSubnetByID(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().GetSubnet(gomock.Any(), "subnet0")
	_, err := NewSubnetService(s.st).Subnet(context.Background(), "subnet0")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *subnetSuite) TestFailRetrieveSubnetByID(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().GetSubnet(gomock.Any(), "unknown-subnet").
		Return(nil, errors.NotFoundf("subnet %q", "unknown-subnet"))
	_, err := NewSubnetService(s.st).Subnet(context.Background(), "unknown-subnet")
	c.Assert(err, gc.ErrorMatches, "subnet \"unknown-subnet\" not found")
}

func (s *subnetSuite) TestRetrieveSubnetByCIDRs(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().GetSubnetsByCIDR(gomock.Any(), "192.168.1.1", "10.0.0.1")
	_, err := NewSubnetService(s.st).SubnetsByCIDR(context.Background(), "192.168.1.1", "10.0.0.1")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *subnetSuite) TestFailRetrieveSubnetByCIDRs(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().GetSubnetsByCIDR(gomock.Any(), "192.168.1.1", "10.0.0.1").
		Return(nil, errors.New("querying subnets"))
	_, err := NewSubnetService(s.st).SubnetsByCIDR(context.Background(), "192.168.1.1", "10.0.0.1")
	c.Assert(err, gc.ErrorMatches, "querying subnets")
}

func (s *subnetSuite) TestUpdateSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().UpdateSubnet(gomock.Any(), "subnet0", "space0")
	err := NewSubnetService(s.st).UpdateSubnet(context.Background(), "subnet0", "space0")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *subnetSuite) TestFailUpdateSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().UpdateSubnet(gomock.Any(), "unknown-subnet", "space0").
		Return(errors.NotFoundf("subnet %q", "unknown-subnet"))
	err := NewSubnetService(s.st).UpdateSubnet(context.Background(), "unknown-subnet", "space0")
	c.Assert(err, gc.ErrorMatches, "subnet \"unknown-subnet\" not found")
}

func (s *subnetSuite) TestRemoveSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().DeleteSubnet(gomock.Any(), "subnet0")
	err := NewSubnetService(s.st).Remove(context.Background(), "subnet0")
	c.Assert(err, jc.ErrorIsNil)
}

func (s *subnetSuite) TestFailRemoveSubnet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.st.EXPECT().DeleteSubnet(gomock.Any(), "unknown-subnet").
		Return(errors.NotFoundf("subnet %q", "unknown-subnet"))
	err := NewSubnetService(s.st).Remove(context.Background(), "unknown-subnet")
	c.Assert(err, gc.ErrorMatches, "subnet \"unknown-subnet\" not found")
}
