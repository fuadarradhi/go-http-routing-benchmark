// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}

var (
	gplusChi http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Chi", func() {
		gplusChi = loadChi(gplusAPI)
	})

	println()
}

func BenchmarkChi_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusChi, req)
}

func BenchmarkChi_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusChi, req)
}

func BenchmarkChi_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusChi, req)
}

func BenchmarkChi_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusChi, gplusAPI)
}
