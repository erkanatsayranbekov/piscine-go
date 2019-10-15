#!/bin/bash

 curl https://api.github.com/users/justauser123 | grep '"id":' | cut -d  " " -f4  | sed s/.$//

