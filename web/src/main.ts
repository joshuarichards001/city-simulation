import { Application } from "pixi.js";
import { WebSocketClient } from "./websocket";
import { CitizenRenderer } from "./citizenRenderer";
import { StartMoveData } from "./types";

(async () => {
  // Create a new application
  const app = new Application();

  // Initialize the application
  await app.init({ background: "#1099bb", resizeTo: window });

  // Append the application canvas to the document body
  document.getElementById("pixi-container")!.appendChild(app.canvas);

  // Create citizen renderer
  const citizenRenderer = new CitizenRenderer(app);
  await citizenRenderer.init();

  // Connect to WebSocket
  const ws = new WebSocketClient();

  // Handle MOVE messages
  ws.on("MOVE", (data) => {
    citizenRenderer.handleStartMove(data as StartMoveData);
  });

  ws.connect();

  // Start animation loop
  citizenRenderer.startAnimationLoop();
})();
