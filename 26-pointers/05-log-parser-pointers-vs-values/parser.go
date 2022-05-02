package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
	// add more metrics if needed...
}

type parser struct {
	sum     map[string]result // total visits per domain
	domains []string          // unique domain names
	total   int               // total visits to all domains
	lines   int               // number of parsed lines
	lerr    error             // last error
}

func newParser() *parser {
	return &parser{
		sum: make(map[string]result),
	}
}

func parse(p *parser, line string) (r result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error

	r.domain = fields[0]
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}

	return
}

func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	// fmt.Printf("update.p   : %p - &p: %p\n\n", p, &p)

	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	p.total += r.visits

	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}

func err(p *parser) error {
	return p.lerr
}
