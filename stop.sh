set -euo pipefail
script_dirpath="$(cd "$(dirname "${0}")" && pwd)"

cd "${script_dirpath}"
docker compose down
