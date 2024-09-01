"use client";

import React, { useState } from "react";
import styles from "./CategorySlider.module.css";
import {Button} from "antd";

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

const CategorySlider: React.FC = () => {
  const [activeCategory, setActiveCategory] = useState<string>(categories[0]);

    return (
        <div className="category-slider">
            {categories.map((category) => (
                <Button
                    key={category}
                    onClick={() => setActiveCategory(category)}
                    className={`category-button ${
                        activeCategory === category ? "active" : ""
                    }`}
                >
                    {category}
                </Button>
            ))}
        </div>
    );
};

export default CategorySlider;
