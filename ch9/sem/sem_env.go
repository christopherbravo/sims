// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"

	"github.com/emer/emergent/env"
	"github.com/emer/emergent/erand"
	"github.com/emer/etable/etensor"
)

// SemEnv presents paragraphs of text, loaded from file(s)
// This assumes files have all been pre-filtered so only relevant words are present.
type SemEnv struct {
	Nm           string          `desc:"name of this environment"`
	Dsc          string          `desc:"description of this environment"`
	Sequential   bool            `desc:"if true, go sequentially through paragraphs -- else permuted"`
	Order        []int           `desc:"permuted order of paras to present if not sequential -- updated every time through the list"`
	TextFiles    []string        `desc:"paths to text files"`
	Words        []string        `desc:"list of words, in alpha order"`
	WordMap      map[string]int  `desc:"map of words onto index in Words list"`
	CurParaState etensor.Float32 `desc:"current para activation state"`
	Paras        [][]string      `desc:"paragraphs"`
	Run          env.Ctr         `view:"inline" desc:"current run of model as provided during Init"`
	Epoch        env.Ctr         `view:"inline" desc:"number of times through Seq.Max number of sequences"`
	Trial        env.Ctr         `view:"inline" desc:"trial is the step counter within epoch -- this is the index into Paras"`
}

func (ev *SemEnv) Name() string { return ev.Nm }
func (ev *SemEnv) Desc() string { return ev.Dsc }

func (ev *SemEnv) Validate() error {
	return nil
}

func (ev *SemEnv) Counters() []env.TimeScales {
	return []env.TimeScales{env.Run, env.Epoch, env.Sequence, env.Trial}
}

func (ev *SemEnv) States() env.Elements {
	sz := ev.CurParaState.Shapes()
	nms := ev.CurParaState.DimNames()
	els := env.Elements{
		{"Input", sz, nms},
	}
	return els
}

func (ev *SemEnv) State(element string) etensor.Tensor {
	switch element {
	case "Input":
		return &ev.CurParaState
	}
	return nil
}

func (ev *SemEnv) Actions() env.Elements {
	return nil
}

func (ev *SemEnv) Defaults() {
}

func (ev *SemEnv) Init(run int) {
	ev.Run.Scale = env.Run
	ev.Epoch.Scale = env.Epoch
	ev.Trial.Scale = env.Trial
	ev.Run.Init()
	ev.Epoch.Init()
	ev.Trial.Init()
	ev.Run.Cur = run
	np := len(ev.Paras)
	ev.Order = rand.Perm(np) // always start with new one so random order is identical
	// and always maintain Order so random number usage is same regardless, and if
	// user switches between Sequential and random at any point, it all works..
	ev.Trial.Max = np
	ev.Trial.Cur = -1 // init state -- key so that first Step() = 0

	nw := len(ev.Words)
	ev.CurParaState.SetShape([]int{1, nw}, nil, []string{"1", "Words"})
}

// OpenTexts opens multiple text files -- use this as main API
// even if only opening one text
func (ev *SemEnv) OpenTexts(txts []string) {
	ev.TextFiles = txts
	ev.Paras = make([][]string, 0, 2000)
	for _, tf := range ev.TextFiles {
		ev.OpenText(tf)
	}
}

// OpenText opens one text file
func (ev *SemEnv) OpenText(fname string) {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return
	}
	scan := bufio.NewScanner(fp) // line at a time
	cur := []string{}
	for scan.Scan() {
		b := scan.Bytes()
		bs := string(b)
		sp := strings.Fields(bs)
		if len(sp) == 0 {
			ev.Paras = append(ev.Paras, cur)
			cur = []string{}
		} else {
			cur = append(cur, sp...)
		}
	}
	if len(cur) > 0 {
		ev.Paras = append(ev.Paras, cur)
	}
}

func (ev *SemEnv) OpenWords(fname string) {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return
	}
	ev.Words = make([]string, 0, 3000)
	scan := bufio.NewScanner(fp) // line at a time
	for scan.Scan() {
		b := scan.Bytes()
		bs := string(b)
		sp := strings.Fields(bs)
		if len(sp) > 0 {
			ev.Words = append(ev.Words, sp...)
		}
	}
	ev.WordMapFmWords()
}

func (ev *SemEnv) WordMapFmWords() {
	ev.WordMap = make(map[string]int, len(ev.Words))
	for i, wrd := range ev.Words {
		ev.WordMap[wrd] = i
	}
}

func (ev *SemEnv) WordsFmWordMap() {
	ev.Words = make([]string, len(ev.WordMap))
	ctr := 0
	for wrd, _ := range ev.WordMap {
		ev.Words[ctr] = wrd
		ctr++
	}
	sort.Strings(ev.Words)
	for i, wrd := range ev.Words {
		ev.WordMap[wrd] = i
	}
}

func (ev *SemEnv) WordsFmText() {
	ev.WordMap = make(map[string]int, len(ev.Words))
	for _, para := range ev.Paras {
		for _, wrd := range para {
			ev.WordMap[wrd] = -1
		}
	}
	ev.WordsFmWordMap()
}

func (ev *SemEnv) Step() bool {
	ev.Epoch.Same()      // good idea to just reset all non-inner-most counters at start
	if ev.Trial.Incr() { // if true, hit max, reset to 0
		erand.PermuteInts(ev.Order)
		ev.Epoch.Incr()
	}
	ev.SetParaState()
	return true
}

func (ev *SemEnv) Action(element string, input etensor.Tensor) {
	// nop
}

func (ev *SemEnv) Counter(scale env.TimeScales) (cur, prv int, chg bool) {
	switch scale {
	case env.Run:
		return ev.Run.Query()
	case env.Epoch:
		return ev.Epoch.Query()
	case env.Trial:
		return ev.Trial.Query()
	}
	return -1, -1, false
}

// Compile-time check that implements Env interface
var _ env.Env = (*SemEnv)(nil)

// String returns the string rep of the LED env state
func (ev *SemEnv) String() string {
	cpar := ev.CurPara()
	if cpar == nil {
		return ""
	}
	str := cpar[0]
	if len(cpar) > 1 {
		str += " " + cpar[1]
		if len(cpar) > 2 {
			str += " ... " + cpar[len(cpar)-1]
		}
	}
	return str
}

// ParaIdx returns the current idx number in Paras, based on Sequential / perumuted Order
func (ev *SemEnv) ParaIdx() int {
	if ev.Trial.Cur < 0 {
		return -1
	}
	if ev.Sequential {
		return ev.Trial.Cur
	}
	return ev.Order[ev.Trial.Cur]
}

// CurPara returns the current paragraph
func (ev *SemEnv) CurPara() []string {
	pidx := ev.ParaIdx()
	if pidx >= 0 && pidx < len(ev.Paras) {
		return ev.Paras[pidx]
	}
	return nil
}

// SetParaState sets the para state from current para
func (ev *SemEnv) SetParaState() {
	cpar := ev.CurPara()
	ev.CurParaState.SetZeros()
	for _, wrd := range cpar {
		widx := ev.WordMap[wrd]
		ev.CurParaState.SetFloat1D(widx, 1)
	}
}