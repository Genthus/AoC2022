package main

import (
	"errors"
	"fmt"
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
	for {
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
		fmt.Println(o.name, ":", o.value1, v1, o.operation, o.value2, v2)

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
}

type value struct {
	name  string
	value int
}

func ParseString(input string) (*operation, *value) {
	parts := strings.Split(input, ": ")
	fmt.Println(parts)
	name := parts[0]
	data := parts[1]
	if strings.ContainsAny(data, "*/+-") {
		opData := strings.Split(data, " ")
		fmt.Println("optData", opData)
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
					if v.name != "humn" {
						valueMap[v.name] = v.value
					}
				} else if o != nil {
					if o.name != "humn" {
						operationMap[o.name] = o
					}
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
