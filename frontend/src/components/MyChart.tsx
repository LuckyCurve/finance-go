import React, { useEffect, useRef } from "react";
import * as echarts from "echarts";

interface Asset {
	Name: string;
	currency_type: string;
	Currency: number;
	ID: number;
	CreateTime: string;
	UpdateTime: string;
}

interface ApiResponse {
	code: number;
	data: Asset[];
}

interface SankeyChartProps {
	apiUrl: string; // 数据接口
	width?: string;
	height?: string;
}

const SankeyChart: React.FC<SankeyChartProps> = ({
	apiUrl,
	width = "800px",
	height = "600px",
}) => {
	const chartRef = useRef<HTMLDivElement>(null);
	const chartInstance = useRef<echarts.EChartsType | null>(null);

	useEffect(() => {
		if (!chartRef.current) return;

		chartInstance.current = echarts.init(chartRef.current);

		fetch(apiUrl)
			.then((res) => res.json())
			.then((data: ApiResponse) => {
				if (data.code !== 200) {
					console.error("接口返回错误");
					return;
				}

				const assets = data.data;

				// 1. 构造节点
				const currencyNodes = Array.from(
					new Set(assets.map((a) => a.currency_type.toUpperCase()))
				).map((name) => ({ name }));
				const assetNodes = assets.map((a) => ({ name: a.Name }));
				const nodes = [...assetNodes, ...currencyNodes];

				// 2. 构造连接
				const links = assets.map((a) => ({
					source: a.Name,
					target: a.currency_type.toUpperCase(),
					value: a.Currency,
				}));

				// 3. 初始化 ECharts option
				const option: echarts.EChartsOption = {
					title: { text: "资产桑基图" },
					tooltip: { trigger: "item", triggerOn: "mousemove" },
					series: [
						{
							type: "sankey",
							data: nodes,
							links: links,
							emphasis: { focus: "adjacency" },
							label: {
								color: "#000",
								// ✅ 删除 RichText 或复杂 borderRadius，避免 TS 报错
								fontSize: 12,
							},
						},
					],
				};

				chartInstance.current?.setOption(option as echarts.EChartOption);
			})
			.catch((err) => console.error(err));

		// 4. 响应式
		const handleResize = () => chartInstance.current?.resize();
		window.addEventListener("resize", handleResize);
		return () => {
			window.removeEventListener("resize", handleResize);
			chartInstance.current?.dispose();
		};
	}, [apiUrl]);

	return <div ref={chartRef} style={{ width, height }} />;
};

export default SankeyChart;
