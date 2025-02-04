# Changelog

## Unreleased (0.17.0)

### 🚨 Breaking changes
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🗑️  Deprecation
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🛠  Improvements
- [593](https://github.com/vegaprotocol/vegawallet/issues/593) - Tell users to aim for the core repository

### 🔥 Removal
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🐛 Fixes
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

## 0.16.1

### 🐛 Fixes
- Fix hardcoded version to match tag and build binaries

## 0.16.0

### 🐛 Fixes
- [575](https://github.com/vegaprotocol/vegawallet/pull/575) - Fix broken docs links in readme
- [579](https://github.com/vegaprotocol/vegawallet/pull/579) - Ensure the wallet doesn't recommend upgrade to pre-release version

## 0.15.1

### 🐛 Fixes
- [558](https://github.com/vegaprotocol/vegawallet/pull/558) - Use user-friendly RPC code as error message
- [569](https://github.com/vegaprotocol/vegawallet/pull/569) - Ensure the winfile scheme for logger is register only once

## 0.15.0

### 🛠  Improvements
- [555](https://github.com/vegaprotocol/vegawallet/issues/555) - Now using full transaction structures for event emit
- [551](https://github.com/vegaprotocol/vegawallet/pull/551) - Enable system-tests with Vegacapsule

### 🐛 Fixes
- [548](https://github.com/vegaprotocol/vegawallet/pull/548) - Printer utility shouldn't write in service response

## 0.14.2

### 🛠  Improvements
- [544](https://github.com/vegaprotocol/vegagwallet/issues/544) - Display transaction hash after sending transaction

## 0.14.1

### 🐛 Fixes
- [542](https://github.com/vegaprotocol/vegawallet/pull/542) - Update `protos` to get last version of governance schema


## 0.14.0

### 🚨 Breaking changes
- [539](https://github.com/vegaprotocol/vegawallet/issues/539) - Add received at timestamp to transaction sign requests
- [497](https://github.com/vegaprotocol/vegawallet/issues/497) - Add ability to confirm the signing of a transaction. To keep running the service with automatic consent, `--automatic-consent` flag is now required.

### 🗑️  Deprecation
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🛠  Improvements
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🔥 Removal
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🐛 Fixes
- [521](https://github.com/vegaprotocol/vegawallet/pull/521) - Update `protos` repository to get last validation updates on commands
- [523](https://github.com/vegaprotocol/vegawallet/pull/521) - Add link in `README` to direct people to desktop wallet app
- [529](https://github.com/vegaprotocol/vegawallet/pull/529) - Ensure log files can be created on Windows.
- [536](https://github.com/vegaprotocol/vegawallet/pull/536) - Docs: replace link for providing feedback

## 0.13.2

### 🚨 Breaking changes
- [504](https://github.com/vegaprotocol/vegawallet/pull/504) - Add support for proof of work sent from the wallet as part of a transaction
- [507](https://github.com/vegaprotocol/vegawallet/issues/507) - Flag `--level` in service run removed

### 🛠  Improvements
- [508](https://github.com/vegaprotocol/vegawallet/issues/508) - Introduce a command to list the service endpoints
- [507](https://github.com/vegaprotocol/vegawallet/issues/507) - Redirect logs from command service run to a file
- [516](https://github.com/vegaprotocol/vegawallet/issues/516) - Update protos for market update proposal

### 🐛 Fixes
- [502](https://github.com/vegaprotocol/vegawallet/pull/502) - Unwrap properly transfer and cancel transfer commands


## 0.13.1

### 🛠  Improvements
- [496](https://github.com/vegaprotocol/vegawallet/pull/496) - Return tx in /api/v1/command response


## 0.13.0

### 🛠  Improvements
- [484](https://github.com/vegaprotocol/vegawallet/pull/484) - Add JSON structure tags to handler requests, to match responses format

### 🐛 Fixes
- [487](https://github.com/vegaprotocol/vegawallet/pull/487) - Update to work with changes in `protos`

## 0.12.0

### 🚨 Breaking changes
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - The key printed on the JSON output for the `key generate` command is now at the root of the response, just like the output of the `key describe` command.
- [450](https://github.com/vegaprotocol/vegawallet/pull/450) - Move `sign` and `verify` commands to a a `message` subcommand, e.g. `message sign` and `message verify`

### 🗑️  Deprecation
- [](https://github.com/vegaprotocol/vegawallet/pull/) -

### 🛠  Improvements
- [427](https://github.com/vegaprotocol/vegawallet/pull/427) - Sign MacOS binaries
- [447](https://github.com/vegaprotocol/vegawallet/pull/447) - Sign Windows binaries
- [448](https://github.com/vegaprotocol/vegawallet/pull/448) - Add command `network locate` to locate the folder in which network configuration files are stored.
- [451](https://github.com/vegaprotocol/vegawallet/pull/451) - Verify network configuration prior exploitation to raise user-friendly errors.
- [435](https://github.com/vegaprotocol/vegawallet/pull/435) - Support signing commands offline with `command sign`
- [453](https://github.com/vegaprotocol/vegawallet/pull/453) - Added support for new liquidity provision commands
- [456](https://github.com/vegaprotocol/vegawallet/pull/456) - Update protos to support non-negative string pegged orders offset
- [463](https://github.com/vegaprotocol/vegawallet/pull/463) - Improve the `vegawallet version` command output to ensure users doesn't mistake the software version with their wallet version
- [443](https://github.com/vegaprotocol/vegawallet/pull/443) - Notarization of MacOS binaries
- [469](https://github.com/vegaprotocol/vegawallet/pull/469) - Support time-triggered trading termination from built-in oracle

### 🔥 Removal
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - Wallet creation with `key generate` has been removed in favor of `create` command
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - Flag `--console-proxy` has been removed in favor of `--with-console`
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - Flag `--mnemonic-file` has been removed in favor of `--recovery-phrase-file`
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - Field `mnemonic` in HTTP response for wallet creation has been removed in favor of `recoveryPhrase`
- [446](https://github.com/vegaprotocol/vegawallet/pull/446) - Field `mnemonic` in HTTP request for wallet import has been removed in favor of `recoveryPhrase`

### 🐛 Fixes
- [440](https://github.com/vegaprotocol/vegawallet/pull/440) - Fix confirmation prompt on windows
- [449](https://github.com/vegaprotocol/vegawallet/pull/449) - Don't use unicode glyphs on windows
- [444](https://github.com/vegaprotocol/vegawallet/pull/444) - Fail gracefully when trying to get password input on msys
- [455](https://github.com/vegaprotocol/vegawallet/pull/455) - New Liquidity Provisioning commands removed requirement for Id field
- [485](https://github.com/vegaprotocol/vegawallet/pull/485) - Update Windows VM version to fix code signing

## 0.11.0

### 🚨 Breaking changes
- [381](https://github.com/vegaprotocol/vegawallet/pull/381) - Make `send tx` and `send command` match `NOUN VERB` grammar
- [384](https://github.com/vegaprotocol/vegawallet/pull/384) - Remove private key from key generate JSON output

### 🗑️  Deprecation
- [418](https://github.com/vegaprotocol/vegawallet/pull/418) - Wallet creation is deprecated with `key generate`, use `create` instead
- [423](https://github.com/vegaprotocol/vegawallet/pull/423) - Flag `--console-proxy` is deprecated in favor of `--with-console`
- [425](https://github.com/vegaprotocol/vegawallet/pull/425) - Flag `--mnemonic-file` is deprecated in favor of `--recovery-phrase-file`
- [425](https://github.com/vegaprotocol/vegawallet/pull/425) - Deprecated `mnemonic` field in HTTP response for wallet creation in favor of `recoveryPhrase`
- [425](https://github.com/vegaprotocol/vegawallet/pull/425) - Deprecated `mnemonic` field in HTTP request for wallet import in favor of `recoveryPhrase`

### 🛠  Improvements
- [383](https://github.com/vegaprotocol/vegawallet/pull/383) - Add wallet version to `key generate` output
- [390](https://github.com/vegaprotocol/vegawallet/pull/390) - Add new project board Github action and update linked PR action version
- [385](https://github.com/vegaprotocol/vegawallet/pull/385) - Add command `delete` to delete a wallet
- [398](https://github.com/vegaprotocol/vegawallet/pull/398) - More idiomatic test suite
- [396](https://github.com/vegaprotocol/vegawallet/pull/396) - Update the "important" section of `key generate`
- [405](https://github.com/vegaprotocol/vegawallet/pull/405) - Add command `network delete` to delete a network
- [409](https://github.com/vegaprotocol/vegawallet/pull/409) - Return the tx hash after sending a transaction
- [406](https://github.com/vegaprotocol/vegawallet/pull/406) - Add auto-completion to command line through `completion` sub-command
- [418](https://github.com/vegaprotocol/vegawallet/pull/418) - Importing a wallet now generates the first key pair
- [418](https://github.com/vegaprotocol/vegawallet/pull/418) - Add the `create` command to create a wallet
- [420](https://github.com/vegaprotocol/vegawallet/pull/420) - Verify initialisation state on service run
- [421](https://github.com/vegaprotocol/vegawallet/pull/421) - Improve help texts
- [423](https://github.com/vegaprotocol/vegawallet/pull/423) - Support launching the token dApp in `service run` command

### 🔥 Removal
- [387](https://github.com/vegaprotocol/vegawallet/pull/387) - Remove network initialisation in favor of `network import`

### 🐛 Fixes
- [395](https://github.com/vegaprotocol/vegawallet/pull/395) - Fix `command send` examples
- [401](https://github.com/vegaprotocol/vegawallet/pull/400) - Update key rotation submission command
- [410](https://github.com/vegaprotocol/vegawallet/pull/410) - Add current pub key hash to key rotation submission command
- [415](https://github.com/vegaprotocol/vegawallet/pull/415) - Update to latest key rotation protos

## 0.10.0
*2021-11-15*

### 🚨 Breaking changes
- [300](https://github.com/vegaprotocol/vegawallet/pull/300) - Move from semi-colon to comma separated metadata on `key annotate` and `key generate`
- [300](https://github.com/vegaprotocol/vegawallet/pull/300) - Trim the `key list` output to only output name and public key
- [358](https://github.com/vegaprotocol/vegawallet/pull/358) - Rename "human" output to "interactive"

### 🗑️ Deprecation
- [368](https://github.com/vegaprotocol/vegawallet/pull/368) - Deprecate `command` subcommand in favor of `send command`

### 🛠 Improvements
- [300](https://github.com/vegaprotocol/vegawallet/pull/300) - Internal restructuring of CLI building blocks to improve testability
- [359](https://github.com/vegaprotocol/vegawallet/pull/359) - Add `key describe` subcommand to allow querying of key information
- [364](https://github.com/vegaprotocol/vegawallet/pull/364) - Add changelog checker Github action
- [361](https://github.com/vegaprotocol/vegawallet/pull/361) - Add `network describe` subcommand to allow querying of network information
- [366](https://github.com/vegaprotocol/vegawallet/pull/366) - Update the changelog checker action to run when required only
- [368](https://github.com/vegaprotocol/vegawallet/pull/368) - Add `send tx` subcommand to send raw transaction
- [367](https://github.com/vegaprotocol/vegawallet/pull/367) - Add `key rotate` subcommand to generate key rotation transaction payload

### 🐛 Fixes
- [356](https://github.com/vegaprotocol/vegawallet/pull/356) - Ensure the interactive printer is listening to color management env vars (`NO_COLOR` and `CLICOLOR`)
- [357](https://github.com/vegaprotocol/vegawallet/pull/357) - Warn without failing when no connection on version verification
- [360](https://github.com/vegaprotocol/vegawallet/pull/360) - Attempt to enable ANSI colors in Windows terminal
- [372](https://github.com/vegaprotocol/vegawallet/pull/372) - Remove wallet name serialisation to allow wallet file to be renamed
- [370](https://github.com/vegaprotocol/vegawallet/pull/370) - Ensure unsupported output is correctly reported

## 0.9.2
*2021-10-27*

### 🛠 Improvements
- [319](https://github.com/vegaprotocol/vegawallet/pull/319) - Enable more golanci lints
- [325](https://github.com/vegaprotocol/vegawallet/pull/325) - Importing a wallet via API default to version 2
- [322](https://github.com/vegaprotocol/vegawallet/pull/322) - Build artifact for mac M1


### 🐛 Fixes
- [323](https://github.com/vegaprotocol/vegawallet/pull/323) - Fix wallets with version 0
- [308](https://github.com/vegaprotocol/vegawallet/pull/308) - Fix command subcommand
- [330](https://github.com/vegaprotocol/vegawallet/pull/330) - Fix output for list keys command


## 0.9.1
*2021-10-23*

### 🐛 Fixes
- [313](https://github.com/vegaprotocol/go-wallet/pull/313) - Fix release github action


## 0.9.0
*2021-10-23*

### Migration from v0.8.0
* Flag `--name` and `-n` has been replaced by `--wallet` and `-w` respectively.
* The service configuration `wallet-service/config.toml` no longer exists.
* The network configurations are located in the `wallet-service/networks` config folder.
* A new `--network` (shorthand `-n`) has been introduced on `command` and `service run` subcommands.
* In the network configuration the `Nodes` configuration in now named `API.GRPC`.


### 🛠 Improvements
- [310](https://github.com/vegaprotocol/go-wallet/pull/310) - Rename repository to vegawallet
- [306](https://github.com/vegaprotocol/go-wallet/pull/301) - Implement hd wallet v2 (new derivation path)
- [301](https://github.com/vegaprotocol/go-wallet/pull/301) - Remove download endpoint
- [293](https://github.com/vegaprotocol/go-wallet/pull/293) - Better versioning
- [292](https://github.com/vegaprotocol/go-wallet/pull/292) - Expose data node endpoints
- [288](https://github.com/vegaprotocol/go-wallet/pull/288) - Update shared library
- [280](https://github.com/vegaprotocol/go-wallet/pull/280) - Update service documentation
- [281](https://github.com/vegaprotocol/go-wallet/pull/281) - Ignore system-tests failures
- [279](https://github.com/vegaprotocol/go-wallet/pull/279) - Run system-tests in CI pipeline
- [273](https://github.com/vegaprotocol/go-wallet/pull/273) - Run golangci-lint as part of Jenkins pipeline
- [272](https://github.com/vegaprotocol/go-wallet/pull/272) - Update the README
- [271](https://github.com/vegaprotocol/go-wallet/pull/271) - Update CLI documentation with network commands
- [174](https://github.com/vegaprotocol/go-wallet/pull/174) - Add golangci-lint
- [268](https://github.com/vegaprotocol/go-wallet/pull/268) - Update build instructions
- [253](https://github.com/vegaprotocol/go-wallet/pull/253) - Support multiple networks
- [257](https://github.com/vegaprotocol/go-wallet/pull/257) - Remove custom requirement verification in favor of cobra built-in's
- [260](https://github.com/vegaprotocol/go-wallet/pull/260) - Support importing networks from local file and URL
- [256](https://github.com/vegaprotocol/go-wallet/pull/256) - Remove getting started as it's the duty of the documentation website
- [244](https://github.com/vegaprotocol/go-wallet/pull/244) - Support key rotation
- [252](https://github.com/vegaprotocol/go-wallet/pull/252) - Readme copy change
- [250](https://github.com/vegaprotocol/go-wallet/pull/250) - Update getting started link to go to docs site
- [247](https://github.com/vegaprotocol/go-wallet/pull/247) - Config endpoint
- [248](https://github.com/vegaprotocol/go-wallet/pull/248) - Update proto services with last changes
- [241](https://github.com/vegaprotocol/go-wallet/pull/241) - SubmitTransactionV2 -> SubmitTransaction
- [234](https://github.com/vegaprotocol/go-wallet/pull/234) - Save XDG file structure
- [235](https://github.com/vegaprotocol/go-wallet/pull/235) - Add command subcommand to the wallet
- [226](https://github.com/vegaprotocol/go-wallet/pull/226) - Clean up the printed service endoint list
- [223](https://github.com/vegaprotocol/go-wallet/pull/223) - Add delegate and undelegate in commands
- [222](https://github.com/vegaprotocol/go-wallet/pull/222) - Ensure mnemonic files can be loaded from the same directory as execution
- [219](https://github.com/vegaprotocol/go-wallet/pull/219) - Update to last protos
- [217](https://github.com/vegaprotocol/go-wallet/pull/217) - Update protos dep
- [216](https://github.com/vegaprotocol/go-wallet/pull/216) - Update dependencies
- [213](https://github.com/vegaprotocol/go-wallet/pull/213) - Allow verification of message from which you don't have the private key
- [210](https://github.com/vegaprotocol/go-wallet/pull/210) - Replace `--mnemonic` flag by `--mnemonic-path`
- [189](https://github.com/vegaprotocol/go-wallet/pull/189) - Use TradingClient only
- [212](https://github.com/vegaprotocol/go-wallet/pull/212) - Shrink docker image size
- [208](https://github.com/vegaprotocol/go-wallet/pull/208) - Add flag to clear metadata
- [206](https://github.com/vegaprotocol/go-wallet/pull/206) - Self documentation of CLI
- [204](https://github.com/vegaprotocol/go-wallet/pull/204) - Improve service store API to get file information
- [193](https://github.com/vegaprotocol/go-wallet/pull/193) - Remove logger from service.GenerateConfig
- [196](https://github.com/vegaprotocol/go-wallet/pull/196) - Make untainting a key possible
- [195](https://github.com/vegaprotocol/go-wallet/pull/195) - Rename meta subcommand by annotate for clarity
- [194](https://github.com/vegaprotocol/go-wallet/pull/194) - Document list subcommand
- [188](https://github.com/vegaprotocol/go-wallet/pull/188) - Add command to list all registered wallets
- [186](https://github.com/vegaprotocol/go-wallet/pull/186) - Bump proto version
- [185](https://github.com/vegaprotocol/go-wallet/pull/185) - Updated protos module
- [183](https://github.com/vegaprotocol/go-wallet/pull/183) - Update protos to point to develop
- [181](https://github.com/vegaprotocol/go-wallet/pull/181) - Remove Txv1
- [180](https://github.com/vegaprotocol/go-wallet/pull/180) - Allow update of the wallet name
- [178](https://github.com/vegaprotocol/go-wallet/pull/178) - Add Jenkinsfile
- [176](https://github.com/vegaprotocol/go-wallet/pull/176) - Expose Console service ouside of the command line
- [173](https://github.com/vegaprotocol/go-wallet/pull/173) - Remove now-useless buf files
- [172](https://github.com/vegaprotocol/go-wallet/pull/172) - Build docker image
- [169](https://github.com/vegaprotocol/go-wallet/pull/169) - Use the protos repo for protos + commands validations
- [168](https://github.com/vegaprotocol/go-wallet/pull/168) - Ensure the the pubKey is not double encoded in Txv2
- [166](https://github.com/vegaprotocol/go-wallet/pull/166) - Remove short flag for passphrase
- [164](https://github.com/vegaprotocol/go-wallet/pull/164) - Allow loading password from file
- [163](https://github.com/vegaprotocol/go-wallet/pull/163) - Add default name to generated keys
- [161](https://github.com/vegaprotocol/go-wallet/pull/161) - Configure logger level based on the config
- [159](https://github.com/vegaprotocol/go-wallet/pull/159) - Restrict file access
- [157](https://github.com/vegaprotocol/go-wallet/pull/157) - Adding delegation commands
- [162](https://github.com/vegaprotocol/go-wallet/pull/162) - Updated the proposal submission to add the trading termination oracle
- [160](https://github.com/vegaprotocol/go-wallet/pull/160) - Add short flag `-r` for root-path
- [156](https://github.com/vegaprotocol/go-wallet/pull/156) - Update readme to link to master for instructions
- [149](https://github.com/vegaprotocol/go-wallet/pull/149) - Split wallet and service management at store level
- [152](https://github.com/vegaprotocol/go-wallet/pull/152) - Fix broken getting_started link
- [145](https://github.com/vegaprotocol/go-wallet/pull/145) - Introduce an `init` command
- [140](https://github.com/vegaprotocol/go-wallet/pull/140) - Initialise store before handler in CLI
- [137](https://github.com/vegaprotocol/go-wallet/pull/137) - Update proposalSubmission validation to be >= 0
- [135](https://github.com/vegaprotocol/go-wallet/pull/135) - Add endpoint that returns app version info
- [133](https://github.com/vegaprotocol/go-wallet/pull/133) - Reorganise the readme to reincorporate starter instructions
- [132](https://github.com/vegaprotocol/go-wallet/pull/132) - Import all core protobuf in the vega wallet repo
- [130](https://github.com/vegaprotocol/go-wallet/pull/130) - Update `pr_verify_linked_issue.yml`
- [131](https://github.com/vegaprotocol/go-wallet/pull/131) - Uniformise the wallet cmd names
- [129](https://github.com/vegaprotocol/go-wallet/pull/129) - Restructure the command line to scope key management
- [112](https://github.com/vegaprotocol/go-wallet/pull/112) - Support for HD Wallet
- [122](https://github.com/vegaprotocol/go-wallet/pull/122) - Create `pr_verify_linked_issue.yml`


### 🐛 Fixes
- [297](https://github.com/vegaprotocol/go-wallet/pull/297) - Fix command help
- [290](https://github.com/vegaprotocol/go-wallet/pull/290) - Fix key generate example #290
- [285](https://github.com/vegaprotocol/go-wallet/pull/285) - Fix commit hash
- [270](https://github.com/vegaprotocol/go-wallet/pull/270) - Fix file extension stripping on network list
- [262](https://github.com/vegaprotocol/go-wallet/pull/262) - Fix golangci-lint offences
- [240](https://github.com/vegaprotocol/go-wallet/pull/240) - Fix Docker image to include node package
- [225](https://github.com/vegaprotocol/go-wallet/pull/225) - Fix wallet check
- [201](https://github.com/vegaprotocol/go-wallet/pull/201) - Remove extra "it"
- [154](https://github.com/vegaprotocol/go-wallet/pull/154) - Remove hyphen


## 0.8.0
*2021-07-06*

### 🛠 Improvements
- [105](https://github.com/vegaprotocol/go-wallet/pull/105) - Update `new_issues_to_sprint_project_board.yml`
- [106](https://github.com/vegaprotocol/go-wallet/pull/106) - Port TxV2 to wallet
- [109](https://github.com/vegaprotocol/go-wallet/pull/109) - Rework storage layer
- [111](https://github.com/vegaprotocol/go-wallet/pull/111) - Add verify endpoint
- [108](https://github.com/vegaprotocol/go-wallet/pull/108) - Add `command/*` into the startup dump docs
- [120](https://github.com/vegaprotocol/go-wallet/pull/120) - Release to master



## 0.7.0
*2021-06-11*

### 🛠 Improvements
- [96](https://github.com/vegaprotocol/go-wallet/pull/96) - Replace localhost with IP address
- [98](https://github.com/vegaprotocol/go-wallet/pull/98) - Add `new_issues_to_sprint_project_board.yml`
- [99](https://github.com/vegaprotocol/go-wallet/pull/99) - Add blockHeight
- [101](https://github.com/vegaprotocol/go-wallet/pull/101) - Log and return better errors from the nodes
- [103](https://github.com/vegaprotocol/go-wallet/pull/103) - Release `v0.7.0`

### 🐛 Fixes
- [92](https://github.com/vegaprotocol/go-wallet/pull/92) - Fix file name in readme text


## 0.6.8
*2021-05-27*

### Improvements
- [89](https://github.com/vegaprotocol/go-wallet/pull/89) - Return if error happen on Health method
- [91](https://github.com/vegaprotocol/go-wallet/pull/91) - Update default configs
- [94](https://github.com/vegaprotocol/go-wallet/pull/94) - Release `v0.6.8`


## 0.6.7
*2021-05-15*

### 🛠 Improvements
- [87](https://github.com/vegaprotocol/go-wallet/pull/87) - Release `v0.6.7` update vega API version


## 0.6.6
*2021-04-23*

### 🛠 Improvements
- [83](https://github.com/vegaprotocol/go-wallet/pull/83) - Check go-wallet newer version when runnig commands
- [86](https://github.com/vegaprotocol/go-wallet/pull/86) - Release `v0.6.6` update vega API version


### 🐛 Fixes
- [82](https://github.com/vegaprotocol/go-wallet/pull/82) - Make the healthcheck more useful and fix some of the endpoint docs printed at startup


## 0.6.5
*2021-04-08*

### 🛠 Improvements
- [72](https://github.com/vegaprotocol/go-wallet/pull/72) - Add endpoint to sign abitrary data
- [74](https://github.com/vegaprotocol/go-wallet/pull/74) - Add arbitray signing
- [80](https://github.com/vegaprotocol/go-wallet/pull/80) - Release `v0.6.5` update the vega API version


## 0.6.4
*2021-03-17*

### 🛠 Improvements
- [71](https://github.com/vegaprotocol/go-wallet/pull/71) - Release `v0.6.4` update vega API version


## 0.6.3
*2021-03-16*

### 🛠 Improvements
- [70](https://github.com/vegaprotocol/go-wallet/pull/70) - Rename api-clients to api


## 0.6.2
*2021-02-18*

### 🛠 Improvements
- [68](https://github.com/vegaprotocol/go-wallet/pull/68) - Add default meta
- [69](https://github.com/vegaprotocol/go-wallet/pull/69) - Release develop to master


## 0.6.1
*2021-02-17*

### 🛠 Improvements
- [64](https://github.com/vegaprotocol/go-wallet/pull/64) - Use `syscall.Stdin` for for read password
- [65](https://github.com/vegaprotocol/go-wallet/pull/65) - Updates for Windows instructions
- [66](https://github.com/vegaprotocol/go-wallet/pull/66) - Release patch version `v0.6.1`


### 🐛 Fixes
- [63](https://github.com/vegaprotocol/go-wallet/pull/63) - Fix windows command


## 0.6.0
*2021-02-15*

### 🛠 Improvements
- [61](https://github.com/vegaprotocol/go-wallet/pull/61) - Update the api to use the new protobugs
- [62](https://github.com/vegaprotocol/go-wallet/pull/62) - Release `v0.6.0`


## 0.5.2
*2020-12-17*

### 🛠 Improvements
- [54](https://github.com/vegaprotocol/go-wallet/pull/54) - Add user friendly at startup of the wallet
- [55](https://github.com/vegaprotocol/go-wallet/pull/55) - Update master with last develop changes


## 0.5.1
*2020-12-11*

### 🛠 Improvements
- [52](https://github.com/vegaprotocol/go-wallet/pull/52) - Implement round robin over all the network nodes when node forwarding
- [53](https://github.com/vegaprotocol/go-wallet/pull/53) - Update master with last develop


## 0.5.0
*2020-12-10*

### 🛠 Improvements
- [26](https://github.com/vegaprotocol/go-wallet/pull/26) - Add custom headers
- [43](https://github.com/vegaprotocol/go-wallet/pull/43) - Update readme - tidy up, combine
- [45](https://github.com/vegaprotocol/go-wallet/pull/45) - Add subheadings
- [49](https://github.com/vegaprotocol/go-wallet/pull/49) -Cchange default port to 1847 and IP is not Host in config
- [50](https://github.com/vegaprotocol/go-wallet/pull/50) - uUdates based on live trial session
- [51](https://github.com/vegaprotocol/go-wallet/pull/51) - Update master with last develop changes


## 0.4.3
*2020-11-20*

### 🛠 Improvements
- [39](https://github.com/vegaprotocol/go-wallet/pull/39) - Change users default home dir for windows
- [40](https://github.com/vegaprotocol/go-wallet/pull/40) - Release develop to master


## 0.4.2
*2020-11-20*

### 🛠 Improvements
- [36](https://github.com/vegaprotocol/go-wallet/pull/36) - Use 7zip to build release
- [37](https://github.com/vegaprotocol/go-wallet/pull/37) - Release develop into master


 ## 0.4.1
*2020-11-20*

### 🛠 Improvements
- [35](https://github.com/vegaprotocol/go-wallet/pull/35) - Release develop to master

### 🐛 Fixes
- [34](https://github.com/vegaprotocol/go-wallet/pull/34) - Windows fixes


## 0.4.0
*2020-11-20*

### 🛠 Improvements
- [21](https://github.com/vegaprotocol/go-wallet/pull/21) - Small README updates
- [24](https://github.com/vegaprotocol/go-wallet/pull/24) - Update the wallet with last vega changes
- [25](https://github.com/vegaprotocol/go-wallet/pull/25) - Remove native window
- [28](https://github.com/vegaprotocol/go-wallet/pull/28) - Release develop to master
- [30](https://github.com/vegaprotocol/go-wallet/pull/30) - Release develop to master


### 🐛 Fixes
- [27](https://github.com/vegaprotocol/go-wallet/pull/27) - Fix version setup in CI release
- [29](https://github.com/vegaprotocol/go-wallet/pull/29) - Fixes for windows release


## 0.3.1
*2020-11-16*

### 🛠 Improvements
- [19](https://github.com/vegaprotocol/go-wallet/pull/19) - cmdline quality of life improvements
- [20](https://github.com/vegaprotocol/go-wallet/pull/20) - Release develop to master


## 0.3.0
*2020-11-16*

### 🛠 Improvements
- [16](https://github.com/vegaprotocol/go-wallet/pull/16) - Generate RSA key default
- [17](https://github.com/vegaprotocol/go-wallet/pull/17) - Start work on native browser UI and try to build release
- [18](https://github.com/vegaprotocol/go-wallet/pull/18) - Release develop code to master


## 0.2.0
*2020-11-14*

### 🛠 Improvements
- [13](https://github.com/vegaprotocol/go-wallet/pull/13) - Update standalone wallet with last vega changes
- [14](https://github.com/vegaprotocol/go-wallet/pull/14) - Add an http server proxy to the console and start it with a flag
- [15](https://github.com/vegaprotocol/go-wallet/pull/15) - Release `v0.2.0`


## 0.1.0
*2020-05-11*

### 🛠 Improvements
- [5](https://github.com/vegaprotocol/go-wallet/pull/5) - WIP release on tag
- [4](https://github.com/vegaprotocol/go-wallet/pull/4) - Create go yaml workflow
- [3](https://github.com/vegaprotocol/go-wallet/pull/3) - Add version to the vega wallet cmdline
