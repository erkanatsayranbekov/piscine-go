#!/bin/bash
find -name '*.sh' | cut -d '/' -f 2 | sed s/^..// | sed  s/...$//
