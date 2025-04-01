#!/bin/bash

# Find all Go files and replace the old import paths with the new one
find . -type f -name "*.go" -exec sed -i '' 's|github.com/jacky-htg/inventory|github.com/nirshpaa/godam-backend|g' {} +
find . -type f -name "*.go" -exec sed -i '' 's|github.com/nishanpandit/inventory|github.com/nirshpaa/godam-backend|g' {} +

echo "Import paths updated successfully" 