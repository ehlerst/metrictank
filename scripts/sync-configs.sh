#!/bin/bash

echo "first make sure metrictank-sample.ini is up to date. its values should match the defaults used by metrictank. and comments should match the descriptions provided by metrictank help menus"
echo "now we will run vimdiff to manually synchronize updates from sample config to other configs:"
echo "try to make every config as closely resembling the sample config as possible, while retaining the customisations that makes each config unique"
echo "hit a key when you're ready...."
read
for cfg in $(find . -name 'metrictank*.ini' | grep -v '\./metrictank-sample.ini'); do
	vimdiff metrictank-sample.ini $cfg
done

echo "updating docs/config.md"
./scripts/config-to-doc.sh > docs/config.md

# metrics2docs .> docs/metrics.md this doesn't work very well yet
