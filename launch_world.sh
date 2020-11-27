#!/bin/bash

./server/gnatsd  &

export GAZEBO_PLUGIN_PATH=${GAZEBO_PLUGIN_PATH}:./plugins
export GAZEBO_MODEL_PATH=${GAZEBO_MODEL_PATH}:./models
gazebo worlds/mopla.world --verbose
