"use client";

import { DataGrid as MuiDataGrid } from "@mui/x-data-grid";
import React from "react";

interface IDataGridProps {
  rows: any;
  columns: any;
  getRowId?: (row: any) => string;
}

export default function DataGrid({ rows, columns, getRowId }: IDataGridProps) {
  console.log(rows);
  return (
    <div style={{ height: "100%", width: "100%" }}>
      <MuiDataGrid
        checkboxSelection
        rows={rows}
        columns={columns}
        getRowId={getRowId}
        sx={{
          border: 0,
          "& .MuiDataGrid-columnHeaders *": {
            border: 0,
            borderRadius: "5px",
            color: "primary.main",
            "& .MuiDataGrid-columnHeaderTitle": {
              fontWeight: 900,
            },
          },
          "& .MuiDataGrid-row, & .MuiDataGrid-columnHeaders": {
            border: 1,
            borderRadius: "5px",
            "*:hover, *:active, *:focus": {
              outline: "transparent",
            },
            backgroundColor: "background.default",
          },
          "& .MuiDataGrid-cell": {
            border: 0,
            py: 1,
            display: "flex",
            alignItems: "center",
            overflow: "hidden",
            textOverflow: "ellipsis",
          },
          "& .MuiCheckbox-root": {
            color: "inherit",
          },
          "*::after": {
            width: "0 !important",
          },
        }}
        autosizeOnMount={true}
        autosizeOptions={{
          columns: columns.map((column: any) => column.field),
          includeHeaders: true,
          expand: true,
        }}
        getRowSpacing={(params) => {
          return {
            top: params.isFirstVisible ? 8 : 4,
            bottom: params.isLastVisible ? 8 : 4,
          };
        }}
        getRowHeight={() => "auto"}
        autoHeight
      />
    </div>
  );
}
