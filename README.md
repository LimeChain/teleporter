## Overview

Teleporter is an EVM compatible cross-subnet communication protocol built on top of [Avalanche Warp Messaging (AWM)](https://docs.avax.network/learn/avalanche/awm), and implemented as a Solidity smart contract. It provides a mechanism to asynchronously invoke smart contract functions on other EVM blockchains within Avalanche. Teleporter provides a handful of useful features on top of AWM, such as specifying relayer incentives for message delivery, replay protection, message delivery and execution retries, and a standard interface for sending and receiving messages within a dApp deployed across multiple subnets.

It's important to understand the distinction between Avalanche Warp Messaging and Teleporter. AWM allows subnets to communicate with each other via authenticated messages by providing signing and verification primitives in Avalanchego. These are used by the blockchain VMs to sign outgoing messages and verify incoming messages.

The Teleporter protocol, on the other hand, is implemented at the smart contract level, and is a user-friendly interface to AWM, aimed at dApp developers. All of the message signing and verification is abstracted away from developers. Instead, developers simply call `sendCrossChainMessage` on the `TeleporterMessenger` contract to send a message invoking a smart contract on another subnet, and implement the `ITeleporterReceiver` interface to receive messages on the destination subnet. Teleporter handles all of the Warp message construction and sending, as well as the message delivery and execution.

- [Overview](#overview)
- [Setup](#setup)
  - [Initialize the repository](#initialize-the-repository)
  - [Dependencies](#dependencies)
- [Structure](#structure)
- [Run a local testnet in Docker](#run-a-local-testnet-in-docker)
  - [Start up the local testnet](#start-up-the-local-testnet)
  - [Additional notes](#additional-notes)
- [E2E tests](#e2e-tests)
  - [Run specific E2E tests](#run-specific-e2e-tests)
  - [Run the E2E tests on another network](#run-the-e2e-tests-on-another-network)
- [Deploy Teleporter to a Subnet](#deploy-teleporter-to-a-subnet)
- [ABI Bindings](#abi-bindings)
- [Docs](#docs)
- [Resources](#resources)

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests
- Docker and Docker Compose v2 for running the local test network
  - The docker image installs the following:
    - [Foundry](https://book.getfoundry.sh/getting-started/installation)
    - [Python3](https://www.python.org/downloads/)

## Structure

- `contracts/` is a [Foundry](https://github.com/foundry-rs/foundry) project that includes the implementation of the `TeleporterMessenger` contract and example dApps that demonstrate how to write contracts that interact with Teleporter.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `scripts/` includes bash scripts for interacting with Teleporter in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` performs Solidity and Golang linting.
  - `scripts/local/` includes scripts for running Teleporter in Docker.
- `docker/` includes configurations for a local, containerized setup of Teleporter.

## Run a local testnet in Docker

A docker setup for running a local network with Teleporter deployed is provided. This setup provides a convenient way to develop and test Teleporter as well as cross-chain applications built on top of Teleporter. Teleporter messages are relayed between subnets using [AWM Relayer](https://github.com/ava-labs/awm-relayer), a fully featured implementation of a Warp message relayer.

### Start up the local testnet

- Run `./scripts/local/run.sh` to run the local testnet in Docker containers with the ability to interact with the nodes directly.

  - `./scripts/local/run.sh` usage is as follows:

  ```
    -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub
    -p, --pause                       Pause the network on stop. Will attempt to restart the paused network on subsequent runs
    -h, --help                        Print this help message
  ```

  - If using `-l, --local` to use a local version of the `awm-relayer` image, build it using `./scripts/build_local_image.sh` from the root of the `awm-relayer` repository.
  - Note that if `-l, --local` is not set, then the latest published `awm-relayer` image will be pulled from Dockerhub.

- After calling `./scripts/local/run.sh`, you'll know the network is ready once the "Waiting for subnets to start up." messages from the `relayer_run` container stop. Once the network is ready, you can interact with the deployed Teleporter contracts directly, or deploy a cross-chain application contract such as `ExampleCrossChainMessenger`. To open a terminal in the container and initialize it with the environment variables needed to interact with the contracts, run:

```
# Open a shell in the container
docker exec -it local_network_run /bin/bash

# In the container:
set -a                        # export all variables so child processes can access
source vars.sh                # source the variables needed to interact with the contracts
```

- Once you've opened a shell in the container, try interacting with the network.

  - For example, send 1 AVAX on the C-Chain using `cast`

  ```
  c_address=0x333d17d3b42bf7930dbc6e852ca7bcf560a69003   # pick an arbitrary address
  cast balance --rpc-url $c_chain_rpc_url $c_address
  cast send --private-key $user_private_key --value 1 $c_address --rpc-url $c_chain_rpc_url
  cast balance --rpc-url $c_chain_rpc_url $c_address
  ```

  - An example of how to interact with Teleporter is provided in `scripts/local/examples/basic_send_receive.sh`. This script sends a dummy payload via Teleporter from the C-Chain to a subnet, and back again.

  ```
  ./scripts/local/examples/basic_send_receive.sh
  ```

  - You should see "Received on Subnet A is true" and "Received on the C-Chain is true" to indicate that Teleporter messages were successfully sent between the C-Chain and Subnet A.
    - These examples can be adapted to send messages between any two subnets, or between the C-Chain and any subnet by changing the RPC URLs.
    - Use these as a starting point to build and interact with your own cross-chain applications on top of Teleporter!

- The script `./scripts/local/run_stop.sh` should be used to gracefully shut down the containers, preserving the local network state between runs. This script is called automatically at the end of `./scripts/local/run.sh`, but can be called at any time from a separate terminal to pause the network.
  - `./scripts/local/run_stop.sh` usage is as follows:
  ```
  ./run_stop.sh             # stop the running containers and preserve the network for subsequent runs
  ./run_stop.sh -c          # stop the running containers and clean the network
  ```

### Additional notes

- The `./scripts/local/run.sh` script runs five local network nodes, with each of the nodes validating the primary network and three subnets (Subnet A, Subnet B, and Subnet C).
- Logs from the subnets on one of the five nodes are printed to stdout when run using either script.
- These logs can also be found at `~/.avalanche-cli/runs/network_<DATE>_<TIMESTAMP>/node{1,5]/logs/<SUBNET_ID>.log` in the `local_network_run` container.

## E2E tests

In addition to the docker setup, end-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests. See the [Contribution Guide](./CONTRIBUTING.md) for additional details.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/local/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Calculate Teleporter message IDs` test:

```bash
GINKGO_FOCUS="Calculate Teleporter message IDs" ./scripts/local/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="Calculate Teleporter" ./scripts/local/e2e_test.sh
```

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all E2E tests for the example cross chain applications:

```bash
	ginkgo.It("Send native tokens from subnet A to B and back",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.NativeTokenBridge(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="cross chain apps" ./scripts/local/e2e_test.sh
```

### Run the E2E tests on another network

The same E2E test flows can be executed against external network by setting the proper environment variables in `.env.testnet` and `.env`, and running the following commands:

```bash
cp .env.example .env # Set proper values after copying.
./scripts/testnet/run_testnet_e2e_flows.sh
```

The user wallet set in `.env` must have native tokens for each of the subnets used in order for the test flows to be able to send transactions on those networks.

## Deploy Teleporter to a Subnet

From the root of the repo, the TeleporterMessenger contract can be deployed by calling

```bash
./scripts/deploy_teleporter.sh <options>
```

Options for this script:

- `--version <version>` Required. Specify the release version to deploy. These will all be of the form `v1.X.0`. Each Teleporter version can only send and receive messages from the **same** Teleporter version on another chain. You can see a list of released versions at https://github.com/ava-labs/teleporter/releases.
- `--rpc-url <url>` Required. Specify the rpc url of the node to use.
- `--fund-deployer <private_key>` Optional. Funds the deployer address with the account held by `<private_key>`

To ensure that Teleporter can be deployed to the same address on every EVM based chain, it uses [Nick's Method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c) to deploy from a static deployer address. Teleporter costs exactly `10eth` in the subnet's native gas token to deploy, which must be sent to the deployer address.

`deploy_teleporter.sh` will send the necessary native tokens to the deployer address if it is provided with a private key for an account with sufficient funds. Alternatively, the deployer address can be funded externally. The deployer address for each version can be found by looking up the appropriate version at https://github.com/ava-labs/teleporter/releases and downloading `TeleporterMessenger_Deployer_Address_<VERSION>.txt`.

## ABI Bindings

To generate Golang ABI bindings for the Solidity smart contracts, run:

```bash
./scripts/abi_go_bindings.sh
```

The auto-generated bindings should be written under the `abi-bindings/` directory.

## Docs

- [Teleporter Protocol Overview](./contracts/src/Teleporter/README.md)
- [Cross Chain Applications](./contracts/src/CrossChainApplications/README.md)
- [Getting Started](./contracts/src/CrossChainApplications/GETTING_STARTED.md)
- [Teleporter Upgradeability](./contracts/src/Teleporter/upgrades/README.md)
- [Contract Deployment](./utils/contract-deployment/README.md)
- [Teleporter CLI](./cmd/teleporter-cli/README.md)

## Resources

- List of blockchain signing cryptography algorithms [here](http://ethanfast.com/top-crypto.html).
- Background on stateful precompiles [here](https://medium.com/avalancheavax/customizing-the-evm-with-stateful-precompiles-f44a34f39efd).
- Background on BLS signature aggregation [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html).
