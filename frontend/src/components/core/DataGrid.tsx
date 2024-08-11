"use client";

import { DataGrid as MuiDataGrid, DataGridProps } from "@mui/x-data-grid";
import React, { useEffect } from "react";

const DataGrid = (props: DataGridProps) => {
  useEffect(() => {
    props.apiRef?.current.autosizeColumns({
      columns: props.columns.map((column: any) => column.field),
      expand: true,
    });
  }, [props.rows]);
  return (
    <div style={{ height: "100%", width: "100%" }}>
      <MuiDataGrid
        {...props}
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
            minHeight: 56,
            maxHeight: 80,
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
          columns: props.columns.map((column: any) => column.field),
          //includeHeaders: true,
          //includeOutliers: true,
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
};

export default DataGrid;
