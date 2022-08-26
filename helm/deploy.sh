#!/usr/bin/env sh

kubectl create ns otus-msa-hw2
helm upgrade --install -n otus-msa-hw2 otus-msa-hw2 helm/chart