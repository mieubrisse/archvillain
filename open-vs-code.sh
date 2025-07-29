set -euo pipefail
script_dirpath="$(cd "$(dirname "${0}")" && pwd)"

# Use 1Password to replace Github token
op run --env-file <(echo "GH_TOKEN=op://njr5tpkb6fbk6zgznwp5n2kb5y/m7idcfhrbople3n72i4u4j4xa4/notesPlain") "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code" "${script_dirpath}"
