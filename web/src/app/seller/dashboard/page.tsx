"use client";
import { Button, Card, Flex, Col, Row, Statistic } from "antd";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

const data = [
  { date: "2024-01", revenue: 0 },
  { date: "2024-02", revenue: 0 },
  { date: "2024-03", revenue: 0 },
  { date: "2024-04", revenue: 0 },
  { date: "2024-05", revenue: 0 },
];

const Page = () => {
  return (
    <Flex vertical gap="middle">
      <Card>
        <Flex justify="space-between">
          <Statistic title="Doanh thu" value={0} />
          <Statistic title="Đơn hàng" value={0} />
          <Statistic title="Lượt xem" value={112893} />
          <Statistic title="Lượt follow" value={0} />
          <Statistic title="Doanh số" value={0} />
        </Flex>
      </Card>
      <Card>
        <LineChart data={data} width={600} height={300}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis
            dataKey="date"
            interval={0}
            tickFormatter={(tick) =>
              new Date(tick).toLocaleDateString("en-US", {
                month: "short",
                year: "numeric",
              })
            } // Format date
          />
          <YAxis />
          <Tooltip />
          <Line type="monotone" dataKey="revenue" stroke="#8884d8" />
        </LineChart>
      </Card>
      <Card>
        <LineChart
          data={[
            {
              date: "2024-10-04T08:00:00",
              totalOrders: 10,
              totalRevenue: 1000,
            },
            {
              date: "2024-10-04T09:00:00",
              totalOrders: 20,
              totalRevenue: 1500,
            },
            {
              date: "2024-10-04T10:00:00",
              totalOrders: 30,
              totalRevenue: 2000,
            },
            {
              date: "2024-10-04T11:00:00",
              totalOrders: 40,
              totalRevenue: 2500,
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

export default Page;
