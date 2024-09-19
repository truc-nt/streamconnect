import { useEffect } from 'react';
import { connectSocket, disconnectSocket } from '@/api/socket';

export const useWebSocket = (onNotificationReceive: any) => {
  useEffect(() => {
    connectSocket(onNotificationReceive);

    return () => {
      disconnectSocket();
    };
  }, []);
};