// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package environupgrader_test

import (
	"github.com/juju/names/v5"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/api/base/testing"
	"github.com/juju/juju/api/controller/environupgrader"
	"github.com/juju/juju/rpc/params"
	coretesting "github.com/juju/juju/testing"
)

var (
	modelTag = names.NewModelTag("e5757df7-c86a-4835-84bc-7174af535d25")
)

var _ = gc.Suite(&EnvironUpgraderSuite{})

type EnvironUpgraderSuite struct {
	coretesting.BaseSuite
}

func (s *EnvironUpgraderSuite) TestModelEnvironVersion(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "EnvironUpgrader")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "ModelEnvironVersion")
		c.Check(arg, jc.DeepEquals, &params.Entities{
			Entities: []params.Entity{{Tag: modelTag.String()}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.IntResults{})
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{
				Result: 1,
			}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	version, err := client.ModelEnvironVersion(modelTag)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(version, gc.Equals, 1)
}

func (s *EnvironUpgraderSuite) TestModelEnvironVersionError(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{
				Error: &params.Error{Message: "foo"},
			}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	_, err := client.ModelEnvironVersion(modelTag)
	c.Assert(err, gc.ErrorMatches, "foo")
}

func (s *EnvironUpgraderSuite) TestModelEnvironArityMismatch(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{}, {}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	_, err := client.ModelEnvironVersion(modelTag)
	c.Assert(err, gc.ErrorMatches, "expected 1 result, got 2")
}

func (s *EnvironUpgraderSuite) TestModelTargetEnvironVersion(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "EnvironUpgrader")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "ModelTargetEnvironVersion")
		c.Check(arg, jc.DeepEquals, &params.Entities{
			Entities: []params.Entity{{Tag: modelTag.String()}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.IntResults{})
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{
				Result: 1,
			}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	version, err := client.ModelTargetEnvironVersion(modelTag)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(version, gc.Equals, 1)
}

func (s *EnvironUpgraderSuite) TestModelTargetEnvironVersionError(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{
				Error: &params.Error{Message: "foo"},
			}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	_, err := client.ModelTargetEnvironVersion(modelTag)
	c.Assert(err, gc.ErrorMatches, "foo")
}

func (s *EnvironUpgraderSuite) TestModelTargetEnvironArityMismatch(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.IntResults)) = params.IntResults{
			Results: []params.IntResult{{}, {}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	_, err := client.ModelTargetEnvironVersion(modelTag)
	c.Assert(err, gc.ErrorMatches, "expected 1 result, got 2")
}

func (s *EnvironUpgraderSuite) TestSetModelEnvironVersion(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "EnvironUpgrader")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "SetModelEnvironVersion")
		c.Check(arg, jc.DeepEquals, &params.SetModelEnvironVersions{
			Models: []params.SetModelEnvironVersion{{
				ModelTag: modelTag.String(),
				Version:  1,
			}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.ErrorResults{})
		*(result.(*params.ErrorResults)) = params.ErrorResults{
			Results: []params.ErrorResult{{Error: &params.Error{Message: "foo"}}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	err := client.SetModelEnvironVersion(modelTag, 1)
	c.Assert(err, gc.ErrorMatches, "foo")
}

func (s *EnvironUpgraderSuite) TestSetModelEnvironVersionArityMismatch(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.ErrorResults)) = params.ErrorResults{
			Results: []params.ErrorResult{{}, {}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	err := client.SetModelEnvironVersion(modelTag, 1)
	c.Assert(err, gc.ErrorMatches, "expected 1 result, got 2")
}

func (s *EnvironUpgraderSuite) TestSetModelStatus(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		c.Check(objType, gc.Equals, "EnvironUpgrader")
		c.Check(version, gc.Equals, 0)
		c.Check(id, gc.Equals, "")
		c.Check(request, gc.Equals, "SetModelStatus")
		c.Check(arg, jc.DeepEquals, params.SetStatus{
			Entities: []params.EntityStatusArgs{{
				Tag:    modelTag.String(),
				Status: "foo",
				Info:   "bar",
				Data: map[string]interface{}{
					"baz": "qux",
				},
			}},
		})
		c.Assert(result, gc.FitsTypeOf, &params.ErrorResults{})
		*(result.(*params.ErrorResults)) = params.ErrorResults{
			Results: []params.ErrorResult{{Error: &params.Error{Message: "foo"}}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	err := client.SetModelStatus(modelTag, "foo", "bar", map[string]interface{}{
		"baz": "qux",
	})
	c.Assert(err, gc.ErrorMatches, "foo")
}

func (s *EnvironUpgraderSuite) TestSetModelStatusArityMismatch(c *gc.C) {
	apiCaller := testing.APICallerFunc(func(objType string, version int, id, request string, arg, result interface{}) error {
		*(result.(*params.ErrorResults)) = params.ErrorResults{
			Results: []params.ErrorResult{{}, {}},
		}
		return nil
	})

	client := environupgrader.NewClient(apiCaller)
	err := client.SetModelStatus(modelTag, "foo", "bar", nil)
	c.Assert(err, gc.ErrorMatches, "expected 1 result, got 2")
}
