import { Application } from "pixi.js";
import { WebSocketClient } from "./websocket";
import { CitizenRenderer } from "./citizenRenderer";
import { StartMoveData } from "./types";
import { CityRenderer } from "./cityRenderer";

(async () => {
  const app = new Application();
  await app.init({ background: "#111", resizeTo: window });
  document.getElementById("pixi-container")!.appendChild(app.canvas);

  const cityRenderer = new CityRenderer(app);
  await cityRenderer.init();

  const citizenRenderer = new CitizenRenderer(app);
  await citizenRenderer.init();

  const ws = new WebSocketClient();
  ws.on("MOVE", (data) => {
    citizenRenderer.handleStartMove(data as StartMoveData);
  });
  ws.connect();

  cityRenderer.renderCity();
  citizenRenderer.startAnimationLoop();
})();
