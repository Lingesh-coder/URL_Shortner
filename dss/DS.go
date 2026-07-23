package dss

import "regexp"

var Ma=make(map[string]string)
var Am=make(map[string]string)
var AnalyticsMap=make(map[string]int)
var MaxURLLen int = 50
var Size int=6
var DomainCheck = regexp.MustCompile(`^https:\/\/[A-Za-z0-9.-]+\.[A-Za-z]{2,64}`)
var ShortURLRegex = regexp.MustCompile(`^test.com\/`)