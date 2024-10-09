"use client";
import { Card, Flex, Statistic } from "antd";
import {
  LineChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Line,
} from "recharts";
import { useMeeting } from "@videosdk.live/react-sdk";
const LivestreamStatistic = () => {
  const { participants } = useMeeting();
  return (
    <Flex vertical gap="middle">
      <Card>
        <Flex justify="space-between">
          <Statistic title="Doanh thu" value={0} />
          <Statistic title="Đơn hàng" value={0} />
          <Statistic title="Lượt xem" value={new Map(participants).size} />
          <Statistic title="Lượt follow" value={0} />
          <Statistic title="Doanh số" value={0} />
        </Flex>
      </Card>
      <Card>
        <LineChart
          data={[
            {
              date: "2024-10-04T08:00:00",
              totalOrders: 0,
              totalRevenue: 0,
            },
            {
              date: "2024-10-04T09:00:00",
              totalOrders: 0,
              totalRevenue: 0,
            },
            {
              date: "2024-10-04T10:00:00",
              totalOrders: 0,
              totalRevenue: 0,
            },
            {
              date: "2024-10-04T11:00:00",
              totalOrders: 0,
              totalRevenue: 0,
            },
            // Add more data points as needed
          ]}
          width={600}
          height={300}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis
            dataKey="date"
            interval={0}
            tickFormatter={(tick) => {
              return new Date(tick).toLocaleTimeString([], {
                hour: "2-digit",
                minute: "2-digit",
              });
            }}
          />
          <YAxis
            yAxisId="left"
            orientation="left"
            label={{
              value: "Total Orders",
              angle: -90,
              position: "insideLeft",
            }}
          />
          <YAxis
            yAxisId="right"
            orientation="right"
            label={{
              value: "Total Revenue",
              angle: -90,
              position: "insideRight",
            }}
          />
          <Tooltip />
          <Line
            type="monotone"
            dataKey="totalOrders"
            stroke="#82ca9d"
            yAxisId="left"
            name="Total Orders"
          />
          <Line
            type="monotone"
            dataKey="totalRevenue"
            stroke="#8884d8"
            yAxisId="right"
            name="Total Revenue"
          />
        </LineChart>
      </Card>
    </Flex>
  );
};

export default LivestreamStatistic;
