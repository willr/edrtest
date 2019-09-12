#! /usr/bin/env bash

./edrtest exec -execPath /bin/cat all_tests.sh

./edrtest file create -filePath ./blah.txt
./edrtest file modify -filePath ./blah.txt
./edrtest file delete -filePath ./blah.txt

./edrtest net -url http://www.google.com
