#!/usr/bin/env bash
# Upgrades terraform-provider-zitadel to its latest upstream release, regenerates
# the schema and SDKs, and opens a pull request.
#
# This replaces pulumi/upgrade-provider, which cannot handle this upstream: it tags
# v3.x releases but its go.mod still declares the /v2 module path, and
# upgrade-provider always rewrites the module path to match the tag's major
# version (`go get .../v3@<sha>`), which Go rejects. Here the release commit is
# pinned as a /v2 pseudo-version instead.
#
# Environment:
#   GH_TOKEN        token for gh (required)
#   TARGET_VERSION  upstream release to upgrade to, e.g. v3.8.6 (default: latest)
#   PR_REVIEWERS    comma-separated GitHub users to request review from (optional)
set -euo pipefail

UPSTREAM_REPO=zitadel/terraform-provider-zitadel
UPSTREAM_MODULE=github.com/zitadel/terraform-provider-zitadel/v2

cd "$(dirname "$0")/.."

target=${TARGET_VERSION:-$(gh release view --repo "$UPSTREAM_REPO" --json tagName --jq .tagName)}
target=v${target#v}

# Prefer the peeled ref, so an annotated tag resolves to its commit.
refs=$(git ls-remote "https://github.com/$UPSTREAM_REPO" "refs/tags/$target" "refs/tags/$target^{}")
sha=$(awk '/\^\{\}$/ {print $1}' <<<"$refs")
sha=${sha:-$(awk '{print $1}' <<<"$refs")}
if [[ -z "$sha" ]]; then
    echo "error: upstream tag $target not found" >&2
    exit 1
fi

current=$(cd provider && go list -m -f '{{.Version}}' "$UPSTREAM_MODULE")
if [[ "${sha:0:12}" == "${current##*-}" ]]; then
    echo "Already on upstream $target ($current), nothing to do."
    exit 0
fi

branch="upgrade-terraform-provider-zitadel-to-$target"
if git ls-remote --exit-code --heads origin "$branch" >/dev/null; then
    echo "Branch $branch already exists on origin, skipping."
    exit 0
fi

echo "Upgrading upstream $current -> $target ($sha)"
git checkout -b "$branch"

(
    cd provider
    go get "$UPSTREAM_MODULE@$sha"
    go mod tidy
)
sed -i -E "s|// Upstream release v[0-9][^ ]*\. |// Upstream release $target. |" provider/go.mod

make tfgen
make build_sdks

git add -A
git commit -m "Upgrade terraform-provider-zitadel to $target"
git push --set-upstream origin "$branch"

pr_args=(
    --base main
    --head "$branch"
    --title "Upgrade terraform-provider-zitadel to $target"
    --body "Upgrades the upstream provider to [$target](https://github.com/$UPSTREAM_REPO/releases/tag/$target) and regenerates the schema and SDKs.

Upstream's go.mod still declares the /v2 module path, so the release commit is pinned as a /v2 pseudo-version (\`$UPSTREAM_MODULE@${sha:0:12}\`)."
)
if [[ -n "${PR_REVIEWERS:-}" ]]; then
    pr_args+=(--reviewer "$PR_REVIEWERS")
fi
gh pr create "${pr_args[@]}"
