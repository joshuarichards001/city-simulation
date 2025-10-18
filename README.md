# City Simulation

## Overview
A online real time city simulation. The city will have 10,000 residents following a daily routine. The city has a commercial and residential area. The residents will have a job and a home.

The actual simulation will happen in the Go server and be passed to the client via WebSockets. The client will render the simulation using HTML5 Canvas.

## Tech
- Frontend
	- HTML
	- Typescript
	- PixiJS
	- Art by LimeZu
- Backend
	- Go
	- Gorilla WebSockets
	- Serve static files from Go app
- Hosting
	- Fly.io
	- Cloudflare
		- Cache frontend
		- Domain
		- Rate limit WebSocket connections
	- Plausible