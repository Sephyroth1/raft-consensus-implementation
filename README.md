# Raft Consensus Algorithm

A from-scratch implementation of the **Raft consensus algorithm**
based on the original Stanford paper "In Search of an Understandable Consensus Algorithm"
with the goal of understanding and implementing distributed consensus, leader election
and reliable communication between nodes.

## Overview

Raft is a consensus algorithm designed to help a group of distributed nodes agree
on a shared state, even when individual nodes experience failures.

This project is an independent implementation of Raft,
built to explore the internal mechanics of distributed systems, including node state,
leader election, heartbeat communication, RPC-based interactions, and coordination between nodes.

The implementation is being developed incrementally, with an emphasis on understanding the protocol,
making deliberate architectural decisions, and validating behavior through testing.

## Goals

- Implement the core mechanics of the Raft consensus algorithm.
- Model the lifecycle and state transitions of individual Raft nodes.
- Implement leader election and heartbeat communication.
- Establish RPC-based communication between nodes.
- Build the main event loop that coordinates protocol behavior.
- Test the implementation under different configurations and failure scenarios.
- Develop a deeper understanding of distributed systems and fault tolerance.

## Architecture

The implementation is organized around individual Raft nodes and a communication layer that allows them to interact over a network.

### Raft Node

Each `RaftNode` represents an individual participant in the distributed system.

The node structure currently includes:

- **Node specification and state:** Holds the node's configuration and protocol-related state.
- **Election timers:** Support the timing mechanisms required for leader election.
- **Heartbeat timers:** Support periodic communication associated with leader activity.
- **RPC channel:** Provides the communication interface used to exchange messages with other nodes.
- **Role information:** Tracks the node's role within the Raft protocol.

### Communication Layer

A wrapper around the node is responsible for registering it with the network as an RPC service.

This separates the node's protocol state from the mechanism used to expose it for communication,
providing a foundation for interaction between distributed nodes.

### Main Event Loop

The main event loop is the next major implementation milestone.
It will coordinate protocol activity, including handling events, managing timers, processing communication, and driving node state transitions.

The event loop is essential to bringing the existing node structure and communication layer together into a functioning Raft implementation.

## Current Status

**Work in progress**

| Component | Status |
| --- | --- |
| Raft node structure | Implemented |
| Node state and role representation | Implemented |
| Election timer structure | Implemented |
| Heartbeat timer structure | Implemented |
| RPC communication layer | Implemented |
| RPC service registration wrapper | Implemented |
| Main event loop | In progress |
| End-to-end protocol execution | Pending |
| Systematic distributed failure testing | Pending |

The current implementation establishes the node structure and communication foundation.
Full protocol execution and rigorous validation remain necessary before the implementation can be considered complete.

## Implementation Approach

The project follows an incremental development approach:

1. **Understand the protocol:** Study the original Raft paper and identify the responsibilities and interactions of individual nodes.
2. **Design the node structure:** Represent node state, roles, timers, and communication interfaces.
3. **Build the communication layer:** Establish RPC communication and register nodes as network-accessible services.
4. **Implement the event loop:** Connect timers, incoming messages, and protocol state transitions.
5. **Integrate the protocol:** Bring the individual components together to support coordinated Raft behavior.
6. **Validate the implementation:** Test normal execution, different node configurations, and distributed failure scenarios.

## Testing Strategy

Testing is a critical part of validating a distributed consensus implementation.
Once the main event loop is complete, testing will focus on both isolated components and the behavior of the complete system.

Planned areas of testing include:

- **Unit testing:** Verify individual components, state transitions, timer behavior, and message handling.
- **Integration testing:** Validate communication and interactions between multiple Raft nodes.
- **Configuration testing:** Exercise the protocol with different node configurations and initial conditions.
- **Failure testing:** Investigate behavior during node failures, communication interruptions, and timing variations.
- **Concurrency testing:** Look for race conditions and timing-sensitive bugs that may not appear during ordinary execution.
- **Repeated execution:** Run tests across varied configurations to improve confidence in protocol correctness and reliability.

The implementation has not yet undergone systematic distributed failure testing, so fault tolerance and correctness under failure are not claimed as verified.

## Design Principles

- **Protocol-driven design:** Keep implementation decisions aligned with the Raft protocol.
- **Separation of concerns:** Distinguish node state and protocol logic from network registration and communication.
- **Incremental development:** Build the core components first, then integrate and validate them.
- **Explicit state management:** Make node roles, state, and protocol transitions understandable and traceable.
- **Correctness before optimization:** Prioritize protocol behavior and reliable testing before performance improvements.

## Project Structure

The project is organized around the Raft node, its communication interfaces, and the components required to coordinate protocol execution.

*The exact directory structure and module responsibilities can be documented here as the implementation evolves.*

## Building and Running

Build and execution instructions will be added once the main event loop and end-to-end execution flow are ready.

For now, the project is under active development and is not yet presented as a fully operational Raft cluster.

## Future Work

- Complete the main event loop.
- Integrate node state, timers, and RPC communication into the protocol's execution flow.
- Implement and validate leader election and heartbeat handling end to end.
- Exercise the implementation with multiple nodes.
- Add systematic tests for node failures and network disruptions.
- Investigate concurrency and timing-related edge cases.
- Improve observability and diagnostics to make protocol behavior easier to debug.
- Document the final architecture, implementation decisions, and test results.

## References

- Ongaro, Diego, and John Ousterhout. **In Search of an Understandable Consensus Algorithm (Extended Version).** Stanford University.

  [Read the Raft paper](https://raft.github.io/raft.pdf)

- [Raft Consensus Algorithm](https://raft.github.io/)

## Motivation

This project is an exploration of distributed systems through implementation.
I'm Building Raft from scratch to understand how distributed nodes coordinate,
how consensus is achieved
how protocol correctness is affected by failures, concurrency, and network communication.

The focus is not only on implementing the algorithm,
but also on understanding the design decisions and
rigorously validating the resulting system as it develops.
