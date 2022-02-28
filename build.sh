#!/bin/bash

go build -gcflags=all="-N -l"  -o bin .
