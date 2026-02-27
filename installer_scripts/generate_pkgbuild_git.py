#!/usr/bin/env python3
from pathlib import Path
from textwrap import dedent
import subprocess

pkgname = "chibi-cli-git"
_pkgname = "chibi-cli"
pkgdesc = "Chibi for AniList - A lightweight anime & manga tracker CLI app powered by AniList (Upstream GIT)."
arch = ["x86_64"]
url = "https://chibi-cli.pages.dev/"
git_url = "https://github.com/CosmicPredator/chibi-cli"
license_ = ["GPL3"]
depends = ["glibc"]
makedepends = ["git", "go>=1.25"]
provides = ["chibi"]
conflicts = ["chibi", "chibi-cli-bin"]
options = ["!debug"]

PKGBUILD_TEMPLATE = """
pkgname={pkgname}
_pkgname={_pkgname}
pkgver={pkgver}
pkgrel=1
pkgdesc=\"{pkgdesc}\"
arch=({arch})
url=\"{url}\"
git_url=\"{git_url}\"
license=({license_})
depends=({depends})
makedepends=({makedepends})
provides=({provides})
conflicts=({conflicts})
options=({options})

source=(\"git+$git_url.git\")
sha256sums=('SKIP')

build() {{
  cd \"$srcdir/$_pkgname\"

  LDFLAGS=\"-X main.VERSION=${{pkgver}} -s -w\"

  GOOS=linux GOARCH=amd64 \\
    go build -ldflags=\"$LDFLAGS\" -o chibi
}}

package() {{
  install -Dm755 \"$srcdir/$_pkgname/chibi\" \"$pkgdir/usr/bin/chibi\"
}}
"""


def bash_array(items):
    return "'" + "' '".join(items) + "'"


def get_pkgver_from_git() -> str:
    """Derive pkgver from git describe in the current repo."""
    result = subprocess.run(
        ["git", "describe", "--tags", "--long"],
        check=True,
        capture_output=True,
        text=True,
    )
    desc = result.stdout.strip()
    if desc.startswith("v"):
        desc = desc[1:]
    desc = desc.replace("-", ".")
    return desc


def generate_pkgbuild(output_path: Path):
    pkgver_val = get_pkgver_from_git()
    content = PKGBUILD_TEMPLATE.format(
        pkgver=pkgver_val,
        pkgname=pkgname,
        _pkgname=_pkgname,
        pkgdesc=pkgdesc,
        arch=bash_array(arch),
        url=url,
        git_url=git_url,
        license_=bash_array(license_),
        depends=bash_array(depends),
        makedepends=bash_array(makedepends),
        provides=bash_array(provides),
        conflicts=bash_array(conflicts),
        options=bash_array(options),
    )

    output_path.write_text(dedent(content).lstrip())
    print(f"PKGBUILD generated at: {output_path}")


if __name__ == "__main__":
    generate_pkgbuild(Path("PKGBUILD"))
