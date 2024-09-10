"use client";

import React, { useState } from "react";
import { Button, Flex } from "antd";

const categories = [
  "Dành cho bạn",
  "Bách hóa",
  "Thời trang",
  "Phong cách sống",
  "Điện tử",
  "Làm đẹp",
  "Giáo dục",
  "Sách",
];

const CategorySlider = () => {
  const [activeCategory, setActiveCategory] = useState<string>(categories[0]);

  return (
    <Flex gap="small">
      {categories.map((category) => (
        <Button
          key={category}
          onClick={() => setActiveCategory(category)}
          type={activeCategory === category ? "primary" : "default"}
        >
          {category}
        </Button>
      ))}
    </Flex>
  );
};

export default CategorySlider;
