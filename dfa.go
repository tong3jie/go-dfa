package dfa

import (
	"strings"
	"sync"
)

type DFA struct {
	mu           sync.RWMutex
	trie         *Trie
	replaceStr   string
	invalidWords map[string]struct{}
	star         int
	question     int
	builder      sync.Pool
}

func New(opts ...Option) *DFA {
	opt := loadOptions(opts...)

	f := &DFA{
		trie:         NewTrie(),
		invalidWords: make(map[string]struct{}),
		builder:      sync.Pool{New: func() any { return new(strings.Builder) }},
	}

	if opt.star > 0 {
		f.star = opt.star
	} else {
		f.star = 5
	}

	if opt.question > 0 {
		f.question = opt.question
	} else {
		f.question = 1
	}

	if opt.defaultStr != "" {
		f.replaceStr = opt.defaultStr
	} else {
		f.replaceStr = defaultReplaceStr
	}
	if opt.invalidWords != "" {
		for _, s := range opt.invalidWords {
			f.invalidWords[string(s)] = struct{}{}
		}
	} else {
		for _, s := range defaultInvalidWorlds {
			f.invalidWords[string(s)] = struct{}{}
		}
	}

	return f
}

func (f *DFA) AddWords(words []string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(words) > 0 {
		for _, s := range words {
			f.trie.Insert(strings.TrimSpace(s))
		}
	}
}

func (f *DFA) RemoveWords(words []string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(words) > 0 {
		for _, s := range words {
			f.trie.Remove(strings.TrimSpace(s))
		}
	}
}

func (f *DFA) Check(txt string) ([]string, bool) {
	_, found, b := f.check(txt, false)
	return found, b
}

func (f *DFA) CheckAndReplace(txt string) (string, []string, bool) {
	return f.check(txt, true)
}

func (f *DFA) check(txt string, replace bool) (dist string, found []string, b bool) {

	str := []rune(txt)
	result := txt
	nodeMap := f.trie.Root().Child
	start := -1
	builder := f.builder.Get().(*strings.Builder)
	defer f.builder.Put(builder)

	f.mu.RLock()
	defer f.mu.RUnlock()

	i := 0
	for i < len(str) {
		if _, ok := f.invalidWords[string(str[i])]; ok {
			i++
			continue
		}

		if node, ok := nodeMap[str[i]]; ok {
			if start == -1 {
				start = i
			}
			if node.IsEnd {
				found = append(found, string(str[start:i+1]))

				if replace {
					builder.WriteString(string(str[:start]))
					builder.WriteString(f.replaceStr)
					result = builder.String() + string(str[i+1:])
					builder.Reset()
				}
				start = -1
				nodeMap = f.trie.Root().Child
			} else {
				nodeMap = node.Child
				if _, ok := nodeMap['?']; ok {
					i += f.question
					found = append(found, string(str[start:i+1]))
					builder.WriteString(string(str[:start]))
					builder.WriteString(f.replaceStr)
					result = builder.String() + string(str[i+1:])
					builder.Reset()
				} else if _, ok := nodeMap['*']; ok {
					i += f.star
					found = append(found, string(str[start:i+1]))
					builder.WriteString(string(str[:start]))
					builder.WriteString(f.replaceStr)
					result = builder.String() + string(str[i+1:])
					builder.Reset()
				}
			}

		} else {
			start = -1
			nodeMap = f.trie.Root().Child
		}
		i++
	}

	b = len(found) > 0
	if replace {
		dist = result
	}
	return
}
