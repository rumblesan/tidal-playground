# Tidal Install

A hopefully simple way to get tidal and tidal-midi installed and running.

## Setup

* Make sure that `stack` is installed and available on the command line.
* Either run the `setup` bash script or directly run `stack setup --install-ghc && stack build`
* Run `tidal` to start up the ghci session, or directly run `stack exec ghci -- -ghci-script Setup.ghci`

Might take a while the first time as everything is downloaded, should be relatively quick after that.

### Stack Install

The instructions to install stack are found on their website.
[Install Stack](https://docs.haskellstack.org/en/stable/install_and_upgrade/)

Installing stack should not require a pre-installed version of haskell, ghc, cabal or the haskell-platform.
