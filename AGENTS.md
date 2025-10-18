# AGENTS.md - AI Agent Guidelines for City Simulation

## Project Overview

This is **City Simulation** (Project SimCityLive) - a persistent, real-time simulation of a digital city with 10,000 residents. The simulation runs 24/7 on a Go backend and streams updates to web clients via WebSockets. Clients render the simulation using HTML5 Canvas.

**Current Phase:** V1 - Building the core technical foundation
**V1 Goal:** A passive viewing experience with 10,000 simulated residents and support for 1,000 concurrent viewers.

## Tech Stack

### Backend
- **Language:** Go (standard library)
- **WebSocket:** gorilla/websocket
- **Database:** PostgreSQL
- **Architecture:** In-memory simulation with persistent storage

### Frontend
- **No frameworks** - Vanilla HTML, CSS, and TypeScript
- **Rendering:** HTML5 Canvas 2D API
- **Communication:** WebSockets

### Deployment
- **Containerization:** Docker / Docker Compose
- **Target Platform:** Google Cloud Run (PaaS) or similar container service

## Architecture Principles

### Core Concepts

1. **Procedural Generation on Server Start**
   - City is a fixed 2D grid (e.g., 1000x1000)
   - Simple zoning: x < 500 = Residential, x >= 500 = Commercial
   - Buildings are coordinate points, not physical objects
   - 10,000 residents with homes and workplaces

2. **Intention-Based Communication**
   - Server sends commands (START_MOVE, SET_STATE), not continuous position updates
   - Minimizes bandwidth and enables high concurrency
   - Client interpolates movement using requestAnimationFrame

3. **Spatial Partitioning**
   - Critical for scaling to 1,000 concurrent users
   - Server only sends updates for residents in client's viewing area
   - Viewing area is fixed (e.g., 0,0 to 50,50) - no panning/zooming in V1

4. **Simple Simulation Logic**
   - 24-hour schedule: Work (09:00-17:00) at WorkCoords, Home (17:01-08:59) at HomeCoords
   - A* pathfinding on grid
   - Fixed movement speed for all residents
   - No collisions - residents pass through each other

## Key V1 Success Metrics

- **Uptime:** 99% over 7 days
- **Concurrency:** 1,000 WebSocket connections with <500ms latency
- **Simulation Integrity:** 10,000 residents running without crashes

## V1 Scope - What's IN

✅ 10,000 residents with simple daily schedules
✅ Procedurally generated city (residential + commercial zones)
✅ A* pathfinding for resident movement
✅ WebSocket streaming to clients
✅ Fixed canvas view of city block
✅ Click resident to see Name and Status
✅ Connection status display
✅ Spatial partitioning for performance
✅ PostgreSQL persistence of world data
✅ Structured logging (log/slog)
✅ Docker containerization

## V1 Scope - What's OUT (Deferred to V2+)

❌ User accounts or authentication
❌ User interaction with simulation (Coin of Fate, etc.)
❌ Panning, zooming, or searching the map
❌ Complex schedules (leisure, shopping, social activities)
❌ Resident-to-resident interaction
❌ Complex building models or road networks
❌ Multiple viewing areas or camera controls

## File Structure & Conventions

### Backend (Go)
- Use standard Go project layout
- `cmd/` for application entry points
- `internal/` for private application code
- `pkg/` for reusable packages (if needed)
- Use `log/slog` for structured logging
- Database migrations should be versioned

### Frontend (TypeScript)
- Keep vanilla - no build frameworks unless absolutely necessary
- Separate concerns: WebSocket client, Canvas renderer, UI state
- Use TypeScript for type safety
- Canvas rendering loop driven by `requestAnimationFrame`

### Database Schema
- **Static Data:** City layout, buildings, resident profiles (generated once)
- **Dynamic Data:** Current positions, states (runtime only, or cached)
- Consider: Does dynamic state need persistence, or can simulation restart fresh?

## Development Guidelines for AI Agents

### When Writing Go Code

1. **Use gorilla/websocket** for WebSocket handling
2. **Implement spatial partitioning early** - don't defer this optimization
3. **Use goroutines efficiently** - one per WebSocket connection is acceptable
4. **Implement graceful shutdown** with context cancellation
5. **Structure logs** with `log/slog` including: connection events, errors, performance metrics
6. **Separate concerns:** simulation logic, WebSocket handling, data persistence
7. **Use channels** for safe communication between simulation and WebSocket handlers

### When Writing Frontend Code

1. **No frameworks** - use vanilla TypeScript/JavaScript
2. **Efficient Canvas rendering:**
   - Clear and redraw only what's needed
   - Use `requestAnimationFrame` for smooth animation
   - Consider double buffering if performance issues arise
3. **Client-side interpolation:**
   - When receiving START_MOVE, calculate intermediate positions
   - Smooth movement between start and end coordinates over duration
4. **Handle WebSocket lifecycle:**
   - Show connection status clearly
   - Auto-reconnect on disconnect
   - Queue commands if connection drops (or discard - decide based on use case)
5. **Keep UI minimal** - this is a passive viewer, not an interactive app

### When Designing APIs/Protocols

1. **WebSocket Message Format:**
   - Use JSON for simplicity in V1
   - Consider binary protocols (MessagePack, Protobuf) if performance issues arise
   - Message types: INIT, START_MOVE, SET_STATE, UPDATE_VIEW, etc.

2. **Message Structure Example:**
   ```json
   {
     "type": "START_MOVE",
     "residentId": "uuid",
     "from": {"x": 10, "y": 20},
     "to": {"x": 15, "y": 25},
     "duration": 5000
   }
   ```

3. **Keep payloads minimal** - bandwidth is the bottleneck for 1,000 users

### Performance Considerations

1. **Spatial Partitioning is Non-Negotiable**
   - Implement grid-based spatial indexing
   - Each client subscribes to specific grid cells
   - Only send updates for residents in subscribed cells

2. **Movement Updates**
   - Send START_MOVE when resident begins moving
   - Client calculates intermediate positions
   - Only send correction if resident's path changes

3. **Database Usage**
   - Load world data into memory on startup
   - Minimize database queries during simulation
   - Consider read-only replicas if database becomes bottleneck

4. **Monitoring**
   - Log connection count
   - Log message throughput
   - Monitor goroutine count
   - Track memory usage

### Testing Strategy

1. **Load Testing:** Simulate 1,000 WebSocket connections early
2. **Simulation Testing:** Verify 10,000 residents move correctly
3. **Integration Testing:** End-to-end WebSocket message flow
4. **Performance Testing:** Measure message latency under load

### Common Pitfalls to Avoid

1. ❌ **Don't send position updates every frame** - use intention-based updates
2. ❌ **Don't broadcast all resident updates to all clients** - use spatial partitioning
3. ❌ **Don't implement features from V2** - stick to V1 scope
4. ❌ **Don't add frameworks** without justification - keep it simple
5. ❌ **Don't ignore the 1,000 user target** - design for scale from the start

### Docker & Deployment

1. **Multi-stage builds** for smaller images
2. **Environment variables** for configuration (DB connection, port, etc.)
3. **Health check endpoint** for container orchestration
4. **Graceful shutdown** to handle SIGTERM properly
5. **Docker Compose** for local development (Go app + PostgreSQL)

## Data Models

### Resident
- UUID (string)
- Name (string)
- HomeCoords (x, y)
- WorkCoords (x, y)
- CurrentCoords (x, y) - runtime only
- Status (string): "Working", "At Home", "Sleeping", "Moving"
- CurrentGoal (x, y) - runtime only

### Building
- ID (int/string)
- Type ("Housing" | "Workplace")
- Coords (x, y)
- Zone ("Residential" | "Commercial")

### ViewArea (Client Subscription)
- ClientID
- BoundsX (min, max)
- BoundsY (min, max)

## Questions to Ask Yourself

When implementing features, consider:

1. **Does this scale to 10,000 residents?**
2. **Does this scale to 1,000 concurrent connections?**
3. **Is this in scope for V1?** (Check the V1 scope section)
4. **Does this align with "intention-based" communication?**
5. **Have I implemented spatial partitioning correctly?**
6. **Is this the simplest solution that works?**

## Resources & References

- **A* Algorithm:** Standard implementation for grid-based pathfinding
- **gorilla/websocket:** https://github.com/gorilla/websocket
- **Canvas 2D API:** MDN Web Docs for reference
- **Spatial Partitioning:** Grid-based approach (divide world into cells)

## Version History

- **V1 (Current):** Passive observation of 10,000 residents, 1,000 concurrent users
- **V2 (Future):** User interaction, "Coin of Fate", social features
- **V3+ (Future):** TBD based on V1/V2 learnings

---

**Remember:** V1 is about proving the technical foundation works. Keep it simple, focus on performance and scalability, and resist feature creep. Every line of code should serve the V1 goals.