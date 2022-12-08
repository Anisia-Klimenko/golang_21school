#!/bin/bash
go test -bench=. | grep Benchmark | awk '{print $1,$3}' | sort -r -nk2 | head | awk '{print $1}' | cut -d '-' -f 1 > top10.txt