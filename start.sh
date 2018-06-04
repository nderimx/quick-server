#!/bin/bash

start-stop-daemon -bx ~/server/server -S

start-stop-daemon -bx ~/server/forward.sh -S
