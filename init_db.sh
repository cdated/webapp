#!/bin/sh

createdb chitchat
psql -f setup.sql -d chitchat
