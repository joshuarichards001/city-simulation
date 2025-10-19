export class WebSocketClient {
  private ws: WebSocket | null = null;
  private url: string;

  constructor() {
    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    const host = import.meta.env.DEV ? "localhost:8080" : window.location.host;
    this.url = `${protocol}//${host}/ws`;
  }

  connect() {
    this.ws = new WebSocket(this.url);

    this.ws.onopen = () => {
      console.log("WebSocket connected");
    };

    this.ws.onmessage = (event) => {
      console.log("Message received:", event.data);
    };

    this.ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.ws.onclose = () => {
      console.log("WebSocket closed");
    };
  }

  send(message: string) {
    this.ws?.send(message);
  }
}
