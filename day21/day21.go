package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type operation struct {
	name      string
	value1    string
	value2    string
	operation string
}

func operate(o *operation, valueMap map[string]int, operationMap map[string]*operation) (int, error) {
	if _, ok := valueMap[o.value1]; !ok {
		if op, ok := operationMap[o.value1]; ok {
			valueMap[o.value1], _ = operate(op, valueMap, operationMap)
		} else {
			return 0, errors.New("not found in operationMap")
		}
	}
	if _, ok := valueMap[o.value2]; !ok {
		if op, ok := operationMap[o.value2]; ok {
			valueMap[o.value2], _ = operate(op, valueMap, operationMap)
		} else {
			return 0, errors.New("not found in operationMap")
		}
	}
	v1, v2 := valueMap[o.value1], valueMap[o.value2]

	switch o.operation {
	case "+":
		return v1 + v2, nil
	case "-":
		return v1 - v2, nil
	case "*":
		return v1 * v2, nil
	case "/":
		return v1 / v2, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

func unoperate(o *operation, valueMap map[string]int, operationMap map[string]*operation) (int, []string) {
	influence := []string{}
	if _, ok := valueMap[o.value1]; !ok {
		if op, ok := operationMap[o.value1]; ok {
			pi := []string{}
			valueMap[o.value1], pi = unoperate(op, valueMap, operationMap)
			if len(pi) > 0 {
				influence = pi
			}
		} else {
			return 0, nil
		}
	}
	if _, ok := valueMap[o.value2]; !ok {
		if op, ok := operationMap[o.value2]; ok {
			pi := []string{}
			valueMap[o.value2], pi = unoperate(op, valueMap, operationMap)
			if len(pi) > 0 {
				influence = pi
			}
		} else {
			return 0, nil
		}
	}
	if len(influence) > 0 || o.value1 == "humn" || o.value2 == "humn" {
		if o.value1 == "humn" || o.value2 == "humn" {
			influence = append(influence, "humn")
		}
		return 0, append(influence, o.name)
	} else if o.name == "root" {
		influence = append(influence, o.name)
	}
	v1, v2 := valueMap[o.value1], valueMap[o.value2]

	switch o.operation {
	case "+":
		return v1 + v2, influence
	case "-":
		return v1 - v2, influence
	case "*":
		return v1 * v2, influence
	case "/":
		return v1 / v2, influence
	default:
		return 0, nil
	}
}

func dispelMystery(valueMap map[string]int, operationMap map[string]*operation, influence []string) int {
	currentNode := influence[len(influence)-1]
	var eq int
	for len(influence) > 0 {
		if currentNode == "humn" {
			break
		}
		var nodeR string
		n := operationMap[influence[len(influence)-1]]
		rhs := false
		if n.value1 == influence[len(influence)-2] {
			nodeR = n.value2
		} else {
			nodeR = n.value1
			rhs = true
		}
		if currentNode == "root" {
			eq = valueMap[nodeR]
		} else {
			switch operationMap[currentNode].operation {
			case "+":
				eq -= valueMap[nodeR]
			case "-":
				if rhs {
					eq = valueMap[nodeR] - eq
				} else {
					eq += valueMap[nodeR]
				}
			case "*":
				eq /= valueMap[nodeR]
			case "/":
				eq *= valueMap[nodeR]
			}
		}
		influence = influence[:len(influence)-1]
		currentNode = influence[len(influence)-1]
	}
	return eq
}

type value struct {
	name  string
	value int
}

func ParseString(input string) (*operation, *value) {
	parts := strings.Split(input, ": ")
	name := parts[0]
	data := parts[1]
	if strings.ContainsAny(data, "*/+-") {
		opData := strings.Split(data, " ")
		return &operation{name, opData[0], opData[2], opData[1]}, nil
	} else {
		num, err := strconv.Atoi(data)
		if err != nil {
			log.Fatal("error parsing value", data)
		}
		return nil, &value{name, num}
	}

}

func QueueChecks(input []string) int {
	valueMap := make(map[string]int)
	operationMap := make(map[string]*operation)

	jobs := make(chan string, len(input))
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				o, v := ParseString(j)
				if v != nil {
					valueMap[v.name] = v.value
				} else if o != nil {
					operationMap[o.name] = o
				}
			} else {
				done <- true
				return
			}
		}
	}()

	for _, v := range input {
		if v != "" {
			jobs <- v
		}
	}
	close(jobs)

	<-done
	res, err := operate(operationMap["root"], valueMap, operationMap)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func MonkeyBrain(input []string) int {
	valueMap := make(map[string]int)
	operationMap := make(map[string]*operation)

	jobs := make(chan string, len(input))
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				o, v := ParseString(j)
				if v != nil {
					valueMap[v.name] = v.value
				} else if o != nil {
					operationMap[o.name] = o
				}
			} else {
				done <- true
				return
			}
		}
	}()

	for _, v := range input {
		if v != "" {
			jobs <- v
		}
	}
	close(jobs)

	<-done
	res, hmn := unoperate(operationMap["root"], valueMap, operationMap)
	res = dispelMystery(valueMap, operationMap, hmn)
	return res
}
