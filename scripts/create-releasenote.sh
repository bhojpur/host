#!/bin/bash

# Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
#
# This script will create a txt file with Kubernetes versions, which will be
# used as (pre) release decription by Drone

set -e -x

RELEASEFILE="./build/bin/hostops-k8sversions.txt"

mkdir -p ./build/bin

echo "Creating ${RELEASEFILE}"

DEFAULT_VERSION=$(./bin/hostops --quiet config --list-version)
if [ $? -ne 0 ]; then
  echo "Non zero exit code while running 'hostops config -l'"
  exit 1
fi

DEFAULT_VERSION_FOUND="false"
echo "# Bhojpur Kubernetes Engine versions" > $RELEASEFILE
for VERSION in $(./bin/hostops --quiet config --all --list-version | sort -V); do
  if [ "$VERSION" == "$DEFAULT_VERSION" ]; then
    echo "- \`${VERSION}\` (default)" >> $RELEASEFILE
    DEFAULT_VERSION_FOUND="true"
  else
    echo "- \`${VERSION}\`" >> $RELEASEFILE
  fi
done

if [ "$DEFAULT_VERSION_FOUND" == "false" ]; then
  echo -e "\nNo default version found!" >> $RELEASEFILE
fi

echo "Done creating ${RELEASEFILE}"

cat $RELEASEFILE