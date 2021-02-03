#!/bin/bash
echo Building domainfinder...
go build -o domainfinderapp

echo Building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonymsapp

echo Building available...
cd ../available
go build -o ../domainfinder/lib/availableapp

echo Building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkleapp

echo Building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolifyapp

echo Building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainifyapp

echo Done.
