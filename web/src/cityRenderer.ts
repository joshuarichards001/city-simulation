import { Application, Container, Graphics } from "pixi.js";

type CityGrid = boolean[][];

export class CityRenderer {
  private app: Application;
  private city: CityGrid | null;
  private container: Container;
  private readonly tileSize = 100;

  constructor(app: Application) {
    this.app = app;
    this.city = null;
    this.container = new Container();
    this.app.stage.addChild(this.container);
  }

  async init(): Promise<void> {
    await this.loadCity();
  }

  renderCity(): void {
    this.container.removeChildren();

    if (!this.city || this.city.length === 0) {
      return;
    }

    for (let y = 0; y < this.city.length; y++) {
      for (let x = 0; x < this.city[y].length; x++) {
        const isBuilding = this.city[y][x] === true;

        const square = new Graphics()
          .rect(
            x * this.tileSize,
            y * this.tileSize,
            this.tileSize,
            this.tileSize,
          )
          .fill(isBuilding ? "brown" : "gray");

        this.container.addChild(square);
      }
    }
  }

  private async loadCity(): Promise<void> {
    try {
      const res = await fetch("/city.json");
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const { grid } = (await res.json()) as { grid: CityGrid };
      this.city = grid ?? [];
    } catch (error) {
      console.error("Error loading city data:", error);
      this.city = [];
    }
  }
}
