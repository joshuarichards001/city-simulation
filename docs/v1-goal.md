# City Simulation V1

## Overview

Project SimCityLive is a persistent, real-time simulation of a digital city and its residents. The goal of V1 is to establish the core technical foundation: a backend capable of simulating 10,000 residents and a web client capable of serving a real-time view of a portion of this city to 1,000 concurrent users. V1 is a passive experience; users can only watch the city. All interactive social features are deferred to V2.

## V1 Goals & Success Metrics

Primary Goal: Launch a stable, 24/7 running simulation that meets the defined performance and scalability targets.

Success Metrics:

Uptime: The Go application maintains 99% uptime over a 7-day period.

Concurrency: The system handles 1,000 concurrent WebSocket connections without significant latency degradation (>500ms).

Simulation Integrity: The simulation of 10,000 residents runs without crashing or freezing. Residents follow their simple schedules correctly.

## V1 User Persona: The Passive Observer

For V1, we are building for a single user type:

The Observer: A user who visits the website to watch the city. They can see the characters move and understand, at a high level, what is happening. They cannot interact with the simulation.

## Functional Requirements (V1)

### Procedural Generation (Executed on Server Start)

City Generation:

The world will be a fixed-size 2D grid (e.g., 1000x1000).

The grid will be divided into two simple zones: a Residential Zone and a Commercial Zone. A simple rule, like x < 500 is residential and x >= 500 is commercial, is sufficient.

Buildings: Buildings are not physical objects in V1. They are simply named coordinate points on the grid. We will generate two types:

Housing (1 per resident): A randomly assigned coordinate within the Residential Zone.

Workplaces (1 per 10 residents): A randomly assigned coordinate within the Commercial Zone.

Resident Generation:

The system will generate 10,000 residents.

Each resident will have the following persistent attributes:

UUID: A unique identifier.

Name: Randomly generated from a predefined list of first/last names.

HomeCoords: The coordinates of their assigned Housing.

WorkCoords: The coordinates of their assigned Workplace.

### Simulation Core (Backend)

Resident Schedule: The V1 schedule is a simple, 24-hour loop based on server UTC time.

Work (09:00 - 17:00): Resident's goal is their WorkCoords. If at work, their status is "Working".

Home (17:01 - 08:59): Resident's goal is their HomeCoords. If at home, their status is "At Home" or "Sleeping".

Movement Logic:

When a resident's schedule changes their goal, they will begin moving from their current coordinates to their goal coordinates.

Pathfinding must be a simple A* (A-star) algorithm on the grid. For V1, residents can pass through each other. There are no obstacles.

Movement speed will be a fixed constant for all residents.

### Client Application (Frontend)

City View:

The client will render a fixed, non-pannable view of a single "block" of the city (e.g., coordinates 0,0 to 50,50).

The view will be rendered on an HTML Canvas. Residents are represented by simple colored pixels or circles.

Data Display:

Clicking on a resident's pixel will display their Name and current Status in a static UI element.

Connection Status: The UI must display the current connection status to the WebSocket server (Connecting, Connected, Disconnected).

## Non-Functional Requirements (NFRs)

Scalability (1,000 Concurrent Users):

The server-client communication must be "Intention-Based". The server sends commands (START_MOVE, SET_STATE), not continuous position updates.

The server must implement spatial partitioning. It will only send updates for residents within the client's fixed viewing area. This is the most critical NFR for achieving concurrency.

Performance:

WebSocket messages from server to client must be minimal. A START_MOVE command should contain personId, start/end coordinates, and total travel duration.

The client-side animation must be smooth (driven by requestAnimationFrame).

Deployment & Operations:

The Go application and the frontend files must be containerized using Docker.

The application will be deployed to a cloud provider. A simple Platform-as-a-Service (PaaS) like Google Cloud Run or a container service on DigitalOcean is recommended for V1.

Logging: The Go application must have structured logging (e.g., using log/slog) to output key events (server start, user connect/disconnect, simulation errors).

## V1 Technical Architecture

Backend:

Language: Go

Framework: Standard library for net/http.

WebSocket Library: gorilla/websocket.

Persistence: PostgreSQL. On initial server start, the procedural generation logic runs and populates the database. The running server then reads this data into memory for the live simulation. This separates the static "world" data from the dynamic "state" data.

Frontend:

Frameworks: None. Vanilla HTML, CSS, and TypeScript.

Rendering: HTML Canvas 2D API.

Deployment:

Containerization: Docker / Docker Compose.

Hosting: Google Cloud Run (recommended for its auto-scaling and simple container deployment).

## Scope & Phasing

To be crystal clear, the following features are OUT OF SCOPE for V1 and are planned for V2 and beyond:

NO user accounts or authentication.

NO "Coin of Fate" or any other user interaction with the simulation.

NO panning, zooming, or searching on the client. The view is fixed.

NO complex resident schedules (e.g., leisure, shopping, social).

NO resident-to-resident interaction.

NO complex building models or road networks.