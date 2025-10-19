export interface StartMoveData {
  citizenId: number;
  fromX: number;
  fromY: number;
  toX: number;
  toY: number;
  duration: number;
}

export interface WebSocketMessage {
  type: string;
  data: unknown;
}
