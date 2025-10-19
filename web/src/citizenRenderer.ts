import { Application, Assets, Sprite, Container } from "pixi.js";
import { StartMoveData } from "./types";

interface CitizenSprite {
  sprite: Sprite;
  startX: number;
  startY: number;
  targetX: number;
  targetY: number;
  startTime: number;
  duration: number;
  moving: boolean;
}

export class CitizenRenderer {
  private app: Application;
  private citizens: Map<number, CitizenSprite>;
  private container: Container;

  constructor(app: Application) {
    this.app = app;
    this.citizens = new Map();
    this.container = new Container();
    this.app.stage.addChild(this.container);
  }

  async init() {
    // Preload the bunny texture
    await Assets.load("/assets/bunny.png");
  }

  handleStartMove(data: StartMoveData) {
    let citizen = this.citizens.get(data.citizenId);

    if (!citizen) {
      // Create new citizen sprite
      const texture = Assets.get("/assets/bunny.png");
      const sprite = new Sprite(texture);
      sprite.anchor.set(0.5);
      sprite.position.set(data.fromX, data.fromY);
      this.container.addChild(sprite);

      citizen = {
        sprite,
        startX: data.fromX,
        startY: data.fromY,
        targetX: data.toX,
        targetY: data.toY,
        startTime: Date.now(),
        duration: data.duration,
        moving: true,
      };

      this.citizens.set(data.citizenId, citizen);
    } else {
      // Update existing citizen with new movement
      citizen.startX = data.fromX;
      citizen.startY = data.fromY;
      citizen.targetX = data.toX;
      citizen.targetY = data.toY;
      citizen.startTime = Date.now();
      citizen.duration = data.duration;
      citizen.moving = true;
      citizen.sprite.position.set(data.fromX, data.fromY);
    }
  }

  update() {
    const now = Date.now();

    this.citizens.forEach((citizen) => {
      if (!citizen.moving) return;

      const elapsed = now - citizen.startTime;
      const progress = Math.min(elapsed / citizen.duration, 1);

      // Linear interpolation
      const currentX =
        citizen.startX + (citizen.targetX - citizen.startX) * progress;
      const currentY =
        citizen.startY + (citizen.targetY - citizen.startY) * progress;

      citizen.sprite.position.set(currentX, currentY);

      // Stop moving when animation is complete
      if (progress >= 1) {
        citizen.moving = false;
      }
    });
  }

  startAnimationLoop() {
    this.app.ticker.add(() => {
      this.update();
    });
  }
}
