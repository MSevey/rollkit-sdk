# Rollkit SDK

RappDK stands for "Rollkit application development kit". RappDK provides a boilerplate for building rollkit app using cosmos-sdk which introduces some new modules designed specifically for rollkit application as well as a default template and recommended setting for the app. The goal here is to "rollkitifize" the cosmos sdk, making the "PoS-chain oriented" cosmos sdk be more suitable for rollkit.

- A default [app template](./simapp/README.md)
- Recommended params for the app's modules
- New modules for rollkit app:
  - [Staking wrapper module](./x/staking/README.md)
  - [Sequencer module](./x/sequencer/README.md)
- Guide to kickstart an example rolllkit application
