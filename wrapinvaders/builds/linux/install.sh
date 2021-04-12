#!/bin/bash

if ! [[ -f $HOME/.config/wrapinvaders ]]; then
	mkdir $HOME/.config/wrapinvaders
	cp -r assets $HOME/.config/wrapinvaders
fi

sudo cp wrapinvaders /bin

wrapinvaders	
