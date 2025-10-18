# Realtime Online City Simulation

A city simulation that users can watch play out in real time in their browser.

## Tech
- Frontend
	- HTML
	- Typescript
	- PixiJS
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
	- Plausible Analytics
	
## Programming approach

Ensure that all your implementations are simple and concise.