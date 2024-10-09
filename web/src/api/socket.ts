import { Client } from "@stomp/stompjs";
export const socketClient = new Client({
  brokerURL: "ws://localhost:8080/ws",
  onWebSocketError: (evt) => {
    console.log(evt);
  },
});
export function connectSocket(onNotificationReceive: any) {
  const token = localStorage.getItem("token");
  if (!token) {
    return;
  }
  socketClient.onConnect = () => {
    socketClient.subscribe("/user/topic/notification", (message) => {
      console.log("receive message", message);
      const notification: Notification = JSON.parse(message.body);
      onNotificationReceive(notification);
    });
  };
  socketClient.connectHeaders = {
    Authorization: `Bearer ${token}`,
  };
  socketClient.activate();
}

export function disconnectSocket() {
  socketClient.deactivate();
}
