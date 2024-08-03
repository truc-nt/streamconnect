"use client";

import React, { useRef, useState } from "react";
import { Box, Typography, IconButton } from "@mui/material";
import { ArrowBack, ArrowForward } from "@mui/icons-material";
//import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import DetailModal from "./DetailModal";

interface Product {
  id: number;
  name: string;
  image: string;
  description: string;
  categories: string[];
  categoryValues: {
    [category: string]: string[];
  };
  prices: {
    platform: string;
    category1: string;
    category2: string;
    price: number;
  }[];
}

const products: Product[] = [
  {
    id: 1,
    name: "Product 1",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 1",
    categories: ["Màu", "Size"],
    categoryValues: {
      Màu: ["Đỏ", "Xanh", "Vàng", "Trắng"],
      Size: ["S", "M"],
    },
    prices: [
      { platform: "Shopee", category1: "Đỏ", category2: "S", price: 100000 },
      { platform: "Shopee", category1: "Đỏ", category2: "M", price: 100000 },
      { platform: "Shopee", category1: "Xanh", category2: "S", price: 105000 },
      { platform: "Shopee", category1: "Xanh", category2: "M", price: 110000 },
      { platform: "Shopee", category1: "Vàng", category2: "S", price: 105000 },
      { platform: "Shopee", category1: "Vàng", category2: "M", price: 110000 },
      { platform: "Shopee", category1: "Trắng", category2: "S", price: 105000 },
      { platform: "Shopee", category1: "Trắng", category2: "M", price: 110000 },
      { platform: "Lazada", category1: "Đỏ", category2: "S", price: 105000 },
      { platform: "Lazada", category1: "Đỏ", category2: "M", price: 125000 },
      { platform: "Lazada", category1: "Xanh", category2: "S", price: 100000 },
      { platform: "Lazada", category1: "Xanh", category2: "M", price: 115000 },
      { platform: "Lazada", category1: "Vàng", category2: "S", price: 105000 },
      { platform: "Lazada", category1: "Vàng", category2: "M", price: 130000 },
      { platform: "Lazada", category1: "Trắng", category2: "S", price: 105000 },
      { platform: "Lazada", category1: "Trắng", category2: "M", price: 125000 },
    ],
  },
  {
    id: 2,
    name: "Product 2",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 2",
    categories: ["Màu", "Size"],
    categoryValues: {
      Màu: ["Đỏ", "Xanh", "Vàng", "Trắng"],
      Size: ["L", "XL"],
    },
    prices: [
      { platform: "Shopee", category1: "Đỏ", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Đỏ", category2: "XL", price: 120000 },
      { platform: "Shopee", category1: "Xanh", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Xanh", category2: "XL", price: 120000 },
      { platform: "Lazada", category1: "Đỏ", category2: "L", price: 105000 },
      { platform: "Lazada", category1: "Đỏ", category2: "XL", price: 125000 },
    ],
  },
  {
    id: 3,
    name: "Product 3",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 3",
    categories: [],
    categoryValues: {},
    prices: [
      { platform: "Shopee", category1: "", category2: "", price: 50000 },
      { platform: "Lazada", category1: "", category2: "", price: 65000 },
    ],
  },
  {
    id: 4,
    name: "Product 4",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 4",
    categories: ["Size"],
    categoryValues: {
      Size: ["S", "M"],
    },
    prices: [
      { platform: "Shopee", category1: "S", category2: "", price: 80000 },
      { platform: "Shopee", category1: "M", category2: "", price: 90000 },
      { platform: "Lazada", category1: "S", category2: "", price: 95000 },
      { platform: "Lazada", category1: "M", category2: "", price: 125000 },
    ],
  },
  {
    id: 5,
    name: "Product 5",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 5",
    categories: ["Màu", "Size"],
    categoryValues: {
      Màu: ["Đỏ", "Xanh"],
      Size: ["L", "XL"],
    },
    prices: [
      { platform: "Shopee", category1: "Đỏ", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Đỏ", category2: "XL", price: 120000 },
      { platform: "Shopee", category1: "Xanh", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Xanh", category2: "XL", price: 120000 },
      { platform: "Lazada", category1: "Đỏ", category2: "L", price: 105000 },
      { platform: "Lazada", category1: "Đỏ", category2: "XL", price: 125000 },
    ],
  },
  {
    id: 6,
    name: "Product 6",
    image:
      "https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg",
    description: "Mô tả sản phẩm 6",
    categories: ["Màu", "Size"],
    categoryValues: {
      Màu: ["Đỏ", "Xanh"],
      Size: ["L", "XL"],
    },
    prices: [
      { platform: "Shopee", category1: "Đỏ", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Đỏ", category2: "XL", price: 120000 },
      { platform: "Shopee", category1: "Xanh", category2: "L", price: 100000 },
      { platform: "Shopee", category1: "Xanh", category2: "XL", price: 120000 },
      { platform: "Lazada", category1: "Đỏ", category2: "L", price: 105000 },
      { platform: "Lazada", category1: "Đỏ", category2: "XL", price: 125000 },
    ],
  },
];

const ProductList: React.FC = () => {
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const formatPrice = (price: number) => {
    return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
  };

  const scrollContainerRef = useRef<HTMLDivElement>(null);

  const scrollLeft = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: -300, behavior: "smooth" });
    }
  };

  const scrollRight = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: 300, behavior: "smooth" });
    }
  };

  const handleProductClick = (product: Product) => {
    setSelectedProduct(product);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
  };

  const getPriceRange = (prices: { price: number }[]) => {
    if (prices.length === 0) return "";
    const minPrice = Math.min(...prices.map((p) => p.price));
    const maxPrice = Math.max(...prices.map((p) => p.price));
    return `${formatPrice(minPrice)} đ`;
  };

  return (
    <>
      <Box
        sx={{
          width: "100%",
          mx: "auto",
          my: 3,
          display: "flex",
          alignItems: "center",
        }}
      >
        <IconButton
          onClick={scrollLeft}
          sx={{
            backgroundColor: "#01E0EE",
            borderRadius: "50%",
            "&:hover": { backgroundColor: "#01E0EE" },
            marginRight: 2,
          }}
        ></IconButton>

        <Box
          ref={scrollContainerRef}
          sx={{
            overflowX: "auto",
            display: "flex",
            gap: 2,
            msOverflowStyle: "none",
            scrollbarWidth: "none",
            "::-webkit-scrollbar": { display: "none" },
          }}
        >
          {products.map((product) => (
            <Box
              key={product.id}
              onClick={() => handleProductClick(product)}
              sx={{
                flex: "0 0 auto",
                width: "30%",
                backgroundColor: "#333",
                padding: 2,
                borderRadius: 2,
                display: "flex",
                alignItems: "center",
                border: "1px solid white",
                cursor: "pointer",
              }}
            >
              <img
                src={product.image}
                alt={product.name}
                style={{
                  width: "50%",
                  height: "auto",
                  borderRadius: "8px",
                  marginRight: "16px",
                  objectFit: "cover",
                }}
              />
              <Box sx={{ display: "flex", flexDirection: "column" }}>
                <Typography
                  variant="body1"
                  sx={{
                    color: "white",
                    fontSize: "14px",
                    fontWeight: "normal",
                    whiteSpace: "nowrap",
                    overflow: "hidden",
                    textOverflow: "ellipsis",
                    maxWidth: "100%",
                  }}
                >
                  {product.name}
                </Typography>
                <Typography
                  variant="body2"
                  sx={{ color: "white", fontSize: "18px" }}
                >
                  {getPriceRange(product.prices)}
                </Typography>
              </Box>
            </Box>
          ))}
        </Box>

        <IconButton
          onClick={scrollRight}
          sx={{
            backgroundColor: "#01E0EE",
            borderRadius: "50%",
            "&:hover": { backgroundColor: "#01E0EE" },
            marginLeft: 2,
          }}
        ></IconButton>
      </Box>

      {selectedProduct && (
        <DetailModal
          open={isModalOpen}
          onClose={handleCloseModal}
          product={selectedProduct}
        />
      )}
    </>
  );
};

export default ProductList;
