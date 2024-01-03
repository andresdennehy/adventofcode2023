package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type PulseType int;

const (
    Low PulseType = iota
    High
)

type QueueItem struct {
    module string
    pulse PulseType
    sourceModule Module
}

type FlipFlop struct {
    name string
    on bool
}

type Conjunction struct {
    name string
    remembers map[Module]PulseType  // Memory for each input module
}

type Broadcaster struct {
    name string
}

type Output struct {
    name string
}

type Module interface {
    getName() string
    ProcessPulse(PulseType, Module) (PulseType, bool)
}

func (f *FlipFlop) ProcessPulse(pulse PulseType, sourceModule Module) (PulseType, bool) {
    if pulse == High {
        return High, false;
    } else {
        switch f.on {
            case true:
                // If it was on, it turns off and sends a low pulse
                f.on = false;
                return Low, true;
            case false:
            // If it was off, it turns on and sends a high pulse
                f.on = true;
                return High, true;
        }
    }

    return Low, false;
}

func (c *Conjunction) ProcessPulse(pulse PulseType, sourceModule Module) (PulseType, bool) {
    c.remembers[sourceModule] = pulse;

    // If it remembers high pulses for all inputs, it sends a low pulse.
    // Otherwise, it sends a high pulse.
    for _, v := range c.remembers {
        if v == Low {
            return High, true;
        }
    }

    //
    return Low, true;
}

func (b *Broadcaster) ProcessPulse(pulse PulseType, sourceModule Module) (PulseType, bool) {
    return Low, true;
}

func (o *Output) ProcessPulse(pulse PulseType, sourceModule Module) (PulseType, bool) {
    return Low, false;
}

func (f *FlipFlop) getName()string {
    return f.name;
}

func (c *Conjunction) getName()string {
    return c.name;
}

func (b *Broadcaster) getName()string {
    return b.name;
}

func (o *Output) getName()string {
    return o.name;
}

func newModule(config string) (Module, string) {
    switch config[0] {
        case '%':
            return &FlipFlop{config[1:], false}, config[1:];
        case '&':
            return &Conjunction{config[1:], make(map[Module]PulseType)}, config[1:];
    }

    switch config {
        case "broadcaster":
            return &Broadcaster{"broadcaster"}, config;
        case "output":
            return &Output{"output"}, config;
    }

    return nil, "";
}

func printTree(tree map[Module][]Module) {
    fmt.Printf("Printing adjacency list\n")
    for k, v := range tree {
        fmt.Printf("Key: %v, Value: %v\n", k, v);
    }
}

func pushButton(tree map[Module][]Module, modules map[string]Module) (int, int) {
    var lowPulses, highPulses int;

    queue := []QueueItem{QueueItem{"broadcaster", Low, nil}};
    for len(queue) > 0 {
        currentModule := modules[queue[0].module];
        currentPulseType := queue[0].pulse;
        switch currentPulseType {
            case Low:
                lowPulses++;
            case High:
                highPulses++;
        }
        sourceModule := queue[0].sourceModule;
        queue = queue[1:];
        newPulse, continues := currentModule.ProcessPulse(currentPulseType, sourceModule);
        if continues {
            for _, v := range tree[currentModule] {
                if v == nil {
                    switch newPulse {
                        case Low:
                            lowPulses++;
                        case High:
                            highPulses++;
                    }
                    continue;
                }
                queue = append(queue, QueueItem{v.getName(), newPulse, currentModule});
            }
        }
    }

    return lowPulses, highPulses;
}

func main() {
    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    rawTree := make(map[string][]string);
    modules := make(map[string]Module);

    // Assuming that the input is well-formed, no loops
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        var source, targets string;
        lineParts := strings.Split(scanner.Text(), " -> ");
        source = lineParts[0];
        targets = lineParts[1];
        targetList := strings.Split(targets, ", ");

        newModule, name := newModule(source);
        modules[name] = newModule;
        rawTree[name] = targetList;
    }

    fmt.Printf("Status before processing\n");
    fmt.Printf("%v", rawTree);

    tree := make(map[Module][]Module);
    for k, v := range rawTree {
        for _, target := range v {
            if _, ok := modules[target]; !ok {
                newModule, name := newModule(target);
                modules[name] = newModule;
            }
            tree[modules[k]] = append(tree[modules[k]], modules[target]);
        }
    }

    // Initialize conjunction memories
    for k, v := range tree {
        for _, v := range v {
            switch v.(type) {
                case *Conjunction:
                    v.(*Conjunction).remembers[k] = Low;
            }
        }
    }

    var lowPulses, highPulses int;
    for i:=0; i<1000; i++ {
        low, high := pushButton(tree, modules);
        lowPulses += low;
        highPulses += high;
    }

    fmt.Printf("Low pulses: %d, High pulses: %d\n", lowPulses, highPulses);
    fmt.Printf("Result is %d\n", lowPulses * highPulses);
}