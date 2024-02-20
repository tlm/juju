// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package logtailer_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/juju/loggo/v2"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4/workertest"
	gc "gopkg.in/check.v1"

	corelogger "github.com/juju/juju/core/logger"
	"github.com/juju/juju/internal/logtailer"
	coretesting "github.com/juju/juju/testing"
)

type TailerSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&TailerSuite{})

func (s *TailerSuite) TestProcessForwardNoTail(c *gc.C) {
	testFileName := filepath.Join(c.MkDir(), "test.log")
	err := os.WriteFile(testFileName, []byte(logContent), 0644)
	c.Assert(err, jc.ErrorIsNil)

	tailer, err := logtailer.NewLogTailer(coretesting.ModelTag.Id(), testFileName, logtailer.LogTailerParams{
		NoTail: true,
	})
	c.Assert(err, jc.ErrorIsNil)

	var records []*corelogger.LogRecord
	logs := tailer.Logs()
	for {
		rec, ok := <-logs
		if !ok {
			break
		}
		records = append(records, rec)
	}
	c.Assert(records, jc.DeepEquals, logRecords)
}

func (s *TailerSuite) TestProcessReverseNoTail(c *gc.C) {
	testFileName := filepath.Join(c.MkDir(), "test.log")
	err := os.WriteFile(testFileName, []byte(logContent), 0644)
	c.Assert(err, jc.ErrorIsNil)

	tailer, err := logtailer.NewLogTailer(coretesting.ModelTag.Id(), testFileName, logtailer.LogTailerParams{
		NoTail:       true,
		InitialLines: 2,
	})
	c.Assert(err, jc.ErrorIsNil)

	var records []*corelogger.LogRecord
	logs := tailer.Logs()
	for {
		rec, ok := <-logs
		if !ok {
			break
		}
		records = append(records, rec)
	}
	c.Assert(records, jc.DeepEquals, logRecords[2:])
}

func (s *TailerSuite) fetchLogs(tailer logtailer.LogTailer, expected int) []*corelogger.LogRecord {
	var records []*corelogger.LogRecord
	timeout := time.After(testing.LongWait)
	for {
		select {
		case rec, ok := <-tailer.Logs():
			if !ok {
				return records
			}
			records = append(records, rec)
			if len(records) == expected {
				return records
			}
		case <-timeout:
			return records
		}
	}
}

func (s *TailerSuite) writeAdditionalLogs(c *gc.C, fileName string, lines []string) {
	go func() {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		c.Assert(err, jc.ErrorIsNil)
		defer func() {
			_ = f.Close()
		}()

		_, _ = fmt.Fprintln(f, "")
		for _, l := range lines {
			if l == "" {
				continue
			}
			_, _ = fmt.Fprintln(f, l)
		}
	}()
}

func (s *TailerSuite) TestProcessForwardTail(c *gc.C) {
	logLines := strings.Split(logContent, "\n")
	testFileName := filepath.Join(c.MkDir(), "test.log")
	f, err := os.OpenFile(testFileName, os.O_CREATE|os.O_WRONLY, 0644)
	c.Assert(err, jc.ErrorIsNil)
	defer func() {
		_ = f.Close()
	}()
	_, err = io.WriteString(f, strings.Join(logLines[:2], "\n"))
	c.Assert(err, jc.ErrorIsNil)

	tailer, err := logtailer.NewLogTailer(coretesting.ModelTag.Id(), testFileName, logtailer.LogTailerParams{})
	c.Assert(err, jc.ErrorIsNil)

	s.writeAdditionalLogs(c, testFileName, logLines[2:])

	records := s.fetchLogs(tailer, 4)
	c.Assert(records, jc.DeepEquals, logRecords)
}

func (s *TailerSuite) TestProcessReverseTail(c *gc.C) {
	logLines := strings.Split(logContent, "\n")
	testFileName := filepath.Join(c.MkDir(), "test.log")
	f, err := os.OpenFile(testFileName, os.O_CREATE|os.O_WRONLY, 0644)
	c.Assert(err, jc.ErrorIsNil)

	_, err = io.WriteString(f, strings.Join(logLines[:3], "\n"))
	c.Assert(f.Close(), jc.ErrorIsNil)
	c.Assert(err, jc.ErrorIsNil)

	tailer, err := logtailer.NewLogTailer(coretesting.ModelTag.Id(), testFileName, logtailer.LogTailerParams{
		InitialLines: 2,
	})
	c.Assert(err, jc.ErrorIsNil)
	records := s.fetchLogs(tailer, 2)

	s.writeAdditionalLogs(c, testFileName, logLines[3:])

	newRecords := s.fetchLogs(tailer, 1)
	result := append(records, newRecords...)
	c.Assert(result, jc.DeepEquals, logRecords[1:])
}

var logContent = `
machine-0 2024-02-15 06:23:22 DEBUG juju.worker.dependency engine.go:598 "db-accessor" manifold worker started at 2024-02-15 06:23:23.006402802 +0000 UTC
machine-0 2024-02-15 06:23:23 INFO juju.worker.dbaccessor worker.go:518 host is configured to use cloud-local address as a Dqlite node
machine-1 2024-02-15 06:23:24 WARNING juju.worker.dependency engine.go:598 "lease-manager" manifold worker started at 2024-02-15 06:23:23.016373586 +0000 UTC
machine-0 2024-02-15 06:23:25 CRITICAL juju.worker.dependency engine.go:598 "change-stream" manifold worker started at 2024-02-15 06:23:23.01677874 +0000 UTC`[1:]

var logRecords = []*corelogger.LogRecord{
	{
		Time:      mustParseTime("2024-02-15 06:23:22"),
		ModelUUID: coretesting.ModelTag.Id(),
		Entity:    "machine-0",
		Level:     loggo.DEBUG,
		Module:    "juju.worker.dependency",
		Location:  "engine.go:598",
		Message:   `"db-accessor" manifold worker started at 2024-02-15 06:23:23.006402802 +0000 UTC`,
	},
	{
		Time:      mustParseTime("2024-02-15 06:23:23"),
		ModelUUID: coretesting.ModelTag.Id(),
		Entity:    "machine-0",
		Level:     loggo.INFO,
		Module:    "juju.worker.dbaccessor",
		Location:  "worker.go:518",
		Message:   "host is configured to use cloud-local address as a Dqlite node",
	},
	{
		Time:      mustParseTime("2024-02-15 06:23:24"),
		ModelUUID: coretesting.ModelTag.Id(),
		Entity:    "machine-1",
		Level:     loggo.WARNING,
		Module:    "juju.worker.dependency",
		Location:  "engine.go:598",
		Message:   `"lease-manager" manifold worker started at 2024-02-15 06:23:23.016373586 +0000 UTC`,
	},
	{
		Time:      mustParseTime("2024-02-15 06:23:25"),
		ModelUUID: coretesting.ModelTag.Id(),
		Entity:    "machine-0",
		Level:     loggo.CRITICAL,
		Module:    "juju.worker.dependency",
		Location:  "engine.go:598",
		Message:   `"change-stream" manifold worker started at 2024-02-15 06:23:23.01677874 +0000 UTC`,
	},
}

func mustParseTime(in string) time.Time {
	out, err := time.Parse("2006-01-02 15:04:05", in)
	if err != nil {
		panic(err)
	}
	return out
}

type LogFilterSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&LogFilterSuite{})

func (s *LogFilterSuite) TestLevelFiltering(c *gc.C) {
	infoLevelRec := &corelogger.LogRecord{Level: loggo.INFO}
	errorLevelRec := &corelogger.LogRecord{Level: loggo.ERROR}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 1, &corelogger.LogRecord{Level: loggo.DEBUG})
		s.writeLogs(c, logFile, 1, infoLevelRec)
		s.writeLogs(c, logFile, 1, errorLevelRec)
		return logFile
	}
	params := logtailer.LogTailerParams{
		MinLevel: loggo.INFO,
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, infoLevelRec, errorLevelRec)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestIncludeEntity(c *gc.C) {
	machine0 := &corelogger.LogRecord{Entity: "machine-0"}
	foo0 := &corelogger.LogRecord{Entity: "unit-foo-0"}
	foo1 := &corelogger.LogRecord{Entity: "unit-foo-1"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 3, machine0)
		s.writeLogs(c, logFile, 2, foo0)
		s.writeLogs(c, logFile, 1, foo1)
		s.writeLogs(c, logFile, 3, machine0)
		return logFile
	}
	params := logtailer.LogTailerParams{
		IncludeEntity: []string{
			"unit-foo-0",
			"unit-foo-1",
		},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, foo0, foo0, foo1)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestIncludeEntityWildcard(c *gc.C) {
	machine0 := &corelogger.LogRecord{Entity: "machine-0"}
	foo0 := &corelogger.LogRecord{Entity: "unit-foo-0"}
	foo1 := &corelogger.LogRecord{Entity: "unit-foo-1"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 3, machine0)
		s.writeLogs(c, logFile, 2, foo0)
		s.writeLogs(c, logFile, 1, foo1)
		s.writeLogs(c, logFile, 3, machine0)
		return logFile
	}
	params := logtailer.LogTailerParams{
		IncludeEntity: []string{
			"unit-foo*",
		},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, foo0, foo0, foo1)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestExcludeEntity(c *gc.C) {
	machine0 := &corelogger.LogRecord{Entity: "machine-0"}
	foo0 := &corelogger.LogRecord{Entity: "unit-foo-0"}
	foo1 := &corelogger.LogRecord{Entity: "unit-foo-1"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 3, machine0)
		s.writeLogs(c, logFile, 2, foo0)
		s.writeLogs(c, logFile, 1, foo1)
		s.writeLogs(c, logFile, 3, machine0)
		return logFile
	}
	params := logtailer.LogTailerParams{
		ExcludeEntity: []string{
			"machine-0",
			"unit-foo-0",
		},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, foo1)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestExcludeEntityWildcard(c *gc.C) {
	machine0 := &corelogger.LogRecord{Entity: "machine-0"}
	foo0 := &corelogger.LogRecord{Entity: "unit-foo-0"}
	foo1 := &corelogger.LogRecord{Entity: "unit-foo-1"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 3, machine0)
		s.writeLogs(c, logFile, 2, foo0)
		s.writeLogs(c, logFile, 1, foo1)
		s.writeLogs(c, logFile, 3, machine0)
		return logFile
	}
	params := logtailer.LogTailerParams{
		ExcludeEntity: []string{
			"machine*",
			"unit-*-0",
		},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, foo1)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestIncludeModule(c *gc.C) {
	mod0 := &corelogger.LogRecord{Module: "foo.bar"}
	mod1 := &corelogger.LogRecord{Module: "juju.thing"}
	subMod1 := &corelogger.LogRecord{Module: "juju.thing.hai"}
	mod2 := &corelogger.LogRecord{Module: "elsewhere"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, mod1)
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, subMod1)
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, mod2)
		return logFile
	}
	params := logtailer.LogTailerParams{
		IncludeModule: []string{"juju.thing", "elsewhere"},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, mod1, subMod1, mod2)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestExcludeModule(c *gc.C) {
	mod0 := &corelogger.LogRecord{Module: "foo.bar"}
	mod1 := &corelogger.LogRecord{Module: "juju.thing"}
	subMod1 := &corelogger.LogRecord{Module: "juju.thing.hai"}
	mod2 := &corelogger.LogRecord{Module: "elsewhere"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, mod1)
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, subMod1)
		s.writeLogs(c, logFile, 1, mod0)
		s.writeLogs(c, logFile, 1, mod2)
		return logFile
	}
	params := logtailer.LogTailerParams{
		ExcludeModule: []string{"juju.thing", "elsewhere"},
	}
	assert := func(tailer logtailer.LogTailer) {
		s.assertTailer(c, tailer, mod0, mod0)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) TestIncludeExcludeModule(c *gc.C) {
	foo := &corelogger.LogRecord{Module: "foo"}
	bar := &corelogger.LogRecord{Module: "bar"}
	barSub := &corelogger.LogRecord{Module: "bar.thing"}
	baz := &corelogger.LogRecord{Module: "baz"}
	qux := &corelogger.LogRecord{Module: "qux"}
	logFile := filepath.Join(c.MkDir(), "logs.log")
	writeLogs := func() string {
		s.writeLogs(c, logFile, 1, foo)
		s.writeLogs(c, logFile, 1, bar)
		s.writeLogs(c, logFile, 1, barSub)
		s.writeLogs(c, logFile, 1, baz)
		s.writeLogs(c, logFile, 1, qux)
		return logFile
	}
	params := logtailer.LogTailerParams{
		IncludeModule: []string{"foo", "bar", "qux"},
		ExcludeModule: []string{"foo", "bar"},
	}
	assert := func(tailer logtailer.LogTailer) {
		// Except just "qux" because "foo" and "bar" were included and
		// then excluded.
		s.assertTailer(c, tailer, qux)
	}
	s.checkLogTailerFiltering(c, params, writeLogs, assert)
}

func (s *LogFilterSuite) checkLogTailerFiltering(
	c *gc.C,
	params logtailer.LogTailerParams,
	writeLogs func() string,
	assertTailer func(logtailer.LogTailer),
) {
	logFile := writeLogs()
	tailer, err := logtailer.NewLogTailer(coretesting.ModelTag.Id(), logFile, params)
	c.Assert(err, jc.ErrorIsNil)
	defer workertest.CleanKill(c, tailer)
	assertTailer(tailer)
}

func (s *LogFilterSuite) assertTailer(c *gc.C, tailer logtailer.LogTailer, template ...*corelogger.LogRecord) {
	c.Assert(template, gc.Not(gc.HasLen), 0)
	timeout := time.After(testing.LongWait)
	count := 0
	for {
		select {
		case log, ok := <-tailer.Logs():
			if !ok {
				c.Fatalf("tailer died unexpectedly: %v", tailer.Wait())
			}
			rec := s.normaliseLogTemplate(template[0])
			template = template[1:]

			c.Assert(log.Entity, gc.Equals, rec.Entity)
			c.Assert(log.Module, gc.Equals, rec.Module)
			c.Assert(log.Location, gc.Equals, rec.Location)
			c.Assert(log.Level, gc.Equals, rec.Level)
			c.Assert(log.Message, gc.Equals, rec.Message)
			c.Assert(log.Labels, gc.DeepEquals, rec.Labels)
			count++
			if len(template) == 0 {
				return
			}
		case <-timeout:
			c.Fatalf("timed out waiting for logs (received %d)", count)
		}
	}
}

func (s *LogFilterSuite) normaliseLogTemplate(template *corelogger.LogRecord) corelogger.LogRecord {
	rec := *template
	if rec.Entity == "" {
		rec.Entity = "not-a-tag"
	}
	if rec.Module == "" {
		rec.Module = "module"
	}
	if rec.Location == "" {
		rec.Location = "loc"
	}
	if rec.Level == loggo.UNSPECIFIED {
		rec.Level = loggo.INFO
	}
	if rec.Message == "" {
		rec.Message = "message"
	}
	var labelsOut []string
	for k, v := range rec.Labels {
		labelsOut = append(labelsOut, fmt.Sprintf("%s:%s", k, v))
	}
	return rec
}

// writeLogs creates count log messages at the current time using
// the supplied template.
func (s *LogFilterSuite) writeLogs(c *gc.C, logFie string, count int, template *corelogger.LogRecord) {
	t := coretesting.ZeroTime()
	s.writeLogsT(c, logFie, t, t, count, template)
}

// writeLogsT creates count log messages between startTime and
// endTime using the supplied template
func (s *LogFilterSuite) writeLogsT(c *gc.C, logFile string, startTime, endTime time.Time, count int, template *corelogger.LogRecord) {
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	c.Assert(err, jc.ErrorIsNil)
	defer func() {
		_ = f.Close()
	}()

	interval := endTime.Sub(startTime) / time.Duration(count)
	t := startTime
	for i := 0; i < count; i++ {
		rec := s.normaliseLogTemplate(template)
		line := strings.Join([]string{
			rec.Entity,
			rec.Time.In(time.UTC).Format("2006-01-02 15:04:05"),
			rec.Level.String(),
			rec.Module,
			rec.Location,
			rec.Message,
			// TODO(debug-log) - implement label filtering when json formatting is added
			//strings.Join(labelsOut, ","),
		}, " ")
		_, err = f.WriteString(line + "\n")
		c.Assert(err, jc.ErrorIsNil)
		t = t.Add(interval)
	}
}
